package config

import (
    "log"
	
    "github.com/spf13/viper"
)

type Env struct {
    JWT_SECRET        string `mapstructure:"JWT_SECRET"`
    SMTP_SERVER       string `mapstructure:"SMTP_SERVER"`
    SMTP_PORT         string `mapstructure:"SMTP_PORT"`
    SMTP_USERNAME     string `mapstructure:"SMTP_USERNAME"`
    SMTP_PASSWORD     string `mapstructure:"SMTP_PASSWORD"`
    SMTP_SENDER_EMAIL string `mapstructure:"SMTP_SENDER_EMAIL"`
    MONGO_URI         string `mapstructure:"MONGO_URI"`
    DB_NAME           string `mapstructure:"DB_NAME"`
}

func NewEnv() *Env {
    viper.SetConfigFile(".env") 
    viper.AutomaticEnv()        

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading .env file: %v", err)
    }

    env := Env{}

    if err := viper.Unmarshal(&env); err != nil {
        log.Fatalf("Error unmarshaling environment variables: %v", err)
    }

    return &env
}
