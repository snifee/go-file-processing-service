package config

import (
	"log"

	"github.com/spf13/viper"
)

func NewViperConfig() *viper.Viper {

	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatalln("fatal error config file: default \n", err)
	}

	return config
}
