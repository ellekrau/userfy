package userrepository

import (
	"database/sql"
	clientconfig "github.com/ellekrau/userfy/config/client-config"
	"github.com/ellekrau/userfy/domain"
)

type postgresUserRepository struct {
	userRepository
}

func newPostgresUserRepository(clientConfig clientconfig.Client, db *sql.DB) *postgresUserRepository {
	return &postgresUserRepository{
		userRepository{
			clientDBConfig: clientConfig.Database,
			db:             db,
		},
	}
}

func (mur postgresUserRepository) CreateUser(user *domain.User) (err error) {
	err = mur.db.QueryRow("INSERT INTO contacts (nome, celular) VALUES($1, $2) RETURNING id;", user.Name, user.Cellphone).Scan(&user.ID)
	return
}
