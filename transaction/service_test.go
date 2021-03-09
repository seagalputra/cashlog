package transaction

import (
	uuid "github.com/satori/go.uuid"
	. "github.com/seagalputra/cashlog/user_account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/big"
	"testing"
	"time"
)

func TestTransactionServiceImpl_CreateOutcomeTransactionNeeds(t *testing.T) {
	transactionRepository := &TransactionRepositoryMock{}
	userAccountRepository := &UserAccountRepositoryMock{}

	transactionService := &TransactionServiceImpl{transactionRepository, userAccountRepository}

	amount := new(big.Float).SetInt64(-15_000)
	request := CreateOutcomeTransactionRequest{
		UserId:          1,
		Title:           "Lunch",
		Amount:          *amount,
		TransactionDate: time.Now(),
		Description:     "A some description",
		TransactionType: "needs",
	}

	savedAccount := &UserAccount{
		Id:         1,
		UserId:     uuid.NewV4().String(),
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
		assert.Equal(t, request.UserId, req.UserAccount.Id)
		assert.Equal(t, request.Title, req.Title)
		assert.Equal(t, request.Amount, req.Amount)
		assert.Equal(t, request.Description, req.Detail.Description)
		assert.Equal(t, request.Amount, req.Detail.Needs)
		assert.Equal(t, big.Float{}, req.Detail.Wants)
		assert.Equal(t, big.Float{}, req.Detail.Invest)
		return true
	})).Return(nil)

	err := transactionService.CreateOutcome(request)
	assert.NoError(t, err)
}

func TestTransactionServiceImpl_CreateOutcomeTransactionWants(t *testing.T) {
	transactionRepository := &TransactionRepositoryMock{}
	userAccountRepository := &UserAccountRepositoryMock{}

	transactionService := &TransactionServiceImpl{transactionRepository, userAccountRepository}

	amount := new(big.Float).SetInt64(-25_000)
	request := CreateOutcomeTransactionRequest{
		UserId:          1,
		Title:           "Snack",
		Amount:          *amount,
		TransactionDate: time.Now(),
		Description:     "A some description",
		TransactionType: "wants",
	}

	savedAccount := &UserAccount{
		Id:         1,
		UserId:     uuid.NewV4().String(),
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
		assert.Equal(t, request.UserId, req.UserAccount.Id)
		assert.Equal(t, request.Title, req.Title)
		assert.Equal(t, request.Amount, req.Amount)
		assert.Equal(t, request.Description, req.Detail.Description)
		assert.Equal(t, request.Amount, req.Detail.Wants)
		assert.Equal(t, big.Float{}, req.Detail.Needs)
		assert.Equal(t, big.Float{}, req.Detail.Invest)
		return true
	})).Return(nil)

	err := transactionService.CreateOutcome(request)
	assert.NoError(t, err)
}

func TestTransactionServiceImpl_CreateIncomeTransaction(t *testing.T) {
	transactionRepository := &TransactionRepositoryMock{}
	userAccountRepository := &UserAccountRepositoryMock{}

	transactionService := &TransactionServiceImpl{transactionRepository, userAccountRepository}

	amount := new(big.Float).SetInt64(3_000_000)
	request := &CreateIncomeTransactionRequest{
		Title:           "Salary",
		Amount:          *amount,
		TransactionDate: time.Now(),
		Description:     "Monthly salary income",
	}

	savedAccount := &UserAccount{
		Id:         1,
		UserId:     uuid.NewV4().String(),
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

		expectedNeeds := new(big.Float).SetFloat64(1_500_000)
		assert.Equal(t, expectedNeeds.String(), req.Detail.Needs.String())

		expectedWants := new(big.Float).SetFloat64(1_200_000)
		assert.Equal(t, expectedWants.String(), req.Detail.Wants.String())

		expectedInvest := new(big.Float).SetFloat64(300_000)
		assert.Equal(t, expectedInvest.String(), req.Detail.Invest.String())
		return true
	})).Return(nil)

	err := transactionService.CreateIncome(*request)
	assert.NoError(t, err)
}
