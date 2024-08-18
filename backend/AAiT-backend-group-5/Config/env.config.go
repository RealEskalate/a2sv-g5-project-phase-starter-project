package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
}

func NewEnv() *Env {
	viper.AutomaticEnv()

	env := Env{}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal("Error getting the enviroment variables")
	}

	return &env
}
