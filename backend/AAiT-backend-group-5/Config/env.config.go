package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	JWT_SECRET        string `mapstructure:"JWT_SECRET"`
	SMTP_SERVER       string `mapstructure:"SMTP_SERVER"`
	SMTP_PORT         string `mapstructure:"SMTP_PORT"`
	SMTP_USERNAME     string `mapstructure:"SMTP_USERNAMR"`
	SMTP_PASSWORD     string `mapstructure:"SMTP_PASSWORD"`
	SMTP_SENDER_EMAIL string `mapstructure:"SMTP_SENDER_EMAIL"`
	BASE_URL          string `mapstructure:"BASE_URL"`

	MONGO_URI string `mapstructure:"MONGO_URI"`
	DB_NAME   string `mapstructure:"DB_NAME"`

	SERVER_ADDRESS  string `mapstructure:"SERVER_ADDRESS"`
	CONTEXT_TIMEOUT int    `mapstructure:"CONTEXT_TIMEOUT"`

	ACCESS_TOKEN_EXPIRY_HOUR  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	REFRESH_TOKEN_EXPIRY_HOUR int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	ACCESS_TOKEN_SECRET       string `mapstructure:"ACCESS_TOKEN_SECRET"`
	REFRESH_TOKEN_SECRET      string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() *Env {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}
	// Explicitly bind environment variables to the struct fields
	viper.BindEnv("JWT_SECRET")
	viper.BindEnv("SMTP_SERVER")
	viper.BindEnv("SMTP_PORT")
	viper.BindEnv("SMTP_USERNAME")
	viper.BindEnv("SMTP_PASSWORD")
	viper.BindEnv("SMTP_SENDER_EMAIL")
	viper.BindEnv("BASE_URL")

	env := Env{}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatalf("Error unmarshaling environment variables: %v", err)
		log.Fatalf("Error unmarshalling environment variables: %v", err)
	}

	return &env
}
