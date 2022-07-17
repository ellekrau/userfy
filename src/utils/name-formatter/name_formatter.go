package nameformatter

import (
	"github.com/ellekrau/mercafacil/config"
	namepatternenum "github.com/ellekrau/mercafacil/utils/name-pattern-enum"
	"strings"
)

func FormatName(inputName string) string {
	// If the NamePatternEnum variable is empty just returns the name input value
	if config.UserData.NamePattern == "" {
		return inputName
	}

	switch namepatternenum.NewNamePatternEnum(config.UserData.NamePattern) {
	case namepatternenum.UpperCase:
		return strings.ToUpper(inputName)
	case namepatternenum.LowerCase:
		return strings.ToLower(inputName)
	}

	return inputName
}
