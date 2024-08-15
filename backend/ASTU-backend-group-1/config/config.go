package config

import "github.com/spf13/viper"

type Database struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Uri      string `mapstructure:"uri"`
	Name     string `mapstructure:"name"`
}
type Config struct {
	Database Database `mapstructure:"database"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return &Config{}, err
	}
	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		return &config, err
	}
	return &config, nil
}
