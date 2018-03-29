package config

import "github.com/spf13/viper"

var Config = viper.New()

func Open(path string) error {
	Config.AddConfigPath(path)
	Config.SetConfigName("development")
	Config.SetConfigType("yaml")
	return Config.ReadInConfig()
}
