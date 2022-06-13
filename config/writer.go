package config

import "github.com/spf13/viper"

type ConfigWriter interface {
	config(config Meta) error
	write() error
}

type viperWriter struct {
	path     string
	fileName string
}

func (v viperWriter) config(config Meta) error {
	path, err := config.GetPath()

	if err != nil {
		return err
	}

	v.path = path
	v.fileName = config.GetName()

	return nil
}

func (v viperWriter) write() error {
	viper.AddConfigPath(v.path)
	viper.SetConfigName(v.fileName)

	err := viper.WriteConfig()

	return err
}
