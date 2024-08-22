package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

type Env struct {
	JWT_SECRET        string `mapstructure:"JWT_SECRET"`
	SMTP_SERVER       string `mapstructure:"SMTP_SERVER"`
	SMTP_PORT         string `mapstructure:"SMTP_PORT"`
	SMTP_USERNAME     string `mapstructure:"SMTP_USERNAME"`
	SMTP_PASSWORD     string `mapstructure:"SMTP_PASSWORD"`
	SMTP_SENDER_EMAIL string `mapstructure:"SMTP_SENDER_EMAIL"`
	BASE_URL          string `mapstructure:"BASE_URL"`

	MONGO_URI string `mapstructure:"MONGO_URI"`
	DB_NAME   string `mapstructure:"DB_NAME"`

	ACCESS_TOKEN_EXPIRY_HOUR  int `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	REFRESH_TOKEN_EXPIRY_HOUR int `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`

	SERVER_ADDRESS  string `mapstructure:"SERVER_ADDRESS"`
	CONTEXT_TIMEOUT int    `mapstructure:"CONTEXT_TIMEOUT"`

	GEMINI_API_KEY    string `mapstructure:"GEMINI_API_KEY"`
	REDIS_BLOG_KEY    string `mapstructure:"REDIS_BLOG_KEY"`
	REDIS_DB_ADDRESS  string `mapstructure:"REDIS_DB_ADDRESS"`
	REDIS_DB_PASSWORD string `mapstructure:"REDIS_DB_PASSWORD"`
	REDIS_DB          int    `mapstructure:"REDIS_DB"`
}

func NewEnv() *Env {
	projectRoot := "/home/mercury/Desktop/a2sv_starter_project/a2sv-g5-project-phase-starter-project/backend/AAiT-backend-group-5"

	viper.SetConfigFile(filepath.Join(projectRoot, ".env"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}
	viper.BindEnv("JWT_SECRET")
	viper.BindEnv("SMTP_SERVER")
	viper.BindEnv("SMTP_PORT")
	viper.BindEnv("SMTP_USERNAME")
	viper.BindEnv("SMTP_PASSWORD")
	viper.BindEnv("SMTP_SENDER_EMAIL")
	viper.BindEnv("BASE_URL")
	viper.BindEnv("MONGO_URI")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("SERVER_ADDRESS")
	viper.BindEnv("CONTEXT_TIMEOUT")
	viper.BindEnv("GEMINI_API_KEY")
	viper.BindEnv("REDIS_BLOG_KEY")
	viper.BindEnv("REDIS_DB_ADDRESS")
	viper.BindEnv("REDIS_DB_PASSWORD")
	viper.BindEnv("REDIS_DB")
	viper.BindEnv("ACCESS_TOKEN_EXPIRY_HOUR")
	viper.BindEnv("REFRESH_TOKEN_EXPIRY_HOUR")

	env := Env{}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatalf("Error unmarshaling environment variables: %v", err)
	}

	return &env
}