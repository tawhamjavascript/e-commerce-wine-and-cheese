package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	USER     string `mapstructure:"USER"`
	PASSWORD string `mapstructure:"PASSWORD"`
	URL      string `mapstructure:"DB_URI"`
	SecretPassword string `mapstructure:"SECRET_PASSWORD"`
}

var Configs *Config

func LoadConfig() (*Config) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Error while reading config file %v", err)

	}
	err := viper.Unmarshal(&Configs)
	if err != nil {
		panic("Error to read file of config")
	}
	return Configs

}
