package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func Make() (*Config, error) {
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
