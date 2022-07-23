package userrepository

import (
	"database/sql"
	"errors"
	clientconfig "github.com/ellekrau/userfy/config/client-config"
	"github.com/ellekrau/userfy/domain"
	"strings"
)

var errInvalidDatabase = "invalid database"

type IUserRepository interface {
	CreateUser(user *domain.User) error
}

type userRepository struct {
	clientDBConfig clientconfig.Database
	db             *sql.DB
}

func NewUserRepository(clientConfig clientconfig.Client, db *sql.DB) (IUserRepository, error) {
	switch strings.ToLower(clientConfig.Database.Database) {
	case "mysql":
		return newMySQLUserRepository(clientConfig, db), nil
	case "postgres":
		return newPostgresUserRepository(clientConfig, db), nil
	}
	return nil, errors.New(errInvalidDatabase)
}
