package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBURI                  string `mapstructure:"DB_URI"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	ResetTokenExpiryHour   int    `mapstructure:"RESET_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	ResetTokenSecret       string `mapstructure:"RESET_TOKEN_SECRET"`
	SMTPServer             string `mapstructure:"SMTP_SERVER"`
	SMTPPort               string `mapstructure:"SMTP_PORT"`
	SMTPUser               string `mapstructure:"SMTP_USER"`
	SMTPPassword           string `mapstructure:"SMTP_PASSWORD"`
	FromAddress            string `mapstructure:"FROM_ADDRESS"`
	FrontendBaseURL        string `mapstructure:"FRONTEND_BASE_URL"`
	API_KEY                string `mapstructure:"API_KEY"`
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

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
