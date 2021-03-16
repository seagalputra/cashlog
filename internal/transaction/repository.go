package transaction

import (
	"database/sql"
	"fmt"
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

// FindByID querying transaction data based on transaction_id.
// It's returns single data because every transaction is unique.
func (t *RepositoryImpl) FindByID(transactionID string) (*Transaction, error) {
	transaction := new(Transaction)
	query := "SELECT * FROM transaction WHERE transaction.transaction_id = $1"
	err := t.DB.QueryRow(query, transactionID).Scan(&transaction.ID, &transaction.TransactionID, &transaction.Title, &transaction.Amount, &transaction.TransactionDate, &transaction.Detail.ID, &transaction.User.ID)
	if err != nil {
		return nil, fmt.Errorf("TransactionRepository.FindByID : %v", err)
	}

	return transaction, nil
}

// Save is function to storing transaction data into database.
// Only save one instance of transaction data
func (t *RepositoryImpl) Save(transaction Transaction) error {
	query := "INSERT INTO transaction (transaction_id, title, amount, transaction_date, transaction_detail_id, user_account_id) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := t.DB.Exec(query, transaction.TransactionID, transaction.Title, transaction.Amount, transaction.TransactionDate, transaction.Detail.ID, transaction.User.ID)
	if err != nil {
		return fmt.Errorf("TransactionRepository.Save : %v", err)
	}

	return nil
}

// SaveDetail will saving the transaction detail and return the Detail with saved id
func (t *RepositoryImpl) SaveDetail(transactionDetail Detail) (*Detail, error) {
	detail := new(Detail)
	query := "INSERT INTO transaction_detail (transaction_detail_id, needs, wants, invest, description, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *"
	err := t.DB.QueryRow(query, transactionDetail.TransactionDetailID, transactionDetail.Needs, transactionDetail.Wants, transactionDetail.Invest, transactionDetail.Description, transactionDetail.Status).Scan(&detail.ID, &detail.TransactionDetailID, &detail.Needs, &detail.Wants, &detail.Invest, &detail.Description, &detail.Status)
	if err != nil {
		return nil, fmt.Errorf("TransactionRepository.SaveDetail : %v", err)
	}

	return detail, nil
}
