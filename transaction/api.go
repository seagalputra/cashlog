package transaction

import (
	"math/big"
	"time"
)

type (
	CreateOutcomeTransactionRequest struct {
		UserId            int64
		Title             string
		Amount            big.Float
		TransactionDate   time.Time
		Description       string
		TransactionStatus string
		TransactionType   string // consist of needs, wants or invest
	}

	CreateIncomeTransactionRequest struct {
		UserId            int64
		Title             string
		Amount            big.Float
		TransactionDate   time.Time
		Description       string
		TransactionStatus string
	}
)
