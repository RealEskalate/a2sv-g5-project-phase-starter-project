package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                  string `mapstructure:"APP_ENV"`
	ServerAddress           string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout          int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                  string `mapstructure:"DB_HOST"`
	DBPort                  string `mapstructure:"DB_PORT"`
	DBUser                  string `mapstructure:"DB_USER"`
	DBPass                  string `mapstructure:"DB_PASS"`
	DBName                  string `mapstructure:"DB_NAME"`
	UserCollection          string `mapstructure:"USER_COLLECTION"`
	BlogCollection          string `mapstructure:"BLOG_COLLECTION"`
	OtpCollection           string `mapstructure:"OTP_COLLECTION"`
	LikeCollection          string `mapstructure:"LIKE_COLLECTION"`
	EmailApiKey             string `mapstructure:"EMAIL_API_KEY"`
	GeminiApiKey            string `mapstructure:"GEMINI_API_KEY"`
	AccessTokenExpiryMinute int    `mapstructure:"ACCESS_TOKEN_EXPIRY_MINUTE"`
	AccessTokenSecret       string `mapstructure:"ACCESS_TOKEN"`
	RefreshTokenSecret      string `mapstructure:"REFRESH_TOKEN"`
	RefreshTokenExpiryHour  int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	RedisHost               string `mapstructure:"REDIS_HOST"`
	RedisPort               string `mapstructure:"REDIS_PORT"`
	RedisPassword           string `mapstructure:"REDIS_PASSWORD"`
}

func NewEnv() *Env {
	env := &Env{}

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env")
	}

	err = viper.Unmarshal(env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return env
}
