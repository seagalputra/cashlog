package user

import "github.com/stretchr/testify/mock"

// RepositoryMock type definition of mock user repository
type RepositoryMock struct {
	mock.Mock
}

// FindByUsername function mock for get user from their username
func (u *RepositoryMock) FindByUsername(username string) (*User, error) {
	args := u.Called(username)
	return args.Get(0).(*User), args.Error(1)
}

// FindByID function mock to get user from their row id
func (u *RepositoryMock) FindByID(id int64) (*User, error) {
	args := u.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

// Save function mock to storing user data to database
func (u *RepositoryMock) Save(user *User) error {
	args := u.Called(user)
	return args.Error(0)
}
