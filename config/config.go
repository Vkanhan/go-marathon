package config

import (
	"log"

	"github.com/spf13/viper"
)

// InitConfig lets viper read the config from the toml file
func InitConfig(fileName string) *viper.Viper {
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
