package transaction

import (
	. "github.com/seagalputra/cashlog/user_account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/big"
	"testing"
)

func TestTransactionServiceImpl_CreateOutcomeTransaction(t *testing.T) {
	transcationRepository := &TransactionRepositoryMock{}
	userAccountRepository := &UserAccountRepositoryMock{}

	transactionService := &TransactionServiceImpl{transcationRepository, userAccountRepository}

	// Get user id from jwt token
	amount := new(big.Float).SetInt64(15_000)
	request := CreateTransactionRequest{
		Title:             "Lunch",
		Amount:            *amount,
		TransactionDate:   "2021-03-07",
		Description:       "A some description",
		TransactionStatus: "outcome",
	}

	savedAccount := &UserAccount{
		Id:         1,
		UserId:     "some user id",
		FirstName:  "Dwiferdio Seagal",
		LastName:   "Putra",
		Username:   "test1234",
		Password:   "12345",
		Email:      "testing1234@gmail.com",
		IsDisabled: false,
		IsVerified: false,
	}

	userAccountRepository.On("FindByID", mock.Anything).Return(savedAccount, nil)
	transcationRepository.On("Save", mock.MatchedBy(func(req *Transaction) bool {
		assert.Equal(t, request.Title, req.Title)
		return true
	})).Return(nil)

	err := transactionService.Create(&request)
	assert.NoError(t, err)
}
