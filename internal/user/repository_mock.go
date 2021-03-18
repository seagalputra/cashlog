package user

import (
	"github.com/seagalputra/cashlog/graph/model"
	"github.com/stretchr/testify/mock"
)

// RepositoryMock type definition of mock user repository
type RepositoryMock struct {
	mock.Mock
}

// FindByUsername function mock for get user from their username
func (u *RepositoryMock) FindByUsername(username string) (*model.User, error) {
	args := u.Called(username)
	return args.Get(0).(*model.User), args.Error(1)
}

// FindByID function mock to get user from their row id
func (u *RepositoryMock) FindByID(id int64) (*model.User, error) {
	args := u.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

// Save function mock to storing user data to database
func (u *RepositoryMock) Save(user *model.User) error {
	args := u.Called(user)
	return args.Error(0)
}
