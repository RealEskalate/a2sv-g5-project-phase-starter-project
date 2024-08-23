// Package config handles application configuration,
// including loading settings from environment variables and configuration files.
package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config holds the application configuration values.
type Config struct {
	ServerHost             string        // Hostname or IP for the server.
	ServerPort             string        // Port number for the server.
	DBName                 string        // Name of the database.
	DBConnectionString     string        // Connection string for the database.
	JWTSecret              string        // Secret key for JWT signing.
	JWTExpirationInSeconds time.Duration // JWT expiration time.
	CacheDB                int           // Cache DB number.
	CacheExpiry            time.Duration // Cache expiry time.
	CacheHost              string        // Cache server host.
	CachePort              string        // Cache server port.
	GoogleApiKey           string        // Google API key.
	MailTrapHost           string        // MailTrap SMTP host.
	MailTrapPort           int           // MailTrap SMTP port.
	MailTrapUsername       string        // MailTrap SMTP username.
	MailTrapPassword       string        // MailTrap SMTP password.
}

// Envs holds the loaded configuration values.
var Envs = initConfig()

// initConfig initializes the configuration by loading environment variables
// and returns a Config object.
func initConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Panicln("Error loading .env file:", err)
	}

	return Config{
		ServerHost:             getEnv("PUBLIC_HOST"),
		ServerPort:             getEnv("PORT"),
		DBConnectionString:     getEnv("DB_CONNECTION_STRING"),
		DBName:                 getEnv("DB_NAME"),
		JWTSecret:              getEnv("JWT_SECRET"),
		JWTExpirationInSeconds: time.Duration(getIntEnv("JWT_EXPIRATION_IN_SECONDS")) * time.Second,
		CacheDB:                getIntEnv("BLOG_CACHE_DB"),
		CacheExpiry:            time.Duration(getIntEnv("BLOG_CACHE_EXPIRY")) * time.Second,
		CacheHost:              getEnv("CACHE_HOST"),
		CachePort:              getEnv("CACHE_PORT"),
		GoogleApiKey:           getEnv("GOOGLE_API_KEY"),
		MailTrapHost:           getEnv("MAILTRAP_HOST"),
		MailTrapPort:           getIntEnv("MAILTRAP_PORT"),
		MailTrapUsername:       getEnv("MAILTRAP_USERNAME"),
		MailTrapPassword:       getEnv("MAILTRAP_PASSWORD"),
	}
}

// getEnv retrieves the value of an environment variable by key.
// If the key is not present, it logs a fatal error and stops the application.
func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return value
}

// getIntEnv retrieves the value of an environment variable as an integer.
// If the key is not present or if there's an error in parsing, it logs a fatal error.
func getIntEnv(key string) int {
	valueStr := getEnv(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("Error parsing environment variable %s as int: %v", key, err)
	}
	return value
}
