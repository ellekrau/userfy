package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type serviceConfig struct {
	Port string `validate:"required"`
}

type authenticationConfig struct {
	JWTKey string `validate:"required"`
}

type dataSecurityConfig struct {
	SecurityDataKey string `validate:"required"`
}

var (
	Service        serviceConfig
	Authentication authenticationConfig
	DataSecurity   dataSecurityConfig
)

func LoadServiceConfig() {
	v := viper.New()
	v.SetConfigFile(".env")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalln(fmt.Sprint("load .env error: ", err.Error()))
	}

	// Loads services variables
	if err := v.Unmarshal(&Service); err != nil {
		log.Fatal(err)
	}

	// Loads auth variables
	if err := v.Unmarshal(&Authentication); err != nil {
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
