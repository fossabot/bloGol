package config

import "github.com/spf13/viper"

type Configuration struct{ *viper.Viper }

var Config *Configuration

func Open(path string) (*Configuration, error) {
	cfg := viper.New()

	cfg.AddConfigPath(path)
	cfg.SetConfigName("development")
	cfg.SetConfigType("yaml")

	if err := cfg.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Configuration{cfg}, nil
}
