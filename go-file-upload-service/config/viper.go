package config

import (
	"log"

	"github.com/spf13/viper"
)

func NewViperConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("fatal error config file: default \n", err)
	}
}
