package transaction

import (
	"time"

	"github.com/seagalputra/cashlog/user_account"
)

type (
	Transaction struct {
		ID              int64
		TransactionID   string
		Title           string
		Amount          string
		TransactionDate time.Time
		Detail          TransactionDetail
		UserAccount     user_account.UserAccount
	}

	TransactionDetail struct {
		ID                  int64
		TransactionDetailID string
		Needs               string
		Wants               string
		Invest              string
		Description         string
		Status              string
	}
)
