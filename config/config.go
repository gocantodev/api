package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"os"
)

type App struct {
	Env      string
	Name     string
	LogLevel string
}

type Config struct {
	App App
}

func New() (*App, error) {
	app := &App{}
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dir)

	err = cleanenv.ReadConfig(dir+"/config.yml", app)

	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(app)

	if err != nil {
		return nil, err
	}

	return app, nil
}
