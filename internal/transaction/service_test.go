package transaction

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/seagalputra/cashlog/graph/model"
	"github.com/seagalputra/cashlog/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransactionServiceImpl_CreateOutcomeNeeds(t *testing.T) {
	transactionRepository := &RepositoryMock{}
	userAccountRepository := &user.RepositoryMock{}

	transactionService := &ServiceImpl{transactionRepository, userAccountRepository}

	desc := "A some description"
	trxType := "needs"
	request := model.CreateTransaction{
		Title:           "Lunch",
		Amount:          "-15000.00",
		TransactionDate: "2020-11-01",
		Description:     &desc,
		Type:            &trxType,
	}

	savedAccount := &model.User{
		ID:         "1",
		UserID:     uuid.NewV4().String(),
		FirstName:  "Dwiferdio Seagal",
		LastName:   "Putra",
		Username:   "test1234",
		Password:   "12345",
		Email:      "testing1234@gmail.com",
		IsDisabled: false,
		IsVerified: false,
	}

	savedDetail := &model.TransactionDetail{
		ID:                  "1",
		TransactionDetailID: uuid.NewV4().String(),
		Needs:               "-15000.00",
		Wants:               "0.00",
		Invest:              "0.00",
		Description:         &desc,
		Status:              "outcome",
	}

	savedTrx := &model.Transaction{
		ID:              "1",
		TransactionID:   uuid.NewV4().String(),
		Title:           "Lunch",
		Amount:          "-15000.00",
		TransactionDate: "2020-11-01",
		Detail:          savedDetail,
		User:            savedAccount,
	}

	userAccountRepository.On("FindByID", mock.Anything).Return(savedAccount, nil)
	transactionRepository.On("Save", mock.MatchedBy(func(req model.Transaction) bool {
		assert.Equal(t, request.Title, req.Title)
		assert.Equal(t, request.Amount, req.Amount)
		assert.Equal(t, request.Description, req.Detail.Description)
		assert.Equal(t, request.Amount, req.Detail.Needs)
		assert.Equal(t, "0.00", req.Detail.Wants)
		assert.Equal(t, "0.00", req.Detail.Invest)
		return true
	})).Return(savedTrx)

	err := transactionService.CreateOutcome(request)
	assert.NoError(t, err)

	userAccountRepository.AssertNumberOfCalls(t, "FindByID", 1)
	transactionRepository.AssertNumberOfCalls(t, "Save", 1)
}

func TestTransactionServiceImpl_CreateOutcomeWants(t *testing.T) {
	transactionRepository := &RepositoryMock{}
	userAccountRepository := &user.RepositoryMock{}

	transactionService := &ServiceImpl{transactionRepository, userAccountRepository}

	desc := "A some description"
	trxType := "wants"
	request := model.CreateTransaction{
		Title:           "Snack",
		Amount:          "-25000.00",
		TransactionDate: "2020-11-01",
		Description:     &desc,
		Type:            &trxType,
	}

	savedAccount := &model.User{
		ID:         "1",
		UserID:     uuid.NewV4().String(),
		FirstName:  "Dwiferdio Seagal",
		LastName:   "Putra",
		Username:   "test1234",
		Password:   "12345",
		Email:      "testing1234@gmail.com",
		IsDisabled: false,
		IsVerified: false,
	}

	savedDetail := &model.TransactionDetail{
		ID:                  "1",
		TransactionDetailID: uuid.NewV4().String(),
		Needs:               "0.00",
		Wants:               "-25000.00",
		Invest:              "0.00",
		Description:         &desc,
		Status:              "outcome",
	}

	savedTrx := &model.Transaction{
		ID:              "1",
		TransactionID:   uuid.NewV4().String(),
		Title:           "Snack",
		Amount:          "-25000.00",
		TransactionDate: "2020-11-01",
		Detail:          savedDetail,
		User:            savedAccount,
	}

	userAccountRepository.On("FindByID", mock.Anything).Return(savedAccount, nil)
	transactionRepository.On("Save", mock.MatchedBy(func(req model.Transaction) bool {
		assert.Equal(t, request.Title, req.Title)
		assert.Equal(t, request.Amount, req.Amount)
		assert.Equal(t, request.Description, req.Detail.Description)
		assert.Equal(t, request.Amount, req.Detail.Wants)
		assert.Equal(t, "0.00", req.Detail.Needs)
		assert.Equal(t, "0.00", req.Detail.Invest)
		return true
	})).Return(savedTrx)

	err := transactionService.CreateOutcome(request)
	assert.NoError(t, err)

	userAccountRepository.AssertNumberOfCalls(t, "FindByID", 1)
	transactionRepository.AssertNumberOfCalls(t, "Save", 1)
}

func TestTransactionServiceImpl_CreateIncome(t *testing.T) {
	transactionRepository := &RepositoryMock{}
	userAccountRepository := &user.RepositoryMock{}

	transactionService := &ServiceImpl{transactionRepository, userAccountRepository}

	desc := "Monthly salary income"
	request := &model.CreateTransaction{
		Title:           "Salary",
		Amount:          "3000000.00",
		TransactionDate: "2020-11-01",
		Description:     &desc,
	}

	savedAccount := &model.User{
		ID:         "1",
		UserID:     uuid.NewV4().String(),
		FirstName:  "Dwiferdio Seagal",
		LastName:   "Putra",
		Username:   "test1234",
		Password:   "12345",
		Email:      "testing1234@gmail.com",
		IsDisabled: false,
		IsVerified: false,
	}

	savedDetail := &model.TransactionDetail{
		ID:                  "1",
		TransactionDetailID: uuid.NewV4().String(),
		Needs:               "1500000.00",
		Wants:               "1200000.00",
		Invest:              "300000.00",
		Description:         &desc,
		Status:              "income",
	}

	savedTrx := &model.Transaction{
		ID:              "1",
		TransactionID:   uuid.NewV4().String(),
		Title:           "Salary",
		Amount:          "3000000.00",
		TransactionDate: "2020-11-01",
		Detail:          savedDetail,
		User:            savedAccount,
	}

	userAccountRepository.On("FindByID", mock.Anything).Return(savedAccount, nil)
	transactionRepository.On("Save", mock.MatchedBy(func(req model.Transaction) bool {
		assert.Equal(t, request.Title, req.Title)
		assert.Equal(t, request.Amount, req.Amount)
		assert.Equal(t, request.Description, req.Detail.Description)
		return true
	})).Return(savedTrx)

	err := transactionService.CreateIncome(*request)

	assert.NoError(t, err)

	userAccountRepository.AssertNumberOfCalls(t, "FindByID", 1)
	transactionRepository.AssertNumberOfCalls(t, "Save", 1)
}
