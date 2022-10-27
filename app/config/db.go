package config

import (
	"os"
)

type DB struct {
	ImageVersion string `mapstructure:"image_version"`
	Version      string
	Name         string
	User         string
	Password     string
	Host         string
	Url          string
	MaxLifetime  int `mapstructure:"max_lifetime"`
	MaxOpenCons  int `mapstructure:"max_open_cons"`
	MaxIdleCons  int `mapstructure:"max_idle_cons"`
}

func (receiver DB) GetImageVersion() string {
	return os.Getenv(receiver.ImageVersion)
}

func (receiver DB) GetVersion() string {
	return os.Getenv(receiver.Version)
}

func (receiver DB) GetName() string {
	return os.Getenv(receiver.Name)
}

func (receiver DB) GetPassword() string {
	return os.Getenv(receiver.Password)
}

func (receiver DB) GetHost() string {
	return os.Getenv(receiver.Host)
}

func (receiver DB) GetUrl() string {
	return os.Getenv(receiver.Url)
}

func (receiver DB) GetMaxLifeTime() int {
	return receiver.MaxLifetime
}

func (receiver DB) GetMaxMaxOpenCons() int {
	return receiver.MaxOpenCons
}

func (receiver DB) GetMaxIdleCons() int {
	return receiver.MaxIdleCons
}
