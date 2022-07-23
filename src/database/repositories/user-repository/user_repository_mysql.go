package userrepository

import (
	"database/sql"
	clientconfig "github.com/ellekrau/userfy/config/client-config"
	"github.com/ellekrau/userfy/domain"
)

type mySQLUserRepository struct {
	userRepository
}

func newMySQLUserRepository(clientConfig clientconfig.Client, db *sql.DB) *mySQLUserRepository {
	return &mySQLUserRepository{
		userRepository{
			clientDBConfig: clientConfig.Database,
			db:             db,
		},
	}
}

func (mur mySQLUserRepository) CreateUser(user *domain.User) (err error) {
	var statement *sql.Stmt
	if statement, err = mur.db.Prepare("INSERT INTO contacts (nome, celular) VALUES (?, ?)"); err != nil {
		return
	}
	defer statement.Close()

	var result sql.Result
	if result, err = statement.Exec(user.Name, user.Cellphone); err != nil {
		return
	}

	var userID int64
	userID, err = result.LastInsertId()
	if err != nil {
		return
	}

	user.ID = userID
	return
}
