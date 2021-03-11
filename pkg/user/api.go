package user

type (
	// RegisterRequest is data contract for registering user account
	RegisterRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Email     string `json:"email"`
	}

	// RegisterResponse is response data when registering user account
	RegisterResponse struct {
		UserAccountID string `json:"id"`
		FirstName     string `json:"first_name"`
		LastName      string `json:"last_name"`
		Username      string `json:"username"`
		Email         string `json:"email"`
	}

	// AuthenticateRequest is data contract for login request
	AuthenticateRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// AuthenticateResponse is response data after successfully login
	AuthenticateResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
)
