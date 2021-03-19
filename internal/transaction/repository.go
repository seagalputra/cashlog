package transaction

import (
	"context"
	"database/sql"
	"log"

	"github.com/seagalputra/cashlog/graph/model"
)

// Repository is interface definition for transaction repository.
type Repository interface {
	FindByTransactionID(transactionID string) *model.Transaction
	Save(transaction model.Transaction) *model.Transaction
}

// RepositoryImpl is implementation details for interacting domain model with database.
type RepositoryImpl struct {
	DB *sql.DB
}

// FindByID querying transaction data based on transaction_id.
// It's returns single data because every transaction is unique.
func (t *RepositoryImpl) FindByTransactionID(transactionID string) *model.Transaction {

	trx := new(model.Transaction)
	query := "SELECT tr.id, tr.transaction_id, tr.title, tr.amount, tr.transaction_date, td.id, td.transaction_detail_id, td.needs, td.wants, td.invest, td.description, td.status FROM transaction tr LEFT JOIN transaction_detail td ON tr.transaction_detail_id = td.id WHERE tr.transaction_id = $1"
	stmt, err := t.DB.Prepare(query)
	if err != nil {
		log.Print(err)
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(transactionID)
	err = row.Scan(&trx.ID, &trx.TransactionID, &trx.Title, &trx.Amount, &trx.TransactionDate, &trx.Detail.ID, &trx.Detail.TransactionDetailID, &trx.Detail.Needs, &trx.Detail.Wants, &trx.Detail.Invest, &trx.Detail.Description, &trx.Detail.Status)
	if err != nil {
		log.Print(err)
		return nil
	}

	return trx
}

// Save is function to storing transaction data into database.
// Only save one instance of transaction data
func (t *RepositoryImpl) Save(transaction model.Transaction) *model.Transaction {
	trx := new(model.Transaction)

	query := `
		WITH trx_detail AS (
			INSERT INTO transaction_detail (transaction_detail_id, needs, wants, invest, description, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
		)
		INSERT INTO transaction (transaction_id, title, amount, transaction_date, transaction_detail_id, user_account_id) VALUES ($7, $8, $9, $10, (SELECT id FROM trx_detail), $11) RETURNING *
	`
	tx, err := t.DB.BeginTx(context.Background(), &sql.TxOptions{})
	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Print(err)
		if err := tx.Rollback(); err != nil {
			log.Print(err)
			return nil
		}
		return nil
	}
	defer stmt.Close()

	var trxDetailID string
	row := stmt.QueryRow(transaction.Detail.TransactionDetailID, transaction.Detail.Needs, transaction.Detail.Wants, transaction.Detail.Invest, transaction.Detail.Description, transaction.Detail.Status, transaction.TransactionID, transaction.Title, transaction.Amount, transaction.TransactionDate, transaction.User.ID)
	err = row.Scan(&trx.ID, &trx.TransactionID, &trx.Title, &trx.Amount, &trx.TransactionDate, &trxDetailID, &trx.User.ID)
	if err != nil {
		log.Print(err)
		if err := tx.Rollback(); err != nil {
			log.Print(err)
			return nil
		}
		return nil
	}
	detail, err := t.getTransactionDetail(trxDetailID)
	if err != nil {
		log.Print(err)
		if err := tx.Rollback(); err != nil {
			log.Print(err)
			return nil
		}
		return nil
	}
	trx.Detail = detail

	if err := tx.Commit(); err != nil {
		log.Print(err)
		if err := tx.Rollback(); err != nil {
			log.Print(err)
			return nil
		}
		return nil
	}

	return trx
}

func (t *RepositoryImpl) getTransactionDetail(ID string) (*model.TransactionDetail, error) {
	trxDetail := new(model.TransactionDetail)
	query := "SELECT * FROM transaction_detail WHERE id = $1"
	stmt, err := t.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(ID)
	err = row.Scan(&trxDetail.ID, &trxDetail.TransactionDetailID, &trxDetail.Needs, &trxDetail.Wants, &trxDetail.Invest, &trxDetail.Description, &trxDetail.Status)
	if err != nil {
		return nil, err
	}

	return trxDetail, nil
}
