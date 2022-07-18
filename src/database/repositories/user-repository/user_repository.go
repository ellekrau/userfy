package userrepository

import (
	"database/sql"
	"errors"
	"github.com/ellekrau/mercafacil/config"
	"github.com/ellekrau/mercafacil/domain"
	"strings"
)

var errInvalidDatabase = "invalid database"

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur userRepository) CreateUser(user *domain.User) (err error) {
	switch strings.ToLower(config.Database.DB) {
	case "postgres":
		return createUserPostgres(ur.db, user)
	case "mysql":
		return createUserMySQL(ur.db, user)
	}
	return errors.New(errInvalidDatabase)
}
