package transaction

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	uuid "github.com/satori/go.uuid"
	"github.com/seagalputra/cashlog/graph/model"
	"github.com/seagalputra/cashlog/internal/user"
)

// Service interface for transaction
type Service interface {
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

type detailAmount struct {
	needs  string
	wants  string
	invest string
}

func (t *ServiceImpl) Show(from string, to string) ([]model.Transaction, error) {
	panic("implement me")
}

func splitAmount(amount string) (*detailAmount, error) {
	amt, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	needs := strconv.FormatFloat(amt*needsRatio, 'f', 2, 64)
	wants := strconv.FormatFloat(amt*wantsRatio, 'f', 2, 64)
	invest := strconv.FormatFloat(amt*investRatio, 'f', 2, 64)

	return &detailAmount{
		needs:  needs,
		wants:  wants,
		invest: invest,
	}, nil
}

func getOutcomeAmount(amount string, trxType string) (*detailAmount, error) {
	outcome := &detailAmount{
		needs:  "0.00",
		wants:  "0.00",
		invest: "0.00",
	}

	switch trxType {
	case "needs":
		outcome.needs = amount
	case "wants":
		outcome.wants = amount
	case "invest":
		outcome.invest = amount
	default:
		return nil, errors.New("failed to save outcome transaction")
	}

	return outcome, nil
}

func (t *ServiceImpl) CreateTransaction(request model.CreateTransaction) (*model.Transaction, error) {
	// TODO: Get user account from context
	var id int64 = 1
	account, err := t.UserRepository.FindByID(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("user with id %d not found", id))
	}

	trxDetail := &model.TransactionDetail{
		TransactionDetailID: uuid.NewV4().String(),
		Description:         request.Description,
		Status:              request.Status,
	}

	var split *detailAmount
	if request.Status == model.TransactionStatusIncome {
		split, err = splitAmount(request.Amount)
		if err != nil {
			return nil, errors.New("failed to split income transaction")
		}
	}

	if request.Status == model.TransactionStatusOutcome {
		split, err = getOutcomeAmount(request.Amount, *request.Type)
		if err != nil {
			return nil, err
		}
	}

	trxDetail.Needs = split.needs
	trxDetail.Wants = split.wants
	trxDetail.Invest = split.invest

	trx := model.Transaction{
		TransactionID:   uuid.NewV4().String(),
		Title:           request.Title,
		Amount:          request.Amount,
		TransactionDate: request.TransactionDate,
		Detail:          trxDetail,
		User:            account,
	}

	savedTrx := t.TransactionRepository.Save(trx)
	if savedTrx == nil {
		return nil, errors.New("something wrong when saving transaction data")
	}

	return savedTrx, nil
}
