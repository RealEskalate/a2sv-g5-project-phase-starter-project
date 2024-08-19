package bootstrap

import (
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	Mail 				   string `mapstructure:"MAIL"`
	MailPassword 		   string `mapstructure:"MAIL_PASSWORD"`
	SmtpServer 			   string `mapstructure:"SMTP_SERVER"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}


func NewEnv() *Env {
	env := Env{}
	if err := godotenv.Load(); err != nil {
    log.Println("No .env file found")
	}

	err := setEnv(&env)

	if err != nil {
		log.Fatalf("Error populating Env struct: %v", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}

func setEnv(envStruct *Env) error{
	val := reflect.ValueOf(envStruct).Elem()
	typ := val.Type()
	for i := range typ.NumField() {
		field := typ.Field(i)
		envKey := field.Tag.Get("mapstructure")

		if envKey != "" {
		envValue := os.Getenv(envKey)
		fieldType := field.Type

		switch fieldType.Kind() {
		case reflect.String:
			val.Field(i).SetString(envValue)
		case reflect.Int:
			intValue, err := strconv.Atoi(envValue)
			if err != nil {
			return err
			}
			val.Field(i).SetInt(int64(intValue))
		}
    	}
	}
	return nil
}