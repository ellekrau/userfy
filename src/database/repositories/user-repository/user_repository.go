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
	clientDBConfig config.Database
	db             *sql.DB
}

func NewUserRepository(clientConfig config.Client, db *sql.DB) *userRepository {
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
