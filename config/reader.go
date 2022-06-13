package config

import "github.com/spf13/viper"

type ConfigReader interface {
	config(config Meta) error
	read() error
}

type viperReader struct {
	path     string
	fileName string
}

func (v viperReader) config(config Meta) error {
	path, err := config.GetPath()

	if err != nil {
		return err
	}

	v.path = path
	v.fileName = config.GetName()

	return nil
}

func (v viperReader) read() error {
	viper.AddConfigPath(v.path)
	viper.SetConfigName(v.fileName)

	err := viper.ReadInConfig()

	return err
}
