package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/seagalputra/cashlog/internal/config"
	"github.com/seagalputra/cashlog/internal/db"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/seagalputra/cashlog/graph"
	"github.com/seagalputra/cashlog/graph/generated"
)

const defaultPort = "8080"

func main() {
	err := config.ReadConfig(os.Getenv("APP_ENV"), ".")
	if err != nil {
		panic(err)
	}

	conn, err := db.Connect(config.Get("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	port := config.Get("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
