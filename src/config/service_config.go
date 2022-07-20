package config

import (
	"github.com/Netflix/go-env"
	"log"
)

type serviceConfig struct {
	Port string `env:"PORT,required=true"`
}

type authenticationConfig struct {
	JWTKey string `env:"JWT_KEY,required=true"`
}

//type dataSecurityConfig struct {
//	SecurityDataKey string `env:"DATA_SECURITY_KEY,required=true"`
//}

var (
	Service        serviceConfig
	Authentication authenticationConfig
	//DataSecurity   dataSecurityConfig
)

func LoadEnvironmentVariables() {
	// Loads services variables
	if _, err := env.UnmarshalFromEnviron(&Service); err != nil {
		log.Fatal(err)
	}

	// Loads auth variables
	if _, err := env.UnmarshalFromEnviron(&Authentication); err != nil {
		log.Fatal(err)
	}

	// Loads data security variables
	//if err := loadDataSecurityConfig(); err != nil {
	//	log.Fatal(err)
	//}
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
