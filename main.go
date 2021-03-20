package main

import (
	"github.com/seagalputra/cashlog/internal/auth"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/seagalputra/cashlog/internal/pkg/config"
	"github.com/seagalputra/cashlog/internal/pkg/db"
	"github.com/seagalputra/cashlog/internal/transaction"
	"github.com/seagalputra/cashlog/internal/user"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/seagalputra/cashlog/graph"
	"github.com/seagalputra/cashlog/graph/generated"
)

const defaultPort = "8080"

func main() {
	err := config.ReadConfig(os.Getenv("APP_ENV"), ".")
	if err != nil {
		log.Panic(err)
	}

	conn, err := db.Connect(config.Get("DATABASE_URL"))
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	router := chi.NewRouter()
	router.Use(auth.Middleware())

	port := config.Get("PORT")
	if port == "" {
		port = defaultPort
	}

	userRepo := &user.RepositoryImpl{DB: conn}
	userService := &user.ServiceImpl{UserRepo: userRepo}
	transactionRepo := &transaction.RepositoryImpl{DB: conn}
	transactionService := &transaction.ServiceImpl{TransactionRepository: transactionRepo, UserRepository: userRepo}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		UserService:        userService,
		TransactionService: transactionService,
	}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
