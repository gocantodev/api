package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

func Make() (*Config, error) {
	config := &Config{}
	dir, err := os.Getwd()

	if err != nil {
		return config, err
	}

	file := dir + "/config.yml"
	fmt.Printf("Reading configuration from: %s\n", file)

	err = cleanenv.ReadConfig(dir+"/config.yml", config)

	if err != nil {
		return nil, fmt.Errorf("error reading the given configuration file %s config error: %w", file, err)
	}

	err = cleanenv.ReadEnv(config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
