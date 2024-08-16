package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// SMTPConfig holds SMTP configuration details
type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

// LoadSMTPConfig loads SMTP configuration from environment or config file
func LoadSMTPConfig() SMTPConfig {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return SMTPConfig{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
		From:     os.Getenv("SMTP_FROM"),
	}
}
