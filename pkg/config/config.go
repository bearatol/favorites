package config

import "github.com/spf13/viper"

const (
	defaultConfigPath     = "./"
	defaultConfigFilename = "config"
	defaultConfigType     = "yml"
)

var conf config

func Conf() *config {
	return &conf
}

func InitConfig(configPath, configFilename, configType string) error {
	if configPath == "" {
		configPath = defaultConfigPath
	}
	if configFilename == "" {
		configFilename = defaultConfigFilename
	}
	if configType == "" {
		configType = defaultConfigType
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configFilename)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&conf); err != nil {
		return err
	}

	return nil
}
