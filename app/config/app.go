package config

import (
	"github.com/gocantodev/api/app/support"
	"os"
)

type App struct {
	Env      string
	Name     string
	LogLevel string `mapstructure:"log_level"`
}

func (receiver App) GetName() string {
	return os.Getenv(receiver.Name)
}

func (receiver App) GetLogLevel() string {
	return os.Getenv(receiver.LogLevel)
}

func (receiver App) GetEnv() string {
	return os.Getenv(receiver.Env)
}

func (receiver App) ToJson() string {
	return support.ToJson(App{
		Name:     receiver.GetName(),
		Env:      receiver.GetEnv(),
		LogLevel: receiver.GetLogLevel(),
	})
}
