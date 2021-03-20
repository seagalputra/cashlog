package user

import (
	"testing"

	"github.com/seagalputra/cashlog/graph/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserServiceImpl_RegisterAccount(t *testing.T) {
	repo := &RepositoryMock{}

	userAccountService := &ServiceImpl{repo}

	userRequest := model.RegisterUser{
		FirstName: "Dwiferdio Seagal",
		LastName:  "Putra",
		Username:  "test1234",
		Password:  "12345",
		Email:     "testing1234@email.com",
	}

	savedAccount := &model.User{
		ID:         "1",
		UserID:     "some user id",
		FirstName:  "Dwiferdio Seagal",
		LastName:   "Putra",
		Username:   "test1234",
		Password:   "$2a$10$bhFsfifxr2EPiqoqxwMwt.UGO2CZQZwnVXuE5h7NgY7UEkq3IWUjq",
		Email:      "testing1234@gmail.com",
		IsDisabled: false,
		IsVerified: false,
	}

	repo.On("FindByUsername", mock.Anything).Return(nil)
	repo.On("Save", mock.MatchedBy(func(req *model.User) bool {
		assert.Equal(t, userRequest.Email, req.Email)
		assert.Equal(t, userRequest.FirstName, req.FirstName)
		assert.Equal(t, userRequest.LastName, req.LastName)
		assert.Equal(t, userRequest.Username, req.Username)
		return true
	})).Return(savedAccount)

	response, err := userAccountService.RegisterAccount(userRequest)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	repo.AssertNumberOfCalls(t, "FindByUsername", 1)
	repo.AssertNumberOfCalls(t, "Save", 1)
}

func TestUserServiceImpl_AuthenticateUser(t *testing.T) {
	repo := &RepositoryMock{}

	userAccountService := &ServiceImpl{repo}

	authenticeRequest := model.Login{
		Username: "test1234",
		Password: "12345",
	}

	savedAccount := &model.User{
		ID:         "1",
		UserID:     "some user id",
		FirstName:  "Dwiferdio Seagal",
		LastName:   "Putra",
		Username:   "test1234",
		Password:   "$2a$10$bhFsfifxr2EPiqoqxwMwt.UGO2CZQZwnVXuE5h7NgY7UEkq3IWUjq",
		Email:      "testing1234@gmail.com",
		IsDisabled: false,
		IsVerified: false,
	}

	repo.On("FindByUsername", mock.Anything).Return(savedAccount, nil)

	response, err := userAccountService.Authenticate(authenticeRequest)

	assert.NoError(t, err)
	assert.NotNil(t, response)

	repo.AssertNumberOfCalls(t, "FindByUsername", 1)
}
