package user

import (
	"database/sql"
	"fmt"
)

// Repository interface definition for interacting with user table in database.
type Repository interface {
	Save(account *User) error
	FindByID(id int64) (*User, error)
	FindByUsername(username string) (*User, error)
}

// RepositoryImpl is implementation details for interacting with user table in database.
type RepositoryImpl struct {
	DB *sql.DB
}

// Save is function to storing user data into user table.
func (u *RepositoryImpl) Save(account *User) error {
	query := "INSERT INTO user_account (user_account_id, first_name, last_name, username, password, email, is_disabled, is_verified) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err := u.DB.Exec(query, account.UserID, account.FirstName, account.LastName, account.Username, account.Password, account.Email, account.IsDisabled, account.IsVerified)
	if err != nil {
		return fmt.Errorf("UserRepository.Save : %v", err)
	}

	return nil
}

// FindByID is function to finding user data based on their row id.
func (u *RepositoryImpl) FindByID(id int64) (*User, error) {
	savedUser := new(User)
	query := "SELECT * FROM user_account WHERE user_account.id = $1"
	err := u.DB.QueryRow(query, id).Scan(&savedUser.ID, &savedUser.UserID, &savedUser.FirstName, &savedUser.LastName, &savedUser.Username, &savedUser.Password, &savedUser.Email, &savedUser.IsDisabled, &savedUser.IsVerified)
	if err != nil {
		return nil, fmt.Errorf("UserRepository.FindByID : %v", err)
	}

	return savedUser, nil
}

// FindByUsername is function to finding user data based on their username.
// Username is always unique. So, there aren't duplicate username in database.
func (u *RepositoryImpl) FindByUsername(username string) (*User, error) {
	savedUser := new(User)
	query := "SELECT * FROM user_account WHERE user_account.username = $1"
	err := u.DB.QueryRow(query, username).Scan(&savedUser.ID, &savedUser.UserID, &savedUser.FirstName, &savedUser.LastName, &savedUser.Username, &savedUser.Password, &savedUser.Email, &savedUser.IsDisabled, &savedUser.IsVerified)
	if err != nil {
		return nil, fmt.Errorf("UserRepository.FindByUsername : %v", err)
	}

	return savedUser, nil
}
