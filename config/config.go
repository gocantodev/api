package config

type App struct {
	Env      string `yaml:"env" json:"env"`
	Name     string `yaml:"name" json:"name"`
	LogLevel string `yaml:"log_level" yaml:"logLevel"`
}

type Database struct {
	ImageVersion string `yaml:"image_version" json:"imageVersion"`
	Version      string `yaml:"version" json:"version"`
	Name         string `yaml:"name" json:"name"`
	User         string `yaml:"user" json:"user"`
	Password     string `yaml:"password" json:"password"`
	Host         string `yaml:"host" json:"host"`
	Url          string `yaml:"url" json:"url"`
	MaxLifetime  int    `yaml:"max_lifetime" json:"maxLifetime"`
	MaxOpenCons  int    `yaml:"max_open_cons" json:"maxOpenCons"`
	MaxIdleCons  int    `yaml:"max_idle_cons" json:"maxIdleCons"`
}

type Config struct {
	App      App      `yaml:"app" json:"app"`
	Database Database `yaml:"database" json:"database"`
}
