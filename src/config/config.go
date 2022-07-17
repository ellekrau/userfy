package config

import (
	"errors"
	"github.com/Netflix/go-env"
	"log"
	"strings"
)

const cellphoneDigitsCount = 13
const errInvalidCellphonePattern = "invalid value in 'CELLPHONE PATTERN'"

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
	// Load database config variables
	if _, err := env.UnmarshalFromEnviron(&Database); err != nil {
		log.Fatal(err)
	}

	// Load user data variables
	if _, err := env.UnmarshalFromEnviron(&UserData); err != nil {
		log.Fatal(err)
	}

	// Validates that cellphone pattern is a valid format
	validateCellphonePattern()
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
