package user_account

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type UserAccountRepositoryMock struct {
	mock.Mock
}

func (u *UserAccountRepositoryMock) DoSomething(number int) (bool, error) {
	args := u.Called(number)
	return args.Bool(0), args.Error(1)
}

func TestMock(t *testing.T) {
	testObj := new(UserAccountRepositoryMock)

	testObj.On("DoSomething", 123).Return(true, nil)
}
