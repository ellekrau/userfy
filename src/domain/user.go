package domain

import (
	cellphoneformatter "github.com/ellekrau/mercafacil/utils/cellphone-formatter"
	nameformatter "github.com/ellekrau/mercafacil/utils/name-formatter"
)

type User struct {
	ID        int64
	Name      string
	Cellphone string
}

func NewUser(name, cellphone string) User {
	return User{
		Name:      nameformatter.FormatName(name),
		Cellphone: cellphoneformatter.FormatCellphone(cellphone),
	}
}
