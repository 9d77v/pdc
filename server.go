package main

import (
	"log"
	"net/http"
	"os"

	"git.9d77v.me/9d77v/pdc/graph"
	"git.9d77v.me/9d77v/pdc/graph/generated"
	"git.9d77v.me/9d77v/pdc/middleware"
	"git.9d77v.me/9d77v/pdc/models"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()
	router.Use(middleware.Auth())
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	if models.DEBUG {
		http.Handle("/docs", playground.Handler("GraphQL playground", "/api"))
	}
	router.Handle("/api", srv)
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/build/index.html")
	})
	router.HandleFunc("/app/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/build/index.html")
	})
	router.Handle("/", http.FileServer(http.Dir("ui/build")))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
