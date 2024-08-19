package bootstrap

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func InitEnv() {
	viper.SetConfigFile(".env") // Specify the file to read

	// Automatically look for environment variables that match
	viper.AutomaticEnv()

	// Reading in the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}

func GetEnv(key string) (string, error) {
	if !viper.IsSet(key) {
		return "", fmt.Errorf("environment variable %s not found", key)
	}

	return viper.GetString(key), nil
}
