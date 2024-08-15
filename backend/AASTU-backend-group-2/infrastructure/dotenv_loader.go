package infrastructure

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// dotenv loader function to extract secret data from .env file
func DotEnvLoader(identifier string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	result, exists := os.LookupEnv(identifier)

	if !exists {
		log.Fatal(".env entry doesn't exist")
	}

	return result
}
