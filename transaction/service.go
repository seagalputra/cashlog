package transaction

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	. "github.com/seagalputra/cashlog/user_account"
	"math/big"
	"time"
)

type TransactionService interface {
	CreateOutcome(request CreateOutcomeTransactionRequest) error
	CreateIncome(request CreateIncomeTransactionRequest) error
	Show(from string, to string) ([]Transaction, error)
}

type TransactionServiceImpl struct {
	transactionRepository TransactionRepository
	userAccountRepository UserAccountRepository
}

const (
	NeedsRatio  = 0.5
	WantsRatio  = 0.4
	InvestRatio = 0.1
)

func (t *TransactionServiceImpl) Show(from string, to string) ([]Transaction, error) {
	panic("implement me")
}

func (t *TransactionServiceImpl) CreateIncome(request CreateIncomeTransactionRequest) error {
	account, err := t.userAccountRepository.FindByID(request.UserId)
	if err != nil {
		return fmt.Errorf("User with id %v not found : %v ", request.UserId, err)
	}

	transactionId := uuid.NewV4().String()
	transactionDetailId := uuid.NewV4().String()

	needs := new(big.Float).Mul(big.NewFloat(NeedsRatio), &request.Amount)
	wants := new(big.Float).Mul(big.NewFloat(WantsRatio), &request.Amount)
	invest := new(big.Float).Mul(big.NewFloat(InvestRatio), &request.Amount)

	transactionDetail := &TransactionDetail{
		TransactionDetailId: transactionDetailId,
		Needs:               *needs,
		Wants:               *wants,
		Invest:              *invest,
		Description:         request.Description,
		Status:              request.TransactionStatus,
	}

	transaction := &Transaction{
		TransactionId:   transactionId,
		Title:           request.Title,
		Amount:          request.Amount,
		TransactionDate: time.Now(),
		Detail:          *transactionDetail,
		UserAccount:     *account,
	}

	err = t.transactionRepository.Save(*transaction)
	if err != nil {
		return fmt.Errorf("Failed to save income transaction : %v ", err)
	}

	return nil
}

func (t *TransactionServiceImpl) CreateOutcome(request CreateOutcomeTransactionRequest) error {
	account, err := t.userAccountRepository.FindByID(request.UserId)
	if err != nil {
		return fmt.Errorf("User with id %v not found : %v ", request.UserId, err)
	}

	transactionId := uuid.NewV4().String()
	transactionDetailId := uuid.NewV4().String()

	transactionDetail := &TransactionDetail{
		TransactionDetailId: transactionDetailId,
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
		TransactionId:   transactionId,
		Title:           request.Title,
		Amount:          request.Amount,
		TransactionDate: time.Now(),
		Detail:          *transactionDetail,
		UserAccount:     *account,
	}

	err = t.transactionRepository.Save(*transaction)
	if err != nil {
		return fmt.Errorf("Failed to save outcome transaction : %v ", err)
	}

	return nil
}
