package config

import (
	"log"

	"github.com/spf13/viper"
)

func initConfig(fileName string) *viper.Viper {
	config := viper.New()
	config.SetConfigName(fileName)
	config.AddConfigPath(".")
	config.AddConfigPath("$HOME")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("error parsing the config file", err)
	}
	return config
}
