package userrepository

import (
	"database/sql"
	"github.com/ellekrau/mercafacil/domain"
)

func createUserMySQL(db *sql.DB, user *domain.User) (err error) {
	var statement *sql.Stmt
	if statement, err = db.Prepare("INSERT INTO contacts (nome, celular) VALUES (?, ?)"); err != nil {
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
