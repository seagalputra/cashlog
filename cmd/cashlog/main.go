package main

import (
	"os"

	"github.com/seagalputra/cashlog/internal/pkg/config"
	"github.com/seagalputra/cashlog/internal/pkg/db"
	"github.com/seagalputra/cashlog/internal/server"
	"github.com/seagalputra/cashlog/internal/transaction"
	"github.com/seagalputra/cashlog/internal/user"

	_ "github.com/seagalputra/cashlog/docs"
)

// @title Cashlog App
// @version 1.0
// @description This is an API for Cashlog Application

// @contact.name Dwiferdio Seagal Putra
// @contact.email seagalputra@gmail.com

// @BasePath /api
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

	userRepo := &user.RepositoryImpl{DB: conn}
	userService := &user.ServiceImpl{UserRepo: userRepo}

	transactionRepo := &transaction.RepositoryImpl{DB: conn}
	transactionService := &transaction.ServiceImpl{TransactionRepository: transactionRepo, UserRepository: userRepo}

	userHandler := user.Handler{UserService: userService}
	transactionHandler := transaction.Handler{TransactionService: transactionService}
	srv := server.Server{
		UserHandler:        userHandler,
		TransactionHandler: transactionHandler,
	}

	srv.Start()
}
