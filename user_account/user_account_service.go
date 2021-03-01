package user_account

import uuid "github.com/satori/go.uuid"

type UserAccountService interface {
	RegisterAccount(request *RegisterAccountRequest) (string, error)
}

type UserAccountServiceImpl struct {
	userAccountRepo UserAccountRepository
}

func (u *UserAccountServiceImpl) RegisterAccount(request *RegisterAccountRequest) (string, error) {
	userId := uuid.NewV4().String()

	// TODO: encrypt and salt user password
	account := &UserAccount{
		UserId:     userId,
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		Username:   request.Username,
		Password:   request.Password,
		Email:      request.Email,
		IsDisabled: false,
		IsVerified: false,
	}

	if err := u.userAccountRepo.Save(account); err != nil {
		return "", err
	}

	return account.UserId, nil
}
