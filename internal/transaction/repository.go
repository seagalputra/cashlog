package transaction

import (
	"context"
	"database/sql"
	"log"

	"github.com/seagalputra/cashlog/graph/model"
)

// Repository is interface definition for transaction repository.
type Repository interface {
	FindByID(transactionID string) *model.Transaction
	Save(transaction model.Transaction) *model.Transaction
	SaveDetail(transactionDetail model.TransactionDetail) *model.TransactionDetail
}

// RepositoryImpl is implementation details for interacting domain model with database.
type RepositoryImpl struct {
	DB *sql.DB
}

// FindByID querying transaction data based on transaction_id.
// It's returns single data because every transaction is unique.
func (t *RepositoryImpl) FindByID(transactionID string) *model.Transaction {
	transaction := new(model.Transaction)
	query := "SELECT * FROM transaction WHERE transaction.transaction_id = $1"
	err := t.DB.QueryRow(query, transactionID).Scan(&transaction.ID, &transaction.TransactionID, &transaction.Title, &transaction.Amount, &transaction.TransactionDate, &transaction.Detail.ID, &transaction.User.ID)
	if err != nil {
		log.Fatal(err)
	}

	return transaction
}

// Save is function to storing transaction data into database.
// Only save one instance of transaction data
func (t *RepositoryImpl) Save(transaction model.Transaction) *model.Transaction {
	trx := new(model.Transaction)
	query := "INSERT INTO transaction (transaction_id, title, amount, transaction_date, transaction_detail_id, user_account_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *"
	tx, err := t.DB.BeginTx(context.Background(), &sql.TxOptions{})
	//stmt, err := t.DB.Prepare(query)
	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Print(err)
		if err := tx.Rollback(); err != nil {
			log.Print(err)
		}
	}
	defer stmt.Close()

	var trxDetailID string
	var userID string
	row := stmt.QueryRow(transaction.TransactionID, transaction.Title, transaction.Amount, transaction.TransactionDate, transaction.Detail.ID, transaction.User.ID)
	err = row.Scan(&trx.ID, &trx.TransactionID, &trx.Title, &trx.Amount, &trx.TransactionDate, &trxDetailID, &userID)
	if err != nil {
		log.Print(err)
	}

	//_, err := t.DB.QueryRow(query, transaction.TransactionID, transaction.Title, transaction.Amount, transaction.TransactionDate, transaction.Detail.ID, transaction.User.ID)
	//if err != nil {
	//	log.Fatal(err)
	//}

	return nil
}

// SaveDetail will saving the transaction detail and return the Detail with saved id
func (t *RepositoryImpl) SaveDetail(transactionDetail model.TransactionDetail) *model.TransactionDetail {
	detail := new(model.TransactionDetail)
	query := "INSERT INTO transaction_detail (transaction_detail_id, needs, wants, invest, description, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *"
	stmt, err := t.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(transactionDetail.TransactionDetailID, transactionDetail.Needs, transactionDetail.Wants, transactionDetail.Invest, transactionDetail.Description, transactionDetail.Status)
	err = row.Scan(&detail.ID, &detail.TransactionDetailID, &detail.Needs, &detail.Wants, &detail.Invest, &detail.Description, &detail.Status)
	if err != nil {
		log.Fatal(err)
	}

	return detail
}
