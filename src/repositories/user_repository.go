package repositories

import (
	"database/sql"
	"github.com/ellekrau/mercafacil/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur userRepository) CreateUser(user *domain.User) (err error) {
	var statement *sql.Stmt
	if statement, err = ur.db.Prepare("INSERT INTO contacts (nome, celular) VALUES (?, ?)"); err != nil {
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
