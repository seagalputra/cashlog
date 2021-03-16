package transaction

import "github.com/stretchr/testify/mock"

// RepositoryMock type definition for mocking interaction with database
type RepositoryMock struct {
	mock.Mock
}

// FindByID is mock of FindByID function for finding transaction by row id
func (t *RepositoryMock) FindByID(transactionID string) (*Transaction, error) {
	args := t.Called(transactionID)
	return args.Get(0).(*Transaction), args.Error(1)
}

// Save is mock of Save function to storing transaction data into database
func (t *RepositoryMock) Save(transaction Transaction) error {
	args := t.Called(transaction)
	return args.Error(0)
}

func (t *RepositoryMock) SaveDetail(transactionDetail Detail) (*Detail, error) {
	args := t.Called(transactionDetail)
	return args.Get(0).(*Detail), args.Error(1)
}
