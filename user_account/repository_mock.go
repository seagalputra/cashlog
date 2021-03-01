package user_account

import "github.com/stretchr/testify/mock"

type UserAccountRepositoryMock struct {
	mock.Mock
}

func (u *UserAccountRepositoryMock) FindByUsername(username string) (*UserAccount, error) {
	args := u.Called(username)
	return args.Get(0).(*UserAccount), args.Error(1)
}

func (u *UserAccountRepositoryMock) FindByID(userId string) (*UserAccount, error) {
	args := u.Called(userId)
	return args.Get(0).(*UserAccount), args.Error(1)
}

func (u *UserAccountRepositoryMock) Save(user *UserAccount) error {
	args := u.Called(user)
	return args.Error(0)
}
