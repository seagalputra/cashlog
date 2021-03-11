package user_account

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type UserAccountService interface {
	RegisterAccount(request *RegisterAccountRequest) (string, error)
	Authenticate(request *AuthenticateUserRequest) (*AuthenticateUserResponse, error)
}

type UserAccountServiceImpl struct {
	userAccountRepo UserAccountRepository
}

func (u *UserAccountServiceImpl) Authenticate(request *AuthenticateUserRequest) (*AuthenticateUserResponse, error) {
	userAccount, err := u.userAccountRepo.FindByUsername(request.Username)
	if err != nil {
		return nil, fmt.Errorf("User with username %s not found : %v ", request.Username, err)
	}

	if userAccount.Username != request.Username || userAccount.Password != request.Password {
		return nil, errors.New("Username or Password is invalid! ")
	}

	token, err := createToken(userAccount.ID)
	if err != nil {
		return nil, fmt.Errorf("Failed to create token : %v ", err)
	}

	response := &AuthenticateUserResponse{Token: token}

	return response, nil
}

func createToken(userID int64) (string, error) {
	// TODO: Change this secret key with key from environment variable
	secret := "asdfghjkl"
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := unsignedToken.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("Failed to sign token : %v ", err)
	}

	return token, nil
}

func (u *UserAccountServiceImpl) RegisterAccount(request *RegisterAccountRequest) (string, error) {
	userID := uuid.NewV4().String()

	// TODO: encrypt and salt user password
	account := &UserAccount{
		UserID:     userID,
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		Username:   request.Username,
		Password:   request.Password,
		Email:      request.Email,
		IsDisabled: false,
		IsVerified: false,
	}

	if err := u.userAccountRepo.Save(account); err != nil {
		return "", fmt.Errorf("Failed to registering account : %v ", err)
	}

	return account.UserID, nil
}
