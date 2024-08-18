package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var ENV struct {
	DB_ADDRESS             string
	DB_NAME                string
	TEST_DB_NAME           string
	JWT_SECRET_TOKEN       string
	ACCESS_TOKEN_LIFESPAN  string
	REFRESH_TOKEN_LIFESPAN string
	PORT                   int
	ROUTE_PREFIX           string
	ROOT_USERNAME          string
	ROOT_PASSWORD          string
}

/* Loads environment variables from .env file and verifies that all required variables are set */
func LoadEnvironmentVariables() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error: %v", err.Error())
	}

	ENV.DB_ADDRESS = os.Getenv("DB_ADDRESS")
	ENV.DB_NAME = os.Getenv("DB_NAME")
	ENV.TEST_DB_NAME = os.Getenv("TEST_DB_NAME")
	ENV.JWT_SECRET_TOKEN = os.Getenv("JWT_SECRET_TOKEN")
	ENV.ACCESS_TOKEN_LIFESPAN = os.Getenv("ACCESS_TOKEN_LIFESPAN")
	ENV.ROUTE_PREFIX = os.Getenv("ROUTE_PREFIX")
	ENV.REFRESH_TOKEN_LIFESPAN = os.Getenv("REFRESH_TOKEN_LIFESPAN")
	ENV.ROOT_USERNAME = os.Getenv("ROOT_USERNAME")
	ENV.ROOT_PASSWORD = os.Getenv("ROOT_PASSWORD")
	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 64)
	if err != nil {
		return fmt.Errorf("error parsing PORT number: %v", err.Error())
	}

	ENV.PORT = int(port)
	if ENV.DB_ADDRESS == "" {
		return fmt.Errorf("error: couldn't load environment variable 'DB_ADDRESS'")
	}

	if ENV.DB_NAME == "" {
		return fmt.Errorf("error: couldn't load environment variable 'DB_NAME'")
	}

	if ENV.TEST_DB_NAME == "" {
		return fmt.Errorf("error: couldn't load environment variable 'TEST_DB_NAME'")
	}

	if ENV.ACCESS_TOKEN_LIFESPAN == "" {
		return fmt.Errorf("error: couldn't load environment variable 'ACCESS_TOKEN_LIFESPAN'")
	}

	if ENV.REFRESH_TOKEN_LIFESPAN == "" {
		return fmt.Errorf("error: couldn't load environment variable 'REFRESH_TOKEN_LIFESPAN'")
	}

	if ENV.JWT_SECRET_TOKEN == "" {
		return fmt.Errorf("error: couldn't load environment variable 'JWT_SECRET_TOKEN'")
	}

	if ENV.ROOT_USERNAME == "" {
		return fmt.Errorf("error: couldn't load environment variable 'ROOT_USERNAME'")
	}

	if ENV.ROOT_PASSWORD == "" {
		return fmt.Errorf("error: couldn't load environment variable 'ROOT_PASSWORD'")
	}

	if ENV.PORT == 0 {
		return fmt.Errorf("error: couldn't load environment variable 'PORT'")
	}

	return nil
}

/* Removes the root credentials from the environment */
func UnsetRootCredentials() {
	ENV.ROOT_USERNAME = ""
	ENV.ROOT_PASSWORD = ""
	os.Unsetenv("ROOT_USERNAME")
	os.Unsetenv("ROOT_PASSWORD")
}
