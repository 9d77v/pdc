package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/9d77v/pdc/consumers"
	"github.com/9d77v/pdc/graph"
	"github.com/9d77v/pdc/graph/generated"
	"github.com/9d77v/pdc/iot/esp8266"
	"github.com/9d77v/pdc/iot/sdk"
	"github.com/9d77v/pdc/middleware"
	"github.com/9d77v/pdc/models"
	"github.com/9d77v/pdc/models/nats"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	mux := http.NewServeMux()
	apiHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	if models.DEBUG {
		http.Handle("/docs", playground.Handler("GraphQL playground", "/api"))
	}
	mux.Handle("/api", middleware.Auth()(apiHandler))

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/build/index.html")
	})
	mux.HandleFunc("/app/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/build/index.html")
	})
	mux.HandleFunc("/admin/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/build/index.html")
	})
	mux.Handle("/", http.FileServer(http.Dir("ui/build")))
	errc := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	qsub1, _ := nats.Client.QueueSubscribe(nats.SubjectVideo,
		nats.GroupVideo, consumers.HandleVideoMSG)
	defer func() {
		qsub1.Unsubscribe()
		qsub1.Close()
	}()
	iotsdk := sdk.NewIotSDK()
	qsub2, _ := iotsdk.NatsSubscribe(consumers.ReplyDeviceMSG)
	defer func() {
		qsub2.Unsubscribe()
	}()

	qsub3, _ := iotsdk.SubscribeDeviceFilter(consumers.FilterDeviceMSG)
	defer func() {
		qsub3.Unsubscribe()
		qsub3.Close()
	}()
	qsub4, _ := iotsdk.SubscribeSaveDeviceData(consumers.SaveDeviceData)
	defer func() {
		qsub4.Unsubscribe()
		qsub4.Close()
	}()

	go func() {
		errc <- srv.ListenAndServe()
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	}()

	ctx, cancel := context.WithCancel(context.Background())
	go esp8266.CollectData(ctx)
	log.Printf("exiting (%v)", <-errc)
	cancel()
	srvCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(srvCtx)
	log.Println("exited")
}
