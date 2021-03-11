package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib" // only use the driver module
)

// Connect the database using database url in configuration file
func Connect(url string) (*sql.DB, error) {
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, fmt.Errorf("Connect : %v", err)
	}
	return db, nil
}
