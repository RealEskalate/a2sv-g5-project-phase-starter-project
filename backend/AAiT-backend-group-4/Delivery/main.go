package main

import (
	bootstrap "aait-backend-group4/Bootstrap"
	"aait-backend-group4/Delivery/routes"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.Close()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	routes.Setup(env, timeout, *db, gin)

	gin.Run(env.ServerAddress)

}
