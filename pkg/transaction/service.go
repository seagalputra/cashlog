package transaction

import (
	"fmt"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/seagalputra/cashlog/pkg/user"
)

// Service interface for transaction
type Service interface {
	CreateOutcome(request CreateOutcomeRequest) error
	CreateIncome(request CreateIncomeRequest) error
	Show(from string, to string) ([]Transaction, error)
}

// ServiceImpl implement Service interface
type ServiceImpl struct {
	TransactionRepository Repository
	UserRepository        user.Repository
}

const (
	needsRatio  = 0.5
	wantsRatio  = 0.4
	investRatio = 0.1
)

// Show transactions from specific date.
// Not using pagination because it's too slow when dealing with large data.
// Instead use filtering and infinite scrolling to handle bunch of data.
func (t *ServiceImpl) Show(from string, to string) ([]Transaction, error) {
	panic("implement me")
}

// CreateIncome for creating incoming transaction.
// It calculate using ratio (needs, wants, and invest) and store them into specific column.
func (t *ServiceImpl) CreateIncome(request CreateIncomeRequest) error {
	account, err := t.UserRepository.FindByID(request.UserID)
	if err != nil {
		return fmt.Errorf("Transaction.CreateIncome : %v", err)
	}

	transactionID := uuid.NewV4().String()
	transactionDetailID := uuid.NewV4().String()
	amount, err := strconv.ParseFloat(request.Amount, 64)
	if err != nil {
		return fmt.Errorf("Transaction.CreateIncome : %v", err)
	}

	needs := strconv.FormatFloat(needsRatio*amount, 'f', 2, 64)
	wants := strconv.FormatFloat(wantsRatio*amount, 'f', 2, 64)
	invest := strconv.FormatFloat(investRatio*amount, 'f', 2, 64)

	transactionDetail := &Detail{
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
		User:            *account,
	}

	err = t.TransactionRepository.Save(*transaction)
	if err != nil {
		return fmt.Errorf("Transaction.CreateIncome : %v", err)
	}

	return nil
}

// CreateOutcome for creating outcoming transaction with specific amount.
// Specify the type of transaction because the amount will be stored in different column based on type.
func (t *ServiceImpl) CreateOutcome(request CreateOutcomeRequest) error {
	account, err := t.UserRepository.FindByID(request.UserID)
	if err != nil {
		return fmt.Errorf("Transaction.CreateOutcome : %v", err)
	}

	transactionID := uuid.NewV4().String()
	transactionDetailID := uuid.NewV4().String()

	transactionDetail := &Detail{
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
		User:            *account,
	}

	err = t.TransactionRepository.Save(*transaction)
	if err != nil {
		return fmt.Errorf("Transaction.CreateOutcome : %v", err)
	}

	return nil
}
