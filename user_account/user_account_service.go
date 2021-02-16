package user_account

type UserAccountService interface {
	RegisterAccount(request *RegisterAccountRequest) RegisterAccountResponse
	AuthenticateUser(request *AuthenticateUserRequest) AuthenticateUserResponse
}

type UserAccountServiceImpl struct {
}

func (u *UserAccountServiceImpl) RegisterAccount(request *RegisterAccountRequest) RegisterAccountResponse {
	panic("implement me!")
}

func (u *UserAccountServiceImpl) AuthenticateUser(request *AuthenticateUserRequest) AuthenticateUserResponse {
	panic("implement me")
}
