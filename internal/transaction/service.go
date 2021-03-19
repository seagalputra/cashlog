package transaction

import (
	"fmt"
	"strconv"

	uuid "github.com/satori/go.uuid"
	"github.com/seagalputra/cashlog/graph/model"
	"github.com/seagalputra/cashlog/internal/user"
)

// Service interface for transaction
type Service interface {
	CreateOutcome(request model.CreateTransaction) error
	CreateIncome(request model.CreateTransaction) error
	CreateWaiting(request model.CreateTransaction) error
	CreateTransaction(request model.CreateTransaction) (*model.Transaction, error)
	Show(from string, to string) ([]model.Transaction, error)
}

// ServiceImpl implement Service interface
type ServiceImpl struct {
	TransactionRepository Repository
	UserRepository        user.Repository
}

// TODO: Create parameterize ratio
const (
	needsRatio  = 0.5
	wantsRatio  = 0.4
	investRatio = 0.1
)

func (t *ServiceImpl) CreateTransaction(request model.CreateTransaction) (*model.Transaction, error) {
	panic("implement me")
}

// Show transactions from specific date.
// Not using pagination because it's too slow when dealing with large data.
// Instead use filtering and infinite scrolling to handle bunch of data.
func (t *ServiceImpl) Show(from string, to string) ([]model.Transaction, error) {
	panic("implement me")
}

// CreateIncome for creating incoming transaction.
// It calculate using ratio (needs, wants, and invest) and store them into specific column.
func (t *ServiceImpl) CreateIncome(request model.CreateTransaction) error {
	// TODO: Get user account from context
	account, err := t.UserRepository.FindByID(1)
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

	transactionDetail := &model.TransactionDetail{
		TransactionDetailID: transactionDetailID,
		Needs:               needs,
		Wants:               wants,
		Invest:              invest,
		Description:         request.Description,
		Status:              request.Status,
	}

	transaction := model.Transaction{
		TransactionID:   transactionID,
		Title:           request.Title,
		Amount:          request.Amount,
		TransactionDate: request.TransactionDate,
		Detail:          transactionDetail,
		User:            account,
	}

	_ = t.TransactionRepository.Save(transaction)

	return nil
}

// CreateOutcome for creating outcoming transaction with specific amount.
// Specify the type of transaction because the amount will be stored in different column based on type.
func (t *ServiceImpl) CreateOutcome(request model.CreateTransaction) error {
	// TODO: Get user account from context
	account, err := t.UserRepository.FindByID(1)
	if err != nil {
		return fmt.Errorf("Transaction.CreateOutcome : %v", err)
	}

	transactionID := uuid.NewV4().String()
	transactionDetailID := uuid.NewV4().String()

	transactionDetail := &model.TransactionDetail{
		TransactionDetailID: transactionDetailID,
		Description:         request.Description,
		Status:              request.Status,
		Needs:               "0.00",
		Wants:               "0.00",
		Invest:              "0.00",
	}

	switch *request.Type {
	case "needs":
		transactionDetail.Needs = request.Amount
	case "wants":
		transactionDetail.Wants = request.Amount
	case "invest":
		transactionDetail.Invest = request.Amount
	}

	transaction := &model.Transaction{
		TransactionID:   transactionID,
		Title:           request.Title,
		Amount:          request.Amount,
		TransactionDate: request.TransactionDate,
		Detail:          transactionDetail,
		User:            account,
	}

	_ = t.TransactionRepository.Save(*transaction)

	return nil
}

func (t *ServiceImpl) CreateWaiting(request model.CreateTransaction) error {
	panic("implement me")
}
