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
}

// Envs holds the loaded configuration values.
var Envs = initConfig()

// initConfig initializes the configuration by loading environment variables
// and returns a Config object.
func initConfig() Config {
	if err := godotenv.Load("example.env"); err != nil {
		log.Panicln("Error loading .env file:", err)
	}

	return Config{
		ServerHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		ServerPort:             getEnv("PORT", "8080"),
		DBConnectionString:     getEnv("DB_CONNECTION_STRING", ""),
		DBName:                 getEnv("DB_NAME", "taskdb"),
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JWTExpirationInSeconds: time.Duration(getTimeEnv("JWT_EXPIRATION_IN_SECONDS", 60*24)) * time.Second,
	}
}

// getEnv retrieves the value of an environment variable by key or returns
// a fallback value if the key is not present.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// getTimeEnv retrieves the value of an environment variable as an integer.
// It falls back to a default value if the variable is not set or if there's
// an error in parsing. The result is returned in seconds.
func getTimeEnv(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}