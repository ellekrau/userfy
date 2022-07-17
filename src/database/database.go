package database

import (
	"database/sql"
	"errors"
)

func GetDatabase() (*sql.DB, error) {
	return nil, errors.New("db connection error")
}
