package main

import (
	"AAiT-backend-group-6/bootstrap"
	"AAiT-backend-group-6/delivery/routers"
	// "time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	println(db)
	defer app.CloseDBConnection()

	// timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	routers.SetupRouter(db, gin)

	gin.Run(env.ServerAddress)
}
