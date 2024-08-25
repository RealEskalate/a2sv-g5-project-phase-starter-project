package main

import (
	"AAiT-backend-group-6/bootstrap"
	"AAiT-backend-group-6/delivery/route"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	db := app.Mongo.Database(env.DBName)

	// if app.Redis == nil {
	// 	panic("Redis is not connected")
	// }
	redisClient := app.Redis
	defer app.Close()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin, redisClient)
	gin.Run(env.ServerAddress)
}
