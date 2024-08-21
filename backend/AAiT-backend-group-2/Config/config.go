package config

import (
	"fmt"
	domain "AAiT-backend-group-2/Domain"

	"github.com/spf13/viper"
)


func LoadConfig(path string) (config domain.Config, err error) {
	viper.AddConfigPath(path)
    viper.SetConfigType("env")
    viper.SetConfigName(".env")

    // Enable Viper to read environment variables directly
    viper.AutomaticEnv()

    // Read the config file
    err = viper.ReadInConfig()
    if err != nil {
        fmt.Println("Error reading config file:", err)
        return
    }

    // Unmarshal the config into the struct
    err = viper.Unmarshal(&config)
    if err != nil {
        fmt.Println("Error unmarshalling config:", err)
        return
    }

    return
}

	 