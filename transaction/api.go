package transaction

import "math/big"

type (
	CreateTransactionRequest struct {
		UserId            int64
		Title             string
		Amount            big.Float
		TransactionDate   string
		Description       string
		TransactionStatus string
	}
)
