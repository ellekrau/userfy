package userrepository

import (
	"database/sql"
	"github.com/ellekrau/userfy/domain"
)

func createUserPostgres(db *sql.DB, user *domain.User) (err error) {
	err = db.QueryRow("INSERT INTO contacts (nome, celular) VALUES($1, $2) RETURNING id;", user.Name, user.Cellphone).Scan(&user.ID)
	return
}
