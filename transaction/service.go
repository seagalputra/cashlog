package transaction

import (
	"fmt"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/seagalputra/cashlog/user_account"
)

type TransactionService interface {
	CreateOutcome(request CreateOutcomeTransactionRequest) error
	CreateIncome(request CreateIncomeTransactionRequest) error
	Show(from string, to string) ([]Transaction, error)
}

type TransactionServiceImpl struct {
	TransactionRepository TransactionRepository
	UserAccountRepository user_account.UserAccountRepository
}

const (
	needsRatio  = 0.5
	wantsRatio  = 0.4
	investRatio = 0.1
)

func (t *TransactionServiceImpl) Show(from string, to string) ([]Transaction, error) {
	panic("implement me")
}

func (t *TransactionServiceImpl) CreateIncome(request CreateIncomeTransactionRequest) error {
	account, err := t.UserAccountRepository.FindByID(request.UserID)
	if err != nil {
		return fmt.Errorf("User with id %v not found : %v ", request.UserID, err)
	}

	transactionID := uuid.NewV4().String()
	transactionDetailID := uuid.NewV4().String()
	amount, err := strconv.ParseFloat(request.Amount, 64)
	if err != nil {
		return fmt.Errorf("Failed to save income transaction : %v ", err)
	}

	needs := strconv.FormatFloat(needsRatio*amount, 'f', 2, 64)
	wants := strconv.FormatFloat(wantsRatio*amount, 'f', 2, 64)
	invest := strconv.FormatFloat(investRatio*amount, 'f', 2, 64)

	transactionDetail := &TransactionDetail{
		TransactionDetailID: transactionDetailID,
		Needs:               needs,
		Wants:               wants,
		Invest:              invest,
		Description:         request.Description,
		Status:              request.TransactionStatus,
	}

	transaction := &Transaction{
		TransactionID:   transactionID,
		Title:           request.Title,
		Amount:          request.Amount,
		TransactionDate: time.Now(),
		Detail:          *transactionDetail,
		UserAccount:     *account,
	}

	err = t.TransactionRepository.Save(*transaction)
	if err != nil {
		return fmt.Errorf("Failed to save income transaction : %v ", err)
	}

	return nil
}

func (t *TransactionServiceImpl) CreateOutcome(request CreateOutcomeTransactionRequest) error {
	account, err := t.UserAccountRepository.FindByID(request.UserID)
	if err != nil {
		return fmt.Errorf("User with id %v not found : %v ", request.UserID, err)
	}

	transactionID := uuid.NewV4().String()
	transactionDetailID := uuid.NewV4().String()

	transactionDetail := &TransactionDetail{
		TransactionDetailID: transactionDetailID,
		Description:         request.Description,
		Status:              request.TransactionStatus,
	}

	switch request.TransactionType {
	case "needs":
		transactionDetail.Needs = request.Amount
	case "wants":
		transactionDetail.Wants = request.Amount
	case "invest":
		transactionDetail.Invest = request.Amount
	}

	transaction := &Transaction{
		TransactionID:   transactionID,
		Title:           request.Title,
		Amount:          request.Amount,
		TransactionDate: time.Now(),
		Detail:          *transactionDetail,
		UserAccount:     *account,
	}

	err = t.TransactionRepository.Save(*transaction)
	if err != nil {
		return fmt.Errorf("Failed to save outcome transaction : %v ", err)
	}

	return nil
}
