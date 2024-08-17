package bootstrap

import (
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

// Env is a struct to hold environment variables.
type Env struct {
	AppEnv                      string `mapstructure:"APP_ENV"`
	ContextTimeout              int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                      string `mapstructure:"DB_HOST"`
	DBPort                      string `mapstructure:"DB_PORT"`
	DBName                      string `mapstructure:"DB_NAME"`
	DBUri                       string `mapstructure:"MONGODB_URI"`
	AccessTokenExpiryHour       int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour      int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret           string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret          string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

// NewEnv initializes and returns a new instance of the Env struct.
func NewEnv() *Env {
	env := Env{}
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	err := envMapToStruct(&env)
	if err != nil {
		log.Fatalf("Error populating Env struct: %v", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}

// envMapToStruct populates a struct with values from environment variables using reflection.
func envMapToStruct(envStruct interface{}) error {
	// Iterate through the fields of the struct and set their values from environment variables
	structValue := reflect.ValueOf(envStruct).Elem()
	structType := structValue.Type()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		envKey := field.Tag.Get("mapstructure")

		if envKey != "" {
			envValue := os.Getenv(envKey)
			fieldType := field.Type

			switch fieldType.Kind() {
			case reflect.String:
				structValue.Field(i).SetString(envValue)
			case reflect.Int:
				intValue, err := strconv.Atoi(envValue)
				if err != nil {
					return err
				}
				structValue.Field(i).SetInt(int64(intValue))
				// Add cases for other supported data types as needed
			}
		}
	}

	return nil
}
