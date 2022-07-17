package config

import (
	"github.com/Netflix/go-env"
	"log"
)

const cellphoneDigitsCount = 13
const errInvalidCellphonePattern = "invalid value in 'CELLPHONE PATTERN'"
const errInvalidNamePattern = "invalid value in 'NAME PATTERN'"

type serviceConfig struct {
	Port string `env:"PORT,required=true"`
}

type authenticationConfig struct {
	ExpirationTimeMinutes int `env:"EXPIRATION_TIME_MINUTES,required=true"`
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
	Service        serviceConfig
	Authentication authenticationConfig
	Database       databaseConfig
	UserData       userDataConfig
)

func LoadEnvironmentVariables() {
	// Loads service variables
	if _, err := env.UnmarshalFromEnviron(&Service); err != nil {
		log.Fatal(err)
	}

	// Loads authentication variables
	if _, err := env.UnmarshalFromEnviron(&Authentication); err != nil {
		log.Fatal(err)
	}

	// Loads database variables
	if _, err := env.UnmarshalFromEnviron(&Database); err != nil {
		log.Fatal(err)
	}

	// Loads user data variables
	if err := loadUserDataConfig(); err != nil {
		log.Fatal(err)
	}
}

func loadUserDataConfig() error {
	if _, err := env.UnmarshalFromEnviron(&UserData); err != nil {
		return err
	}

	if err := validateCellphonePattern(); err != nil {
		return err
	}

	if err := validateNamePattern(); err != nil {
		return err
	}

	return nil
}
