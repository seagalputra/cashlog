package transaction

import (
	"database/sql"
)

// Repository is interface definition for transaction repository.
type Repository interface {
	FindByID(transactionID string) (*Transaction, error)
	Save(transaction Transaction) error
	SaveDetail(transactionDetail Detail) (*Detail, error)
}

// RepositoryImpl is implementation details for interacting domain model with database.
type RepositoryImpl struct {
	DB *sql.DB
}

// FindByID querying transaction data based on their row id.
// It's returns single data because every transaction is unique.
func (t *RepositoryImpl) FindByID(transactionID string) (*Transaction, error) {
	panic("implement me")
}

// Save is function to storing transaction data into database.
// Only save one instance of transaction data
func (t *RepositoryImpl) Save(transaction Transaction) error {
	panic("implement me")
}

// SaveDetail will saving the transaction detail and return the Detail with saved id
func (t *RepositoryImpl) SaveDetail(transactionDetail Detail) (*Detail, error) {
	panic("implement me")
}
