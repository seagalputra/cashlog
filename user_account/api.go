package user_account

type (
	RegisterAccountRequest struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Email     string `json:"email"`
	}

	RegisterAccountResponse struct {
		UserAccountId string `json:"id"`
		FirstName     string `json:"firstName"`
		LastName      string `json:"lastName"`
		Username      string `json:"username"`
		Email         string `json:"email"`
	}

	AuthenticateUserRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	AuthenticateUserResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refreshToken"`
	}
)
