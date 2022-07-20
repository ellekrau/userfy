package domain

import (
	clientconfig "github.com/ellekrau/userfy/config/client-config"
	cellphoneformatter "github.com/ellekrau/userfy/utils/cellphone-formatter"
	nameformatter "github.com/ellekrau/userfy/utils/name-formatter"
)

type User struct {
	ID        int64
	Name      string
	Cellphone string
}

func NewUser(clientUserConfig clientconfig.User, name, cellphone string) User {
	return User{
		Name:      nameformatter.FormatName(clientUserConfig, name),
		Cellphone: cellphoneformatter.FormatCellphone(clientUserConfig, cellphone),
	}
}
