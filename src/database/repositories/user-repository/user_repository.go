package userrepository

import (
	"database/sql"
	"errors"
	clientconfig "github.com/ellekrau/userfy/config/client-config"
	"github.com/ellekrau/userfy/domain"
	"strings"
)

var errInvalidDatabase = "invalid database"

type userRepository struct {
	clientDBConfig clientconfig.Database
	db             *sql.DB
}

func NewUserRepository(clientConfig clientconfig.Client, db *sql.DB) *userRepository {
	return &userRepository{
		db:             db,
		clientDBConfig: clientConfig.Database,
	}
}

func (ur userRepository) CreateUser(user *domain.User) (err error) {
	switch strings.ToLower(ur.clientDBConfig.Database) {
	case "postgres":
		return createUserPostgres(ur.db, user)
	case "mysql":
		return createUserMySQL(ur.db, user)
	}
	return errors.New(errInvalidDatabase)
}
