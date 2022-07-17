package config

import (
	"errors"
	namepatternenum "github.com/ellekrau/mercafacil/utils/name-pattern-enum"
	"strings"
)

func validateCellphonePattern() error {
	if UserData.CellphonePattern == "" {
		return nil
	}

	patternCellphoneDigits := strings.Count(UserData.CellphonePattern, "X")
	if patternCellphoneDigits == cellphoneDigitsCount {
		return nil
	}

	return errors.New(errInvalidCellphonePattern)
}

func validateNamePattern() error {
	if UserData.NamePattern == "" {
		return nil
	}

	if namepatternenum.IsNamePatternEnumValue(UserData.NamePattern) {
		return nil
	}

	return errors.New(errInvalidNamePattern)
}
