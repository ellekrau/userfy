package domain

import (
	"github.com/ellekrau/mercafacil/config"
	cellphoneformatter "github.com/ellekrau/mercafacil/utils/cellphone-formatter"
	nameformatter "github.com/ellekrau/mercafacil/utils/name-formatter"
)

type User struct {
	ID        int64
	Name      string
	Cellphone string
}

func NewUser(clientUserConfig config.User, name, cellphone string) User {
	return User{
		Name:      nameformatter.FormatName(clientUserConfig, name),
		Cellphone: cellphoneformatter.FormatCellphone(clientUserConfig, cellphone),
	}
}
