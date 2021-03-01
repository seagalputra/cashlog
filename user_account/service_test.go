package user_account

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserAccountServiceImpl_RegisterAccount(t *testing.T) {
	repo := &UserAccountRepositoryMock{}

	userAccountService := &UserAccountServiceImpl{repo}

	userRequest := RegisterAccountRequest{
		FirstName: "Dwiferdio Seagal",
		LastName:  "Putra",
		Username:  "test1234",
		Password:  "12345",
		Email:     "testing1234@email.com",
	}

	repo.On("Save", mock.MatchedBy(func(req *UserAccount) bool {
		assert.Equal(t, userRequest.Email, req.Email)
		assert.Equal(t, userRequest.FirstName, req.FirstName)
		assert.Equal(t, userRequest.LastName, req.LastName)
		assert.Equal(t, userRequest.Username, req.Username)
		assert.Equal(t, userRequest.Password, req.Password)
		return true
	})).Return(nil)

	userId, err := userAccountService.RegisterAccount(&userRequest)
	assert.NoError(t, err)
	assert.NotNil(t, userId)
}

func TestUserAccountServiceImpl_AuthenticateUser(t *testing.T) {
	repo := &UserAccountRepositoryMock{}

	userAccountService := &UserAccountServiceImpl{repo}

	authenticeRequest := &AuthenticateUserRequest{
		Username: "test1234",
		Password: "12345",
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

	repo.On("FindByUsername", mock.Anything).Return(savedAccount, nil)

	response, err := userAccountService.Authenticate(authenticeRequest)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}
