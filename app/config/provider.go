package config

import (
	"github.com/spf13/viper"
	"path/filepath"
)

func Make(source string) (Configuration, error) {
	target := Configuration{}

	if file, err := filepath.Abs(source); err != nil {
		return target, err
	} else {
		target.FilePath = File{path: file, name: "config", extension: "yaml"}
	}

	if err := target.mapValues(); err != nil {
		return target, err
	}

	return target, nil
}

func (current *Configuration) mapValues() error {
	viper.AddConfigPath(current.FilePath.path)
	viper.SetConfigName(current.FilePath.name)
	viper.SetConfigType(current.FilePath.extension)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(current); err != nil {
		return err
	}

	return nil
}
