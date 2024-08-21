package routers

import (
	"context"

	config "github.com/aait.backend.g5.main/backend/Config"
	middelwares "github.com/aait.backend.g5.main/backend/Delivery/middlewares"
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
	protectedRouter.Use(middelwares.JWTAuthMiddelware(jwt_service))

	adminRouter.Use(
		middelwares.JWTAuthMiddelware(jwt_service),
		middelwares.AuthenticateAdmin(),
	)

	NewAuthenticationRouter(env, db, publicRouter)
	NewForgotPasswordRouter(env, db, protectedRouter)
	NewLogoutRouter(env, db, protectedRouter)
	NewPromoteDemoteRouter(db, adminRouter)
	NewRefreshRouter(env, db, protectedRouter)

	NewAISuggestionRouter(env, publicRouter)

	// NewBlogRouter(env, db, protectedRouter)
	NewBlogRouter(env, db, publicRouter, redisClient)
}
