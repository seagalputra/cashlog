package user

import (
	"fmt"
	"github.com/seagalputra/cashlog/graph/model"
	"github.com/seagalputra/cashlog/internal/pkg/token"
	"log"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Service interface definition for dealing with business logic related with user.
type Service interface {
	RegisterAccount(request model.RegisterUser) (*model.AuthPayload, error)
	Authenticate(request model.Login) (*model.AuthPayload, error)
}

// ServiceImpl is implementation details for dealing with business logic related with user.
type ServiceImpl struct {
	UserRepo Repository
}

// Authenticate is function for authentice user account.
func (u *ServiceImpl) Authenticate(request model.Login) (*model.AuthPayload, error) {
	user := u.UserRepo.FindByUsername(request.Username)
	if user == nil {
		return &model.AuthPayload{}, fmt.Errorf("user doesn't exist")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		log.Print(err)
		return &model.AuthPayload{}, fmt.Errorf("wrong password")
	}

	jwt, err := token.Generate(user.ID)
	if err != nil {
		log.Print(err)
		return nil, fmt.Errorf("failed to authenticate account")
	}

	response := &model.AuthPayload{Token: jwt}

	return response, nil
}

// RegisterAccount is function for registering given account
func (u *ServiceImpl) RegisterAccount(request model.RegisterUser) (*model.AuthPayload, error) {
	errMsg := "failed to register account"
	userID := uuid.NewV4().String()

	user := u.UserRepo.FindByUsername(request.Username)
	if user != nil {
		return &model.AuthPayload{}, fmt.Errorf("user already registered")
	}

	encrypted, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
		return &model.AuthPayload{}, fmt.Errorf(errMsg)
	}

	account := &model.User{
		UserID:     userID,
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		Username:   request.Username,
		Password:   string(encrypted),
		Email:      request.Email,
		IsDisabled: false,
		IsVerified: false,
	}

	savedAccount := u.UserRepo.Save(account)
	jwt, err := token.Generate(savedAccount.ID)
	if err != nil {
		log.Print(err)
		return &model.AuthPayload{}, fmt.Errorf(errMsg)
	}

	return &model.AuthPayload{Token: jwt}, nil
}
