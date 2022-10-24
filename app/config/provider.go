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
		target.Source = File{path: file, name: "config", extension: "yaml"}
	}

	if err := target.mapValues(); err != nil {
		return target, err
	}

	return target, nil
}

func (current *Configuration) mapValues() error {
	viper.AddConfigPath(current.Source.path)
	viper.SetConfigName(current.Source.name)
	viper.SetConfigType(current.Source.extension)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(current); err != nil {
		return err
	}

	return nil
}
