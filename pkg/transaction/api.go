package transaction

import (
	"time"
)

type (
	// CreateOutcomeRequest definition for creating outcome transaction
	CreateOutcomeRequest struct {
		UserID            int64     `json:"user_id"`
		Title             string    `json:"title"`
		Amount            string    `json:"amount"`
		TransactionDate   time.Time `json:"transaction_date"`
		Description       string    `json:"description"`
		TransactionStatus string    `json:"status"`
		TransactionType   string    `json:"type"`
	}

	// CreateIncomeRequest definition for creating income transaction
	CreateIncomeRequest struct {
		UserID            int64     `json:"user_id"`
		Title             string    `json:"title"`
		Amount            string    `json:"amount"`
		TransactionDate   time.Time `json:"transaction_date"`
		Description       string    `json:"description"`
		TransactionStatus string    `json:"status"`
	}
)
