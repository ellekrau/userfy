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
	JWTKey string `env:"JWT_KEY,required=true"`
}

type dataSecurityConfig struct {
	SecurityDataKey string `env:"DATA_SECURITY_KEY,required=true"`
}

type databaseConfig struct {
	DB       string `env:"DB,required=true"`
	Host     string `env:"DB_HOST,required=true"`
	Port     string `env:"DB_PORT,required=true"`
	Name     string `env:"DB_NAME,required=true"`
	User     string `env:"DB_USER,required=true"`
	Password string `env:"DB_PASSWORD,required=true"`
}

type userDataConfig struct {
	NamePattern      string `env:"NAME_PATTERN"`
	CellphonePattern string `env:"CELLPHONE_PATTERN"`
}

var (
	Service        serviceConfig
	Authentication authenticationConfig
	DataSecurity   dataSecurityConfig
	Database       databaseConfig
	UserData       userDataConfig
)

func LoadEnvironmentVariables() {
	// Loads services variables
	if _, err := env.UnmarshalFromEnviron(&Service); err != nil {
		log.Fatal(err)
	}

	// Loads authentication variables
	if _, err := env.UnmarshalFromEnviron(&Authentication); err != nil {
		log.Fatal(err)
	}

	// Loads data security variables
	//if err := loadDataSecurityConfig(); err != nil {
	//	log.Fatal(err)
	//}

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

//func loadDataSecurityConfig() error {
//	if _, err := env.UnmarshalFromEnviron(&DataSecurity); err != nil {
//		return err
//	}
//
//	if err := validateDataSecurityKey(); err != nil {
//		return err
//	}
//
//	return nil
//}
