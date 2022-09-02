package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App AppConfig
	DB  DatabaseConfig
}

type DatabaseConfig struct {
	Version     string
	Name        string
	User        string
	Password    string
	Host        string
	Url         string
	MaxLifetime int
	MaxOpenCons int
	MaxIdleCons int
}

type AppConfig struct {
	AppName   string
	LogsLevel string
}

func MakeConfig() (Config, error) {
	var configuration *Config

	viper.SetConfigFile("")
	viper.SetConfigName("gocanto")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("There was an issue reading the config file: %s", err)

		return Config{}, err
	}

	err := viper.Unmarshal(&configuration)

	if err != nil {
		log.Fatalf("There was an issue parsing the config file: %s", err)

		return Config{}, err
	}

	return *configuration, nil
}
