package transaction

import "github.com/stretchr/testify/mock"

type TransactionRepositoryMock struct {
	mock.Mock
}

func (t *TransactionRepositoryMock) FindByID(transactionId string) (*Transaction, error) {
	args := t.Called(transactionId)
	return args.Get(0).(*Transaction), args.Error(1)
}

func (t *TransactionRepositoryMock) Save(transaction Transaction) error {
	args := t.Called(transaction)
	return args.Error(0)
}
