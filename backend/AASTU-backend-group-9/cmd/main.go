package main

import (
	"time"

	"blog/config"
	route "blog/delivery/route"

	"github.com/gin-gonic/gin"
)

func main() {

	app := config.App()
	config.CreateRootUser(&app.Mongo, app.Env)
	env := app.Env
	db := app.Mongo.Database(env.DBName)
	ai := config.NewAIConfig(env)
	defer app.CloseDBConnection()
	defer config.CloseAIConnection(ai)

	timeout := time.Duration(env.ContextTimeout) * time.Second
	gin := gin.Default()

	route.Setup(env, timeout, db, gin, ai)

	gin.Run()
}
