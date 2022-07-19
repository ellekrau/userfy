package config

import (
	"errors"
	namepatternenum "github.com/ellekrau/mercafacil/utils/name-pattern-enum"
	"strings"
)

func validateCellphonePattern() error {
	if OldUserData.CellphonePattern == "" {
		return nil
	}

	patternCellphoneDigits := strings.Count(OldUserData.CellphonePattern, "X")
	if patternCellphoneDigits == cellphoneDigitsCount {
		return nil
	}

	return errors.New(errInvalidCellphonePattern)
}

func validateNamePattern() error {
	if OldUserData.NamePattern == "" {
		return nil
	}

	if namepatternenum.IsNamePatternEnumValue(OldUserData.NamePattern) {
		return nil
	}

	return errors.New(errInvalidNamePattern)
}
