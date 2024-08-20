package main

import (
	"backend-starter-project/bootstrap"
	"backend-starter-project/delivery/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main()  {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	route := gin.Default()

	// Global CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	route.Use(cors.New(config))

	// Setup routes
	router.Setup(env, db, route)

	route.Run()
}