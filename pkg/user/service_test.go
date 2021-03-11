package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserServiceImpl_RegisterAccount(t *testing.T) {
	repo := &RepositoryMock{}

	userAccountService := &ServiceImpl{repo}

	userRequest := RegisterRequest{
		FirstName: "Dwiferdio Seagal",
		LastName:  "Putra",
		Username:  "test1234",
		Password:  "12345",
		Email:     "testing1234@email.com",
	}

	repo.On("Save", mock.MatchedBy(func(req *User) bool {
		assert.Equal(t, userRequest.Email, req.Email)
		assert.Equal(t, userRequest.FirstName, req.FirstName)
		assert.Equal(t, userRequest.LastName, req.LastName)
		assert.Equal(t, userRequest.Username, req.Username)
		return true
	})).Return(nil)

	userID, err := userAccountService.RegisterAccount(&userRequest)
	assert.NoError(t, err)
	assert.NotNil(t, userID)
}

func TestUserServiceImpl_AuthenticateUser(t *testing.T) {
	repo := &RepositoryMock{}

	userAccountService := &ServiceImpl{repo}

	authenticeRequest := &AuthenticateRequest{
		Username: "test1234",
		Password: "12345",
	}

	savedAccount := &User{
		ID:         1,
		UserID:     "some user id",
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
