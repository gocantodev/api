package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"os"
)

type App struct {
	Env      string `yaml:"env"`
	Name     string `yaml:"name"`
	LogLevel string `yaml:"log_level"`
}

type Config struct {
	App App `yaml:"app"`
}

func New() (*Config, error) {
	config := &Config{}
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dir)

	err = cleanenv.ReadConfig(dir+"/config.yml", config)

	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
