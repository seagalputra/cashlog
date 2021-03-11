package main

import (
	"github.com/seagalputra/cashlog/pkg/server"
	"github.com/seagalputra/cashlog/pkg/transaction"
	"github.com/seagalputra/cashlog/pkg/user"
)

func main() {

	userHandler := user.Handler{}
	transactionhandler := transaction.Handler{}
	server := server.Server{
		UserHandler:        userHandler,
		TransactionHandler: transactionhandler,
	}

	server.Start()
}
