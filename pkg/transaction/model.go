package transaction

import (
	"time"

	"github.com/seagalputra/cashlog/pkg/user"
)

type (
	// Transaction definition for domain model
	Transaction struct {
		ID              int64
		TransactionID   string
		Title           string
		Amount          string
		TransactionDate time.Time
		Detail          Detail
		User            user.User
	}

	// Detail definition for detail of transaction domain model
	Detail struct {
		ID                  int64
		TransactionDetailID string
		Needs               string
		Wants               string
		Invest              string
		Description         string
		Status              string
	}
)
