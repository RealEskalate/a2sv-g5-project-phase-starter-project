package main

import (
	"backend-starter-project/bootstrap"
	"backend-starter-project/delivery/router"
	"backend-starter-project/infrastructure/middleware"
	"backend-starter-project/repository"
	"backend-starter-project/service"

	"github.com/gin-gonic/gin"
)

func main()  {
	app := bootstrap.App()

	env := app.Env

	mongoClient := app.Mongo

	db := (*mongoClient).Database(env.DBName)
	
	defer app.CloseDBConnection()
	defer app.CloseModelClient()
	defer app.CloseRedisConnection()

	tr := repository.NewTokenRepository(db, db.Collection("tokens"))
	ur := repository.NewUserRepository(db.Collection("users"))
	ts := service.NewTokenService(env.AccessTokenSecret,env.RefreshTokenSecret, tr,ur)
	authMiddleware := middleware.NewAuthMiddleware(ts)

	// Do something with the database
	route.Setup(env, db, gin.Default(), authMiddleware, app.GenAi, app.Redis)
}