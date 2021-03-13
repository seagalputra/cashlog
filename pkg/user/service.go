package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Service interface definition for dealing with business logic related with user.
type Service interface {
	RegisterAccount(request RegisterRequest) (*RegisterResponse, error)
	Authenticate(request AuthenticateRequest) (*AuthenticateResponse, error)
}

// ServiceImpl is implementation details for dealing with business logic related with user.
type ServiceImpl struct {
	UserRepo Repository
}

// Authenticate is function for authentice user account.
func (u *ServiceImpl) Authenticate(request AuthenticateRequest) (*AuthenticateResponse, error) {
	userAccount, err := u.UserRepo.FindByUsername(request.Username)
	if err != nil {
		return nil, fmt.Errorf("User.Authenticate : %v", err)
	}

	// TODO : Validate username and encoded password
	if userAccount.Username != request.Username || userAccount.Password != request.Password {
		return nil, errors.New("Username or Password is invalid")
	}

	token, err := u.createToken(userAccount.ID)
	if err != nil {
		return nil, fmt.Errorf("User.Authenticate : %v", err)
	}

	response := &AuthenticateResponse{Token: token}

	return response, nil
}

func (u *ServiceImpl) createToken(userID int64) (string, error) {
	// TODO: Change this secret key with key from environment variable
	secret := "asdfghjkl"
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := unsignedToken.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("User.createToken : %v", err)
	}

	return token, nil
}

// RegisterAccount is function for registering given account
func (u *ServiceImpl) RegisterAccount(request RegisterRequest) (*RegisterResponse, error) {
	userID := uuid.NewV4().String()

	_, err := u.UserRepo.FindByUsername(request.Username)
	if err == nil {
		return nil, errors.New("User already exist")
	}

	encrypted, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("User.RegisterAccount : %v", err)
	}

	account := &User{
		UserID:     userID,
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		Username:   request.Username,
		Password:   string(encrypted),
		Email:      request.Email,
		IsDisabled: false,
		IsVerified: false,
	}

	if err := u.UserRepo.Save(account); err != nil {
		return nil, fmt.Errorf("User.RegisterAccount : %v", err)
	}

	return &RegisterResponse{
		UserAccountID: userID,
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		Username:      request.Username,
		Email:         request.Email,
	}, nil
}
