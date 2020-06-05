package main

import (
	"log"
	"net/http"
	"os"

	"git.9d77v.me/9d77v/pdc/graph"
	"git.9d77v.me/9d77v/pdc/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/docs", playground.Handler("GraphQL playground", "/api"))
	http.Handle("/api", srv)
	http.HandleFunc("/app/", handleRedirect)
	http.Handle("/", http.FileServer(http.Dir("ui/build")))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ui/build/index.html")
}
