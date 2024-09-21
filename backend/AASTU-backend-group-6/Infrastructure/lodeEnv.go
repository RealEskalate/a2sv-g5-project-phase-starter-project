package infrastructure

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl              string
	RedisURL                 string
	Port                     int
	DbName                   string
	UserCollection           string
	BlogCollection           string
	CommentCollection        string
	ActiveUserCollection     string
	UnverifiedUserCollection string
	ContextTimeout           int
	AccessTokenExpiryHour    int
	RefreshTokenExpiryHour   int
	AccessTokenSecret        string
	RefreshTokenSecret       string
	ClientID                 string
	ClientSecret             string
	RedirectURL              string
	OauthSecret              string
	GeminiAPIKey             string
}

func LoadEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		// return nil, err
	}

	dbURL := Getenv("DATABASE_URL", "")
	portStr := Getenv("PORT", "")
	dbname := Getenv("DB_NAME", "")
	usercoll := Getenv("user_collection", "")
	blogcoll := Getenv("blog_collection", "")
	commentcoll := Getenv("comment_collection", "")
	activeusercoll := Getenv("ACTIVE_USER_COLLECTION", "")
	unverifiedusercoll := Getenv("UNVERIFIED_USER_COLLECTION", "")
	contextTimeoutStr := Getenv("CONTEXT_TIMEOUT", "")
	accessTokenExpiryHourStr := Getenv("ACCESS_TOKEN_EXPIRY_HOUR", "")
	refreshTokenExpiryHourStr := Getenv("REFRESH_TOKEN_EXPIRY_HOUR", "")
	accessTokenSecret := Getenv("ACCESS_TOKEN_SECRET", "")
	refreshTokenSecret := Getenv("REFRESH_TOKEN_SECRET", "")
	clientId := Getenv("CLIENT_ID", "")
	clientSecret := Getenv("CLIENT_SECRET", "")
	redirectURL := Getenv("REDIRECT_URI", "")
	oauthSecret := Getenv("OAUTH_STATE_STRING", "")
	geminiAPIKey := Getenv("GEMINI_API_KEY", "")

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
		DatabaseUrl:              dbURL,
		Port:                     port,
		DbName:                   dbname,
		UserCollection:           usercoll,
		BlogCollection:           blogcoll,
		CommentCollection:        commentcoll,
		ActiveUserCollection:     activeusercoll,
		UnverifiedUserCollection: unverifiedusercoll,
		ContextTimeout:           contextTimeout,
		AccessTokenExpiryHour:    accessTokenExpiryHour,
		RefreshTokenExpiryHour:   refreshTokenExpiryHour,
		AccessTokenSecret:        accessTokenSecret,
		RefreshTokenSecret:       refreshTokenSecret,
		ClientID:                 clientId,
		ClientSecret:             clientSecret,
		RedirectURL:              redirectURL,
		OauthSecret:              oauthSecret,
		GeminiAPIKey:             geminiAPIKey,
	}

	return config, nil
}
func Getenv(key,defaul string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Println("No value found for ", key)
		return defaul
	}
	return value
}
