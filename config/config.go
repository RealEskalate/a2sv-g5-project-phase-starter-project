package config

import "github.com/spf13/viper"

type Database struct {
	Url      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}
type Email struct {
	Key     string `mapstructure:"key"`
	Address string `mapstructure:"address"`
}
type Jwt struct {
	Secret string `mapstructure:"secret"`
}
type Server struct {
	Port string `mapstructure:"port"`
	Url  string `mapstructure:"url"`
}
type Config struct {
	Database Database `mapstructure:"database"`
	Server   Server   `mapstructure:"server"`
	Email    Email    `mapstructure:"email"`
	Jwt      Jwt      `mapstructure:"jwt"`
}

func LoadConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	var c Config
	viper.Unmarshal(&c)
	return c, nil
}
