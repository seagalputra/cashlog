package transaction

type (
	CreateTransactionRequest struct {
		Title           string `json:"title"`
		Amount          int64  `json:"amount"`
		TransactionDate string `json:"transactionDate"`
		Description     string `json:"description"`
	}
)
