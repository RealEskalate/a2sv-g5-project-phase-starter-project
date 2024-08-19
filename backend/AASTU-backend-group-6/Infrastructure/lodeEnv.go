package infrastructure

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl            string
	Port                   int
	DbName                 string
	UserCollection         string
	BlogCollection         string
	ActiveUserCollection   string
	ContextTimeout         int
	AccessTokenExpiryHour  int
	RefreshTokenExpiryHour int
	AccessTokenSecret      string
	RefreshTokenSecret     string
}

func LoadEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	dbURL := os.Getenv("DATABASE_URL")
	portStr := os.Getenv("PORT")
	dbname := os.Getenv("DB_NAME")
	usercoll := os.Getenv("user_collection")
	blogcoll := os.Getenv("blog_collection")
	activeusercoll := os.Getenv("ACTIVE_USER_COLLECTION")
	contextTimeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	accessTokenExpiryHourStr := os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR")
	refreshTokenExpiryHourStr := os.Getenv("REFRESH_TOKEN_EXPIRY_HOUR")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	refreshTokenSecret := os.Getenv("REFRESH_TOKEN_SECRET")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("Invalid PORT value")
		return nil, err
	}

	contextTimeout, err := strconv.Atoi(contextTimeoutStr)
	if err != nil {
		log.Fatal("Invalid CONTEXT_TIMEOUT value")
		return nil, err
	}

	accessTokenExpiryHour, err := strconv.Atoi(accessTokenExpiryHourStr)
	if err != nil {
		log.Fatal("Invalid ACCESS_TOKEN_EXPIRY_HOUR value")
		return nil, err
	}

	refreshTokenExpiryHour, err := strconv.Atoi(refreshTokenExpiryHourStr)
	if err != nil {
		log.Fatal("Invalid REFRESH_TOKEN_EXPIRY_HOUR value")
		return nil, err
	}

	config := &Config{
		DatabaseUrl:            dbURL,
		Port:                   port,
		DbName:                 dbname,
		UserCollection:         usercoll,
		BlogCollection:         blogcoll,
		ActiveUserCollection:   activeusercoll,
		ContextTimeout:         contextTimeout,
		AccessTokenExpiryHour:  accessTokenExpiryHour,
		RefreshTokenExpiryHour: refreshTokenExpiryHour,
		AccessTokenSecret:      accessTokenSecret,
		RefreshTokenSecret:     refreshTokenSecret,
	}

	return config, nil
}
