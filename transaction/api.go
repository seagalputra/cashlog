package transaction

import (
	"time"
)

type (
	CreateOutcomeTransactionRequest struct {
		UserID            int64     `json:"userId"`
		Title             string    `json:"title"`
		Amount            string    `json:"amount"`
		TransactionDate   time.Time `json:"date"`
		Description       string    `json:"description"`
		TransactionStatus string    `json:"status"`
		TransactionType   string    `json:"type"`
	}

	CreateIncomeTransactionRequest struct {
		UserID            int64
		Title             string
		Amount            string
		TransactionDate   time.Time
		Description       string
		TransactionStatus string
	}
)
