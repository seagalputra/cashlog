package user_account

type UserAccountHandler struct {
	UserAccountService UserAccountService
}

func (h *UserAccountHandler) RegisterAccount(request RegisterAccountRequest) *RegisterAccountResponse {
	panic("Implement me")
}

func (h *UserAccountHandler) Authenticate(request AuthenticateUserRequest) *AuthenticateUserResponse {
	panic("Implement me")
}
