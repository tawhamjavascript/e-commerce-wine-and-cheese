package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	USER     string `mapstructure:"USER"`
	PASSWORD string `mapstructure:"PASSWORD"`
	URL      string `mapstructure:"DB_URI"`
}

func LoadConfig() (config *Config) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Error while reading config file %v", err)

	}
	err := viper.Unmarshal(&config)
	if err != nil {
		panic("Error to read file of config")
	}
	return

}
