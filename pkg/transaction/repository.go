package transaction

import (
	"database/sql"
)

// Repository is interface definition for transaction repository.
type Repository interface {
	FindByID(transactionID string) (*Transaction, error)
	Save(transaction Transaction) error
}

// RepositoryImpl is implementation details for interacting domain model with database.
type RepositoryImpl struct {
	db *sql.DB
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
