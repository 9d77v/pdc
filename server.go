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
	"github.com/9d77v/pdc/middleware"
	"github.com/9d77v/pdc/models"
	"github.com/9d77v/pdc/models/nats"
	"github.com/9d77v/wspush/redishub"
	"github.com/nats-io/stan.go"
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
	mux.HandleFunc("/card/", middleware.HandleCard())
	mux.HandleFunc("/ws/iot/device", middleware.HandleIotDevice())
	mux.HandleFunc("/ws/iot/telemetry", redishub.Hub.HandlerDynamicChannel())
	mux.HandleFunc("/ws/iot/health", redishub.Hub.HandlerDynamicChannel())
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
	qsub1, err := nats.Client.QueueSubscribe(nats.SubjectVideo,
		nats.GroupVideo, consumers.HandleVideoMSG)
	if err != nil {
		log.Panicln("QueueSubscribe error:", err)
	}
	defer func() {
		err = qsub1.Unsubscribe()
		if err != nil {
			log.Println("qsub1 Unsubscribe error:", err)
		}
		err = qsub1.Close()
		if err != nil {
			log.Println("qsub1 Close error:", err)
		}
	}()
	qsub2, err := nats.Client.QueueSubscribe(nats.SubjectDeviceData, nats.GroupSaveDeviceData, consumers.HandleDeviceMSG, stan.DurableName("dur"))
	if err != nil {
		log.Panicln("SubscribeDeviceAttribute error:", err)
	}
	defer func() {
		err = qsub2.Unsubscribe()
		if err != nil {
			log.Println("qsub2 Unsubscribe error:", err)
		}
		err = qsub2.Close()
		if err != nil {
			log.Println("qsub2 Close error:", err)
		}
	}()
	qsub3, err := nats.Client.QueueSubscribe(nats.SubjectDeviceData, nats.GroupPublishDeviceData, consumers.PublishDeviceData, stan.DurableName("dur"))
	if err != nil {
		log.Panicln("SubscribeDeviceAttribute error:", err)
	}
	defer func() {
		err = qsub3.Unsubscribe()
		if err != nil {
			log.Println("qsub3 Unsubscribe error:", err)
		}
		err = qsub3.Close()
		if err != nil {
			log.Println("qsub3 Close error:", err)
		}
	}()
	go consumers.SaveDeviceTelemetry()
	go consumers.SaveDeviceHealth()
	go func() {
		errc <- srv.ListenAndServe()
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	}()

	log.Printf("exiting (%v)", <-errc)
	srvCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = srv.Shutdown(srvCtx)
	if err != nil {
		log.Println("server shut down error:", err)
	}
	log.Println("exited")
}
