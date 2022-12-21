package config

type App struct {
	Env      string `yaml:"env"`
	Name     string `yaml:"name"`
	LogLevel string `yaml:"log_level"`
}

type Database struct {
	ImageVersion string `yaml:"image_version"`
	Version      string `yaml:"version"`
	Name         string `yaml:"name"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Url          string `yaml:"url"`
	MaxLifetime  int    `yaml:"max_lifetime"`
	MaxOpenCons  int    `yaml:"max_open_cons"`
	MaxIdleCons  int    `yaml:"max_idle_cons"`
}

type Config struct {
	App      App      `yaml:"app"`
	Database Database `yaml:"database"`
}
