package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	JWT_SECRET        string `mapstructure:"JWT_SECRET"`
	SMTP_SERVER       string `mapstructure:"SMTP_SERVER"`
	SMTP_PORT         string `mapstructure:"SMTP_PORT"`
	SMTP_USERNAMR     string `mapstructure:"SMTP_USERNAMR"`
	SMTP_PASSWORD     string `mapstructure:"SMTP_PASSWORD"`
	SMTP_SENDER_EMAIL string `mapstructure:"SMTP_SENDER_EMAIL"`

	SERVER_ADDRESS  string `mapstructure:"SERVER_ADDRESS"`
	CONTEXT_TIMEOUT int    `mapstructure:"CONTEXT_TIMEOUT"`

	ACCESS_TOKEN_EXPIRY_HOUR  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	REFRESH_TOKEN_EXPIRY_HOUR int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	ACCESS_TOKEN_SECRET       string `mapstructure:"ACCESS_TOKEN_SECRET"`
	REFRESH_TOKEN_SECRET      string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() *Env {
	viper.AutomaticEnv()

	env := Env{}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal("Error getting the enviroment variables")
	}

	return &env
}
