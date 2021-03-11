package transaction

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/seagalputra/cashlog/pkg/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransactionServiceImpl_CreateOutcomeNeeds(t *testing.T) {
	transactionRepository := &RepositoryMock{}
	userAccountRepository := &user.RepositoryMock{}

	transactionService := &ServiceImpl{transactionRepository, userAccountRepository}

	request := CreateOutcomeRequest{
		UserID:          1,
		Title:           "Lunch",
		Amount:          "-15000.00",
		TransactionDate: time.Now(),
		Description:     "A some description",
		TransactionType: "needs",
	}

	savedAccount := &user.User{
		ID:         1,
		UserID:     uuid.NewV4().String(),
		FirstName:  "Dwiferdio Seagal",
		LastName:   "Putra",
		Username:   "test1234",
		Password:   "12345",
		Email:      "testing1234@gmail.com",
		IsDisabled: false,
		IsVerified: false,
	}

	userAccountRepository.On("FindByID", mock.Anything).Return(savedAccount, nil)
	transactionRepository.On("Save", mock.MatchedBy(func(req Transaction) bool {
		assert.Equal(t, request.UserID, req.User.ID)
		assert.Equal(t, request.Title, req.Title)
		assert.Equal(t, request.Amount, req.Amount)
		assert.Equal(t, request.Description, req.Detail.Description)
		assert.Equal(t, request.Amount, req.Detail.Needs)
		assert.Equal(t, "", req.Detail.Wants)
		assert.Equal(t, "", req.Detail.Invest)
		return true
	})).Return(nil)

	err := transactionService.CreateOutcome(request)
	assert.NoError(t, err)
}

func TestTransactionServiceImpl_CreateOutcomeWants(t *testing.T) {
	transactionRepository := &RepositoryMock{}
	userAccountRepository := &user.RepositoryMock{}

	transactionService := &ServiceImpl{transactionRepository, userAccountRepository}

	request := CreateOutcomeRequest{
		UserID:          1,
		Title:           "Snack",
		Amount:          "-25000.00",
		TransactionDate: time.Now(),
		Description:     "A some description",
		TransactionType: "wants",
	}

	savedAccount := &user.User{
		ID:         1,
		UserID:     uuid.NewV4().String(),
		FirstName:  "Dwiferdio Seagal",
		LastName:   "Putra",
		Username:   "test1234",
		Password:   "12345",
		Email:      "testing1234@gmail.com",
		IsDisabled: false,
		IsVerified: false,
	}

	userAccountRepository.On("FindByID", mock.Anything).Return(savedAccount, nil)
	transactionRepository.On("Save", mock.MatchedBy(func(req Transaction) bool {
		assert.Equal(t, request.UserID, req.User.ID)
		assert.Equal(t, request.Title, req.Title)
		assert.Equal(t, request.Amount, req.Amount)
		assert.Equal(t, request.Description, req.Detail.Description)
		assert.Equal(t, request.Amount, req.Detail.Wants)
		assert.Equal(t, "", req.Detail.Needs)
		assert.Equal(t, "", req.Detail.Invest)
		return true
	})).Return(nil)

	err := transactionService.CreateOutcome(request)
	assert.NoError(t, err)
}

func TestTransactionServiceImpl_CreateIncome(t *testing.T) {
	transactionRepository := &RepositoryMock{}
	userAccountRepository := &user.RepositoryMock{}

	transactionService := &ServiceImpl{transactionRepository, userAccountRepository}

	request := &CreateIncomeRequest{
		Title:           "Salary",
		Amount:          "3000000.00",
		TransactionDate: time.Now(),
		Description:     "Monthly salary income",
	}

	savedAccount := &user.User{
		ID:         1,
		UserID:     uuid.NewV4().String(),
		FirstName:  "Dwiferdio Seagal",
		LastName:   "Putra",
		Username:   "test1234",
		Password:   "12345",
		Email:      "testing1234@gmail.com",
		IsDisabled: false,
		IsVerified: false,
	}

	userAccountRepository.On("FindByID", mock.Anything).Return(savedAccount, nil)
	transactionRepository.On("Save", mock.MatchedBy(func(req Transaction) bool {
		assert.Equal(t, request.Title, req.Title)
		assert.Equal(t, request.Amount, req.Amount)
		assert.Equal(t, request.Description, req.Detail.Description)

		assert.Equal(t, "1500000.00", req.Detail.Needs)
		assert.Equal(t, "1200000.00", req.Detail.Wants)
		assert.Equal(t, "300000.00", req.Detail.Invest)
		return true
	})).Return(nil)

	err := transactionService.CreateIncome(*request)
	assert.NoError(t, err)
}
