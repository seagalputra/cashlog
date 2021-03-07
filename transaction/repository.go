package transaction

type TransactionRepository interface {
	FindAll() ([]Transaction, error)
	FindByID(transactionId string) (*Transaction, error)
	Save(transaction *Transaction) error
}

type TransactionRepositoryImpl struct{}

func (t *TransactionRepositoryImpl) Save(transaction *Transaction) error {
	panic("implement me")
}

func (t *TransactionRepositoryImpl) FindAll() ([]Transaction, error) {
	panic("implement me")
}

func (t *TransactionRepositoryImpl) FindByID(transactionId string) (*Transaction, error) {
	panic("implement me")
}
