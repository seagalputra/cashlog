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
func (u *RepositoryMock) FindByUsername(username string) *model.User {
	args := u.Called(username)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*model.User)
}

// FindByID function mock to get user from their row id
func (u *RepositoryMock) FindByID(id int64) *model.User {
	args := u.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*model.User)
}

// Save function mock to storing user data to database
func (u *RepositoryMock) Save(user *model.User) *model.User {
	args := u.Called(user)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*model.User)
}
