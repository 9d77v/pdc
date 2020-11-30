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
	"github.com/9d77v/wspush/redishub"
	"github.com/nats-io/stan.go"

	"github.com/9d77v/pdc/internal/consts"
	"github.com/9d77v/pdc/internal/consumer"
	"github.com/9d77v/pdc/internal/db/mq"
	"github.com/9d77v/pdc/internal/graph"
	"github.com/9d77v/pdc/internal/graph/generated"
	_ "github.com/9d77v/pdc/internal/init"
	"github.com/9d77v/pdc/internal/middleware"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: getServerMux(),
	}
	qsubs := initSubScriptions()
	defer unSubscribe(qsubs)
	go consumer.SaveDeviceTelemetry()
	go consumer.SaveDeviceHealth()
	go func() {
		errc <- srv.ListenAndServe()
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	}()

	log.Printf("exiting (%v)", <-errc)
	srvCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := srv.Shutdown(srvCtx)
	if err != nil {
		log.Println("server shut down error:", err)
	}
	log.Println("exited")
}

func getServerMux() *http.ServeMux {
	mux := http.NewServeMux()
	apiHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)
	if consts.DEBUG {
		http.Handle("/docs", playground.Handler("GraphQL playground", "/api"))
	}
	mux.Handle("/api", middleware.Auth()(apiHandler))
	mux.HandleFunc("/pdc/", middleware.HandleCard())
	mux.HandleFunc("/ws/iot/device", middleware.HandleIotDevice())
	mux.HandleFunc("/ws/iot/telemetry",
		redishub.Hub.HandlerDynamicChannel(mq.SubjectDeviceTelemetryPrefix, middleware.CheckToken))
	mux.HandleFunc("/ws/iot/health",
		redishub.Hub.HandlerDynamicChannel(mq.SubjectDeviceHealthPrefix, middleware.CheckToken))
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/build/index.html")
	})
	mux.HandleFunc("/app/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/build/index.html")
	})
	mux.HandleFunc("/admin/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/build/index.html")
	})
	mux.Handle("/", http.FileServer(http.Dir("web/build")))
	return mux
}

func initSubScriptions() []stan.Subscription {
	qsub1, err := mq.GetClient().QueueSubscribe(mq.SubjectVideo,
		mq.GroupVideo, consumer.HandleVideoMSG)
	if err != nil {
		log.Panicln("QueueSubscribe error:", err)
	}
	qsub2, err := mq.GetClient().QueueSubscribe(mq.SubjectDeviceData, mq.GroupSaveDeviceData,
		consumer.HandleDeviceMsg, stan.DurableName("dur"))
	if err != nil {
		log.Panicln("SubscribeDeviceAttribute error:", err)
	}
	qsub3, err := mq.GetClient().QueueSubscribe(mq.SubjectDeviceData, mq.GroupPublishDeviceData,
		consumer.PublishDeviceData, stan.DurableName("dur"))
	if err != nil {
		log.Panicln("SubscribeDeviceAttribute error:", err)
	}
	return []stan.Subscription{qsub1, qsub2, qsub3}
}

func unSubscribe(qsubs []stan.Subscription) {
	for _, qsub := range qsubs {
		err := qsub.Unsubscribe()
		if err != nil {
			log.Println("qsub Unsubscribe error:", err)
		}
	}
}
