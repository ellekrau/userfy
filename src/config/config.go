package config

import (
	"errors"
	"github.com/Netflix/go-env"
	namepatternenum "github.com/ellekrau/mercafacil/utils/name-pattern-enum"
	"log"
	"strings"
)

const cellphoneDigitsCount = 13
const errInvalidCellphonePattern = "invalid value in 'CELLPHONE PATTERN'"
const errInvalidNamePattern = "invalid value in 'NAME PATTERN'"

type serviceConfig struct {
	Port string `env:"PORT,required=true"`
}

type databaseConfig struct {
	Port    int32  `env:"DB_PORT,required=true"`
	Address string `env:"DB_ADDRESS,required=true"`
}

type userDataConfig struct {
	NamePattern      string `env:"NAME_PATTERN"`
	CellphonePattern string `env:"CELLPHONE_PATTERN"`
}

var (
	Service  serviceConfig
	Database databaseConfig
	UserData userDataConfig
)

func LoadEnvironmentVariables() {
	// Loads service variables
	if _, err := env.UnmarshalFromEnviron(&Service); err != nil {
		log.Fatal(err)
	}

	// Loads database variables
	if _, err := env.UnmarshalFromEnviron(&Database); err != nil {
		log.Fatal(err)
	}

	// Loads user data variables
	if _, err := env.UnmarshalFromEnviron(&UserData); err != nil {
		log.Fatal(err)
	}

	// Validates if cellphone pattern is a valid format
	if err := validateCellphonePattern(); err != nil {
		log.Fatal(err)
	}

	// Validates if NamePatternEnum is an expected value
	if err := validateNamePattern(); err != nil {
		log.Fatal(err)
	}
}

func validateCellphonePattern() error {
	if UserData.CellphonePattern == "" {
		return nil
	}

	patternCellphoneDigits := strings.Count(UserData.CellphonePattern, "0")
	if patternCellphoneDigits == cellphoneDigitsCount {
		return nil
	}

	return errors.New(errInvalidCellphonePattern)
}

func validateNamePattern() error {
	if namepatternenum.IsNamePatternEnumValue(UserData.NamePattern) {
		return nil
	}
	return errors.New(errInvalidNamePattern)
}
