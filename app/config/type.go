package config

type DB struct {
	ImageVersion string `mapstructure:"image_version"`
	Version      string
	Name         string `mapstructure:"name"`
	User         string
	Password     string
	Host         string
	Url          string
	MaxLifetime  int `mapstructure:"max_lifetime"`
	MaxOpenCons  int `mapstructure:"max_open_cons"`
	MaxIdleCons  int `mapstructure:"max_idle_cons"`
}

type App struct {
	Name string
}

type Logs struct {
	Level string
}

type Configuration struct {
	FilePath File
	EnvPath  File
	App      App  `mapstructure:"app"`
	DB       DB   `mapstructure:"database"`
	Logs     Logs `mapstructure:"logs"`
}

type File struct {
	path      string
	name      string
	extension string
}
