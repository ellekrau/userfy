package nameformatter

import (
	"github.com/ellekrau/mercafacil/config/client-config"
	namepatternenum "github.com/ellekrau/mercafacil/utils/name-pattern-enum"
	"strings"
)

func FormatName(clientUserConfig clientconfig.User, inputName string) string {
	// If the NamePatternEnum variable is empty just returns the name input value
	if clientUserConfig.Name == "" {
		return inputName
	}

	switch namepatternenum.NewNamePatternEnum(clientUserConfig.Name) {
	case namepatternenum.UpperCase:
		return strings.ToUpper(inputName)
	case namepatternenum.LowerCase:
		return strings.ToLower(inputName)
	}

	return inputName
}
