package graph

import (
	"github.com/seagalputra/cashlog/internal/transaction"
	"github.com/seagalputra/cashlog/internal/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService        user.Service
	TransactionService transaction.Service
}
