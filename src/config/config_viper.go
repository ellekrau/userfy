package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func LoadConfig() {
	v := viper.New()
	dir, _ := os.Getwd()
	v.SetConfigFile(dir + "../../config/clients.json")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	var clientsConfig Config
	if err := v.Unmarshal(&clientsConfig); err != nil {
		log.Fatalln(err)
	}
}

type Config struct {
	Clients []Client `json:"clients"`
}

type Client struct {
	Key         string      `json:"key"`
	Name        string      `json:"name"`
	Database    Database    `json:"database"`
	DataPattern DataPattern `json:"dataPattern"`
}

type Database struct {
	Database string `json:"database"`
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type DataPattern struct {
	User User `json:"user"`
}

type User struct {
	Name      string `json:"name"`
	Cellphone string `json:"cellphone"`
}
