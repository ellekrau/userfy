package clientconfig

import (
	"fmt"
	namepatternenum "github.com/ellekrau/mercafacil/utils/name-pattern-enum"
	"strings"
)

type DataPattern struct {
	User User `json:"user"`
}

type User struct {
	Name      string `validate:"required"`
	Cellphone string `validate:"required"`
}

func (u User) ValidateUser() error {
	if err := validateUserName(u.Name); err != nil {
		return err
	}

	if err := validateUserCellphone(u.Cellphone); err != nil {
		return err
	}

	return nil
}

func validateUserName(name string) error {
	if name == "" {
		return nil
	}

	if !namepatternenum.IsNamePatternEnumValue(name) {
		return fmt.Errorf("name pattern isn't an accepted value: %s", name)
	}

	return nil
}

func validateUserCellphone(cellphone string) error {
	if cellphone == "" {
		return nil
	}

	if strings.Count(strings.ToUpper(cellphone), "X") != 13 {
		return fmt.Errorf("celphone pattern isn't an accepted value: %s", cellphone)
	}

	return nil
}
