package config

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

type App struct {
	Env      string
	Name     string
	LogLevel string `mapstructure:"log_level"`
}

type Configuration struct {
	Source File
	App    App `mapstructure:"app"`
	DB     DB  `mapstructure:"database"`
}

type File struct {
	path      string
	name      string
	extension string
}
