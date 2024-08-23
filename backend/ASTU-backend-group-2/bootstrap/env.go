package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                     string `mapstructure:"APP_ENV"`
	ServerAddress              string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout             int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                     string `mapstructure:"DB_HOST"`
	DBPort                     string `mapstructure:"DB_PORT"`
	DBUser                     string `mapstructure:"DB_USER"`
	DBPass                     string `mapstructure:"DB_PASS"`
	DBName                     string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour      int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour     int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret          string `mapstructure:"ACCESS_TOKEN_SECRET"`
	VerificationTokenExpiryMin int    `mapstructure:"VERIFICATION_TOKEN_EXPIRY_MIN"`
	VerificationTokenSecret    string `mapstructure:"VERIFICATION_TOKEN_SECRET"`
	RefreshTokenSecret         string `mapstructure:"REFRESH_TOKEN_SECRET"`
	SenderEmail                string `mapstructure:"SENDER_EMAIL"`
	SmtpPort                   string `mapstructure:"SMTP_PORT"`
	SmtpHost                   string `mapstructure:"SMTP_HOST"`
	SenderPassword             string `mapstructure:"SENDER_PASSWORD"`
	GoogleClientID             string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret         string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	GoogleUrlCallback          string `mapstructure:"GOOGLE_URL_CALLBACK"`
	SessionSecret              string `mapstructure:"SESSION_SECRET"`
	GeminiAPIKey               string `mapstructure:"GEMINI_API_KEY"`
	GeminiWordCount            string `mapstructure:"GEMINI_WORD_COUNT"`
	PassResetCodeExpirationMin int    `mapstructure:"PASS_RESET_CODE_EXPIRATION_MIN"`
	CloudinaryName             string `mapstructure:"CLOUDINARY_NAME"`
	CloudinaryKey              string `mapstructure:"CLOUDINARY_KEY"`
	CloudinarySecret           string `mapstructure:"CLOUDINARY_SECRET"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	log.Println(env)
	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
