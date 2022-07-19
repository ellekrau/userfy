package nameformatter

import (
	"github.com/ellekrau/mercafacil/config"
	namepatternenum "github.com/ellekrau/mercafacil/utils/name-pattern-enum"
	"strings"
)

func FormatName(inputName string) string {
	// If the NamePatternEnum variable is empty just returns the name input value
	if config.OldUserData.NamePattern == "" {
		return inputName
	}

	switch namepatternenum.NewNamePatternEnum(config.OldUserData.NamePattern) {
	case namepatternenum.UpperCase:
		return strings.ToUpper(inputName)
	case namepatternenum.LowerCase:
		return strings.ToLower(inputName)
	}

	return inputName
}
