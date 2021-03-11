package user

import (
	"database/sql"
)

// Repository interface definition for interacting with user table in database.
type Repository interface {
	Save(account *User) error
	FindByID(id int64) (*User, error)
	FindByUsername(username string) (*User, error)
}

// RepositoryImpl is implementation details for interacting with user table in database.
type RepositoryImpl struct {
	db *sql.DB
}

// Save is function to storing user data into user table.
func (u *RepositoryImpl) Save(account *User) error {
	panic("implement me")
}

// FindByID is function to finding user data based on their row id.
func (u *RepositoryImpl) FindByID(id int64) (*User, error) {
	panic("implement me")
}

// FindByUsername is function to finding user data based on their username.
// Username is always unique. So, there aren't duplicate username in database.
func (u *RepositoryImpl) FindByUsername(username string) (*User, error) {
	panic("implement me")
}
