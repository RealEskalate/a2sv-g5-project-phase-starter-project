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
	REDIS_BLOG_KEY    string `mapstructure:"REDIS_BLOG_KEY"`

	MONGO_URI string `mapstructure:"MONGO_URI"`
	DB_NAME   string `mapstructure:"DB_NAME"`

	SERVER_ADDRESS  string `mapstructure:"SERVER_ADDRESS"`
	CONTEXT_TIMEOUT int    `mapstructure:"CONTEXT_TIMEOUT"`

	ACCESS_TOKEN_EXPIRY_HOUR  int `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	REFRESH_TOKEN_EXPIRY_HOUR int `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
}

func NewEnv() *Env {
	viper.SetConfigFile("/home/mercury/Desktop/a2sv_starter_project/a2sv-g5-project-phase-starter-project/backend/AAiT-backend-group-5/.env")
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
	viper.BindEnv("REDIS_BLOG_KEY")

	env := Env{}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatalf("Error unmarshaling environment variables: %v", err)
		log.Fatalf("Error unmarshalling environment variables: %v", err)
	}

	return &env
}
