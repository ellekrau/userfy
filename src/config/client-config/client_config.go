package clientconfig

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var clientsConfig Config

func LoadClientsConfig() {
	var err error

	v := viper.New()
	v.SetConfigFile("./../clients.json")

	if err = v.ReadInConfig(); err != nil {
		log.Fatalln(fmt.Sprint("load clients.json config file error: ", err.Error()))
	}

	if err = v.Unmarshal(&clientsConfig); err != nil {
		log.Fatalln(fmt.Sprint("unmarshal clients.json config file error: ", err.Error()))
	}

	for _, client := range clientsConfig.Clients {
		if err = client.DataPattern.User.ValidateUser(); err != nil {
			log.Fatalf("user data pattern error key['%s']: %v", client.Key, err)
		}
	}
}

type Config struct {
	Clients []Client
}

func GetClient(clientKey string) (Client, error) {
	for _, c := range clientsConfig.Clients {
		if c.Key == clientKey {
			return c, nil
		}
	}

	return Client{}, fmt.Errorf("client not found by ['key']: %s", clientKey)
}
