package user_account

type UserAccountService interface {
	RegisterAccount(request *RegisterAccountRequest) RegisterAccountResponse
	AuthenticateUser(request *AuthenticateUserRequest) AuthenticateUserResponse
}

func (u *UserAccount) RegisterAccount(request *RegisterAccountRequest) RegisterAccountResponse {
	panic("implement me!")
}

func (u *UserAccount) AuthenticateUser(request *AuthenticateUserRequest) AuthenticateUserResponse {
	panic("implement me")
}
