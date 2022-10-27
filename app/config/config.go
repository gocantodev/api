package config

type Configuration struct {
	Source File
	App    App `mapstructure:"app"`
	DB     DB  `mapstructure:"database"`
}
