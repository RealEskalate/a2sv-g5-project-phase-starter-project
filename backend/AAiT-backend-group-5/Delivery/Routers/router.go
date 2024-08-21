package routers

import (
	"context"

	config "github.com/aait.backend.g5.main/backend/Config"
	middlewares "github.com/aait.backend.g5.main/backend/Delivery/middlewares"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *config.Env, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	protectedRouter := gin.Group("")
	adminRouter := gin.Group("")
	redisClient := config.NewRedisClient(*env, context.Background())

	jwt_service := infrastructure.NewJwtService(env)
	protectedRouter.Use(middlewares.JWTAuthMiddelware(jwt_service))

	adminRouter.Use(
		middlewares.JWTAuthMiddelware(jwt_service),
		middlewares.AuthenticateAdmin(),
	)

	NewAuthenticationRouter(env, db, publicRouter)
	NewForgotPasswordRouter(env, db, publicRouter)
	NewLogoutRouter(env, db, protectedRouter)
	NewPromoteDemoteRouter(db, adminRouter)
	NewBlogRouter(env, db, protectedRouter, redisClient)
}
