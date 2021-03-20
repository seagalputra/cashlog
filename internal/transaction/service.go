package transaction

import (
	"fmt"
	"log"
	"strconv"

	uuid "github.com/satori/go.uuid"
	"github.com/seagalputra/cashlog/graph/model"
	"github.com/seagalputra/cashlog/internal/user"
)

// Service interface for transaction
type Service interface {
	CreateTransaction(request TransactionReq) (*model.Transaction, error)
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

type TransactionReq struct {
	UserID string
	Trx model.CreateTransaction
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
		return nil, fmt.Errorf("failed to save outcome transaction")
	}

	return outcome, nil
}

func (t *ServiceImpl) CreateTransaction(request TransactionReq) (*model.Transaction, error) {
	var err error
	id, err := strconv.Atoi(request.UserID)
	if err != nil {
		log.Print(err)
		return &model.Transaction{}, fmt.Errorf("failed to save transaction")
	}
	account := t.UserRepository.FindByID(int64(id))

	trxRequest := request.Trx
	trxDetail := &model.TransactionDetail{
		TransactionDetailID: uuid.NewV4().String(),
		Description:         trxRequest.Description,
		Status:              trxRequest.Status,
	}

	var split *detailAmount
	if trxRequest.Status == model.TransactionStatusIncome {
		split, err = splitAmount(trxRequest.Amount)
		if err != nil {
			return nil, fmt.Errorf("failed to split income transaction")
		}
	}

	if trxRequest.Status == model.TransactionStatusOutcome {
		split, err = getOutcomeAmount(trxRequest.Amount, *trxRequest.Type)
		if err != nil {
			return nil, err
		}
	}

	trxDetail.Needs = split.needs
	trxDetail.Wants = split.wants
	trxDetail.Invest = split.invest

	trx := model.Transaction{
		TransactionID:   uuid.NewV4().String(),
		Title:           trxRequest.Title,
		Amount:          trxRequest.Amount,
		TransactionDate: trxRequest.TransactionDate,
		Detail:          trxDetail,
		User:            account,
	}

	savedTrx := t.TransactionRepository.Save(trx)
	if savedTrx == nil {
		return nil, fmt.Errorf("something wrong when saving transaction data")
	}
	savedTrx.User = account

	return savedTrx, nil
}
