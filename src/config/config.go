package config

import (
	"github.com/Netflix/go-env"
	"log"
)

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
	// Load service variables
	if _, err := env.UnmarshalFromEnviron(&Service); err != nil {
		log.Fatal(err)
	}

	// Load database config variables
	if _, err := env.UnmarshalFromEnviron(&Database); err != nil {
		log.Fatal(err)
	}

	// Load user data variables
	if _, err := env.UnmarshalFromEnviron(&UserData); err != nil {
		log.Fatal(err)
	}
}
