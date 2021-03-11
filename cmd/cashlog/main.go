package main

import (
	"os"

	"github.com/seagalputra/cashlog/pkg/config"
	"github.com/seagalputra/cashlog/pkg/db"
	"github.com/seagalputra/cashlog/pkg/server"
	"github.com/seagalputra/cashlog/pkg/transaction"
	"github.com/seagalputra/cashlog/pkg/user"
)

func main() {
	config, err := config.GetConfig(os.Getenv("APP_ENV"), ".")
	if err != nil {
		panic(err)
	}

	conn, err := db.Connect(config.DBUrl)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	userRepo := &user.RepositoryImpl{DB: conn}
	userService := &user.ServiceImpl{UserRepo: userRepo}

	transactionRepo := &transaction.RepositoryImpl{DB: conn}
	transactionService := &transaction.ServiceImpl{TransactionRepository: transactionRepo, UserRepository: userRepo}

	userHandler := user.Handler{UserService: userService}
	transactionhandler := transaction.Handler{TransactionService: transactionService}
	server := server.Server{
		Config:             *config,
		UserHandler:        userHandler,
		TransactionHandler: transactionhandler,
	}

	server.Start()
}
