package user_account

import "github.com/stretchr/testify/mock"

type UserAccountRepositoryMock struct {
	mock.Mock
}

func (u *UserAccountRepositoryMock) Save(user *UserAccount) error {
	args := u.Called(user)
	return args.Error(0)
}
