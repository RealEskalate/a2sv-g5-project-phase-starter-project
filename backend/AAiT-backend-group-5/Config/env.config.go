package config

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	JWT_SECRET        string `mapstructure:"JWT_SECRET"`
	SMTP_SERVER       string `mapstructure:"SMTP_SERVER"`
	SMTP_PORT         string `mapstructure:"SMTP_PORT"`
	SMTP_USERNAME     string `mapstructure:"SMTP_USERNAME"`
	SMTP_PASSWORD     string `mapstructure:"SMTP_PASSWORD"`
	SMTP_SENDER_EMAIL string `mapstructure:"SMTP_SENDER_EMAIL"`
	BASE_URL          string `mapstructure:"BASE_URL"`
	REDIS_BLOG_KEY    string `mapstructure:"REDIS_BLOG_KEY"`
}

func NewEnv() *Env {
	viper.AutomaticEnv()

	// Explicitly bind environment variables to the struct fields
	viper.BindEnv("JWT_SECRET")
	viper.BindEnv("SMTP_SERVER")
	viper.BindEnv("SMTP_PORT")
	viper.BindEnv("SMTP_USERNAME")
	viper.BindEnv("SMTP_PASSWORD")
	viper.BindEnv("SMTP_SENDER_EMAIL")
	viper.BindEnv("BASE_URL")
	viper.BindEnv("REDIS_BLOG_KEY")

	env := Env{}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatalf("Error unmarshalling environment variables: %v", err)
	}

	return &env
}
