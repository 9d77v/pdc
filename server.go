package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/9d77v/pdc/consumers"
	"github.com/9d77v/pdc/graph"
	"github.com/9d77v/pdc/graph/generated"
	"github.com/9d77v/pdc/middleware"

	"github.com/9d77v/pdc/models"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	if models.DEBUG {
		http.Handle("/docs", playground.Handler("GraphQL playground", "/api"))
	}
	http.Handle("/api", middleware.Auth()(srv))

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/build/index.html")
	})
	http.HandleFunc("/app/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/build/index.html")
	})
	http.HandleFunc("/admin/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/build/index.html")
	})
	http.Handle("/", http.FileServer(http.Dir("ui/build")))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	qsub1, _ := models.NatsClient.QueueSubscribe(models.SubjectVideo,
		models.GroupVideo, consumers.HandleVideoMSG)
	defer func() {
		qsub1.Unsubscribe()
		qsub1.Close()
	}()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
