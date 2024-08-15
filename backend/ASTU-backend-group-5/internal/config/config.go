package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	MONGO_URI             string
	EMAIL_PROVIDER        string
	SMTP_HOST             string
  	SMTP_PORT             string
  	EMAIL_SENDER_EMAIL    string
  	EMAIL_SENDER_PASSWORD string
  	APP_DOMAIN            string
}

func Load() (*Config, error) {
	err := godotenv.Load("../../.env")
  	if err != nil {
    	return nil, err
  	}
  	Config := &Config{
		MONGO_URI:             os.Getenv("MONGO_URI"),
    	EMAIL_PROVIDER:        os.Getenv("EMAIL_PROVIDER"),
    	SMTP_HOST:             os.Getenv("SMTP_HOST"),
    	SMTP_PORT:             os.Getenv("SMTP_PORT"),
    	EMAIL_SENDER_EMAIL:    os.Getenv("SENDER_EMAIL"),
    	EMAIL_SENDER_PASSWORD: os.Getenv("SENDER_PASSWORD"),
    	APP_DOMAIN:            os.Getenv("APP_DOMAIN"),
  	}
  	return Config, nil
}
