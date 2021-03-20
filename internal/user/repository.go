package user

import (
	"database/sql"
	"log"

	"github.com/seagalputra/cashlog/graph/model"
)

// Repository interface definition for interacting with user table in database.
type Repository interface {
	Save(account *model.User) *model.User
	FindByID(id int64) *model.User
	FindByUsername(username string) *model.User
}

// RepositoryImpl is implementation details for interacting with user table in database.
type RepositoryImpl struct {
	DB *sql.DB
}

// Save is function to storing user data into user table.
func (u *RepositoryImpl) Save(account *model.User) *model.User {
	user := new(model.User)
	query := "INSERT INTO user_account (user_account_id, first_name, last_name, username, password, email, is_disabled, is_verified) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *"
	stmt, err := u.DB.Prepare(query)
	if err != nil {
		log.Print(err)
		return nil
	}
	defer stmt.Close()
	row := stmt.QueryRow(account.UserID, account.FirstName, account.LastName, account.Username, account.Password, account.Email, account.IsDisabled, account.IsVerified)
	err = row.Scan(&user.ID, &user.UserID, &user.FirstName, &user.LastName, &user.Username, &user.Password, &user.Email, &user.IsDisabled, &user.IsVerified)
	if err != nil {
		log.Print(err)
		return nil
	}

	return user
}

// FindByID is function to finding user data based on their row id.
func (u *RepositoryImpl) FindByID(id int64) *model.User {
	user := new(model.User)
	query := "SELECT * FROM user_account WHERE user_account.id = $1"
	stmt, err := u.DB.Prepare(query)
	if err != nil {
		log.Print(err)
		return nil
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(&user.ID, &user.UserID, &user.FirstName, &user.LastName, &user.Username, &user.Password, &user.Email, &user.IsDisabled, &user.IsVerified)
	if err != nil {
		log.Print(err)
		return nil
	}

	return user
}

// FindByUsername is function to finding user data based on their username.
// Username is always unique. So, there aren't duplicate username in database.
func (u *RepositoryImpl) FindByUsername(username string) *model.User {
	user := new(model.User)
	query := "SELECT * FROM user_account WHERE user_account.username = $1"
	stmt, err := u.DB.Prepare(query)
	if err != nil {
		log.Print(err)
		return nil
	}
	defer stmt.Close()
	row := stmt.QueryRow(username)
	err = row.Scan(&user.ID, &user.UserID, &user.FirstName, &user.LastName, &user.Username, &user.Password, &user.Email, &user.IsDisabled, &user.IsVerified)
	if err != nil {
		log.Print(err)
		return nil
	}

	return user
}
