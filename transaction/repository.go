package transaction

import (
	"fmt"

	"github.com/seagalputra/cashlog/common"
)

type TransactionRepository interface {
	FindByID(transactionID string) (*Transaction, error)
	Save(transaction Transaction) error
}

type TransactionRepositoryImpl struct {
}

func (t *TransactionRepositoryImpl) FindByID(transactionID string) (*Transaction, error) {
	panic("implement me")
}

func (t *TransactionRepositoryImpl) Save(transaction Transaction) error {
	db, err := common.Connect()
	if err != nil {
		return fmt.Errorf("Failed to save transaction data : %v ", err)
	}
	defer db.Close()

	query := "INSERT INTO transaction (transaction_id, title, amount, transaction_date, transaction_detail_id, user_account_id) VALUES($1, $2, $3, $4, $5, $6, $7)"

	if _, err := db.Exec(query, transaction.TransactionID, transaction.Title, transaction.Amount, transaction.TransactionDate, transaction.Detail.ID, transaction.UserAccount.ID); err != nil {
		return fmt.Errorf("Failed to save transaction data : %v ", err)
	}

	return nil
}
