package transaction

type TransactionRepository interface {
	FindByID(transactionId string) (*Transaction, error)
	Save(transaction Transaction) error
}
