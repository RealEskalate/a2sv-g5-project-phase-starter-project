package env

import (
	"blog_api/domain"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var ENV domain.EnvironmentVariables

// Loads environment variables from .env file and verifies that all required variables are set
func LoadEnvironmentVariables(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		return fmt.Errorf("error: %v", err.Error())
	}

	ENV.DB_ADDRESS = os.Getenv("DB_ADDRESS")
	ENV.DB_NAME = os.Getenv("DB_NAME")
	ENV.TEST_DB_NAME = os.Getenv("TEST_DB_NAME")
	ENV.JWT_SECRET_TOKEN = os.Getenv("JWT_SECRET_TOKEN")
	ENV.ROUTE_PREFIX = os.Getenv("ROUTE_PREFIX")
	ENV.ROOT_USERNAME = os.Getenv("ROOT_USERNAME")
	ENV.ROOT_PASSWORD = os.Getenv("ROOT_PASSWORD")
	ENV.SMTP_GMAIL = os.Getenv("SMTP_GMAIL")
	ENV.SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")
	ENV.REDIS_URL = os.Getenv("REDIS_URL")
	ENV.GOOGLE_CLIENT_ID = os.Getenv("GOOGLE_CLIENT_ID")
	ENV.GOOGLE_CLIENT_SECRET = os.Getenv("GOOGLE_CLIENT_SECRET")
	ENV.GEMINI_API_KEY = os.Getenv(("GEMINI_API_KEY"))

	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 64)
	if err != nil {
		return fmt.Errorf("error parsing PORT number: %v", err.Error())
	}

	accessTkLifespan, err := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_LIFESPAN_MINUTES"), 10, 64)
	if err != nil {
		return fmt.Errorf("error parsing accessTkLifespan number: %v", err.Error())
	}

	refreshTkLifespan, err := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_LIFESPAN_HOURS"), 10, 64)
	if err != nil {
		return fmt.Errorf("error parsing refreshTkLifespan number: %v", err.Error())
	}

	cacheExpiration, err := strconv.ParseInt(os.Getenv("CACHE_EXPIRATION"), 10, 64)
	if err != nil {
		return fmt.Errorf("error parsing cacheExpiration number: %v", err.Error())
	}

	ENV.PORT = int(port)
	ENV.ACCESS_TOKEN_LIFESPAN = int(accessTkLifespan)
	ENV.REFRESH_TOKEN_LIFESPAN = int(refreshTkLifespan)
	ENV.CACHE_EXPIRATION = int(cacheExpiration)

	switch {
	case ENV.DB_ADDRESS == "":
		return fmt.Errorf("error: couldn't load environment variable 'DB_ADDRESS'")
	case ENV.DB_NAME == "":
		return fmt.Errorf("error: couldn't load environment variable 'DB_NAME'")
	case ENV.TEST_DB_NAME == "":
		return fmt.Errorf("error: couldn't load environment variable 'TEST_DB_NAME'")
	case ENV.ROUTE_PREFIX == "":
		return fmt.Errorf("error: couldn't load environment variable 'ROUTE_PREFIX'")
	case ENV.JWT_SECRET_TOKEN == "":
		return fmt.Errorf("error: couldn't load environment variable 'JWT_SECRET_TOKEN'")
	case ENV.ROOT_USERNAME == "":
		return fmt.Errorf("error: couldn't load environment variable 'ROOT_USERNAME'")
	case ENV.ROOT_PASSWORD == "":
		return fmt.Errorf("error: couldn't load environment variable 'ROOT_PASSWORD'")
	case ENV.SMTP_GMAIL == "":
		return fmt.Errorf("error: couldn't load environment variable 'SMTP_GMAIL'")
	case ENV.SMTP_PASSWORD == "":
		return fmt.Errorf("error: couldn't load environment variable 'SMTP_PASSWORD'")
	case ENV.REDIS_URL == "":
		return fmt.Errorf("error: couldn't load environment variable 'REDIS_URL'")
	case ENV.PORT == 0:
		return fmt.Errorf("error: couldn't load environment variable 'PORT'")
	case ENV.GOOGLE_CLIENT_ID == "":
		return fmt.Errorf("error: couldn't load environment variable 'GOOGLE_CLIENT_ID'")
	case ENV.GOOGLE_CLIENT_SECRET == "":
		return fmt.Errorf("error: couldn't load environment variable 'GOOGLE_CLIENT_SECRET'")
	case ENV.GEMINI_API_KEY == "":
		return fmt.Errorf("error: couldn't load environment variable 'GEMINI_API_KEY'")
	case ENV.CACHE_EXPIRATION == 0:
		return fmt.Errorf("error: couldn't load environment variable 'CACHE_EXPIRATION'")
	default:
		return nil
	}
}

// Removes the root credentials from the environment
func UnsetRootCredentials() {
	ENV.ROOT_USERNAME = ""
	ENV.ROOT_PASSWORD = ""
	os.Unsetenv("ROOT_USERNAME")
	os.Unsetenv("ROOT_PASSWORD")
}
