package transaction

import (
	"github.com/seagalputra/cashlog/graph/model"
	"github.com/stretchr/testify/mock"
)

// RepositoryMock type definition for mocking interaction with database
type RepositoryMock struct {
	mock.Mock
}

// FindByID is mock of FindByID function for finding transaction by row id
func (t *RepositoryMock) FindByTransactionID(transactionID string) *model.Transaction {
	args := t.Called(transactionID)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*model.Transaction)
}

// Save is mock of Save function to storing transaction data into database
func (t *RepositoryMock) Save(transaction model.Transaction) *model.Transaction {
	args := t.Called(transaction)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*model.Transaction)
}
