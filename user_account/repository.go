package user_account

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"
)

type UserAccountRepository interface {
	Save(account *UserAccount) error
	FindByID(id int64) (*UserAccount, error)
	FindByUsername(username string) (*UserAccount, error)
}

type UserAccountRepositoryImpl struct {
	db *sql.DB
}

func (u *UserAccountRepositoryImpl) Save(account *UserAccount) error {
	panic("implement me")
}

func (u *UserAccountRepositoryImpl) FindByID(id int64) (*UserAccount, error) {
	return &UserAccount{
		ID:         1,
		UserID:     uuid.NewV4().String(),
		FirstName:  "Dwiferdio Seagal",
		LastName:   "Putra",
		Username:   "test1234",
		Password:   "12345",
		Email:      "testing1234@gmail.com",
		IsDisabled: false,
		IsVerified: false,
	}, nil
}

func (u *UserAccountRepositoryImpl) FindByUsername(username string) (*UserAccount, error) {
	panic("implement me")
}
