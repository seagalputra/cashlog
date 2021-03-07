package transaction

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	. "github.com/seagalputra/cashlog/user_account"
	"time"
)

type TransactionService interface {
	Create(request *CreateTransactionRequest) error
}

type TransactionServiceImpl struct {
	transactionRepository TransactionRepository
	userAccountRepository UserAccountRepository
}

func (t *TransactionServiceImpl) Create(request *CreateTransactionRequest) error {
	account, err := t.userAccountRepository.FindByID(request.UserId)
	if err != nil {
		return fmt.Errorf("User with id %v not found : %v ", request.UserId, err)
	}

	transactionId := uuid.NewV4().String()
	transactionDetailId := uuid.NewV4().String()

	transactionDetail := TransactionDetail{
		TransactionDetailId: transactionDetailId,
		Amount:              request.Amount,
		Description:         request.Description,
		Status:              request.TransactionStatus,
	}

	transaction := &Transaction{
		TransactionId:   transactionId,
		Title:           request.Title,
		Amount:          request.Amount,
		TransactionDate: time.Now(),
		Detail:          transactionDetail,
		UserAccount:     *account,
	}

	err = t.transactionRepository.Save(transaction)
	if err != nil {
		return fmt.Errorf("Failed to save transaction %v ", err)
	}

	return nil
}
