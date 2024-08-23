package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Loads environment variables from .env file and verifies that all required variables are set

// Config struct to hold all environment variables
type EnvironmentVariables struct {
	DB_URI                  string
	PORT                    string
	DB_NAME                 string
	USER_COLLECTION         string
	BLOG_COLLECTION         string
	COMMENT_COLLECTION_NAME string
	LIKE_COLLECTION_NAME    string
	JWT_SECRET              string
	EMAIL_HOST              string
	EMAIL_USERNAME          string
	EMAIL_PASSWORD          string
	EMAIL                   string
	GEMINI_API_KEY          string
	REDIS_ADDR              string
	REDIS_PASSWORD          string
	GOOGLE_CLIENT_ID		string
	GOOGLE_CLIENT_SECRET	string
}

var ENV EnvironmentVariables

func LoadEnvironmentVariables() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error: %v", err.Error())
	}

	// set ENV fields from env
	ENV.DB_URI = os.Getenv("DB_URI")
	ENV.PORT = os.Getenv("PORT")
	ENV.DB_NAME = os.Getenv("DB_NAME")
	ENV.USER_COLLECTION = os.Getenv("USER_COLLECTION")
	ENV.BLOG_COLLECTION = os.Getenv("BLOG_COLLECTION")
	ENV.COMMENT_COLLECTION_NAME = os.Getenv("COMMENT_COLLECTION_NAME")
	ENV.LIKE_COLLECTION_NAME = os.Getenv("LIKE_COLLECTION_NAME")
	ENV.JWT_SECRET = os.Getenv("JWT_SECRET")
	ENV.EMAIL_HOST = os.Getenv("EMAIL_HOST")
	ENV.EMAIL_USERNAME = os.Getenv("EMAIL_USERNAME")
	ENV.EMAIL_PASSWORD = os.Getenv("EMAIL_PASSWORD")
	ENV.EMAIL = os.Getenv("EMAIL")
	ENV.GEMINI_API_KEY = os.Getenv("GEMINI_API_KEY")
	ENV.REDIS_ADDR = os.Getenv("REDIS_ADDR")
	ENV.REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	ENV.GOOGLE_CLIENT_ID = os.Getenv("GOOGLE_CLIENT_ID")
	ENV.GOOGLE_CLIENT_SECRET = os.Getenv("GOOGLE_CLIENT_SECRET")

	//Switch through the fields and make sure they are not empty
	switch {
		case ENV.DB_URI == "":
			return fmt.Errorf("error: couldn't load environment variable 'DB_URI'")
		case ENV.PORT == "":
			return fmt.Errorf("error: couldn't load environment variable 'PORT'")
		case ENV.DB_NAME == "":
			return fmt.Errorf("error: couldn't load environment variable 'DB_NAME'")
		case ENV.USER_COLLECTION == "":
			return fmt.Errorf("error: couldn't load environment variable 'USER_COLLECTION'")
		case ENV.BLOG_COLLECTION == "":
			return fmt.Errorf("error: couldn't load environment variable 'BLOG_COLLECTION'")
		case ENV.COMMENT_COLLECTION_NAME == "":
			return fmt.Errorf("error: couldn't load environment variable 'COMMENT_COLLECTION_NAME'")
		case ENV.LIKE_COLLECTION_NAME == "":
			return fmt.Errorf("error: couldn't load environment variable 'LIKE_COLLECTION_NAME'")
		case ENV.JWT_SECRET == "":
			return fmt.Errorf("error: couldn't load environment variable 'JWT_SECRET'")
		case ENV.EMAIL_HOST == "":
			return fmt.Errorf("error: couldn't load environment variable 'EMAIL_HOST'")
		case ENV.EMAIL_USERNAME == "":
			return fmt.Errorf("error: couldn't load environment variable 'EMAIL_USERNAME'")
		case ENV.EMAIL_PASSWORD == "":
			return fmt.Errorf("error: couldn't load environment variable 'EMAIL_PASSWORD'")
		case ENV.EMAIL == "":
			return fmt.Errorf("error: couldn't load environment variable 'EMAIL'")
		case ENV.GEMINI_API_KEY == "":
			return fmt.Errorf("error: couldn't load environment variable 'GEMINI_API_KEY'")
		case ENV.REDIS_ADDR == "":
			return fmt.Errorf("error: couldn't load environment variable 'REDIS_ADDR'")
		case ENV.REDIS_PASSWORD == "":
			return fmt.Errorf("error: couldn't load environment variable 'REDIS_PASSWORD'")
		case ENV.GOOGLE_CLIENT_ID == "":
			return fmt.Errorf("error: couldn't load environment variable 'GOOGLE_CLIENT_ID'")
		case ENV.GOOGLE_CLIENT_SECRET == "":
			return fmt.Errorf("error: couldn't load environment variable 'GOOGLE_CLIENT_SECRET'")
		default:
			return nil
	}
}