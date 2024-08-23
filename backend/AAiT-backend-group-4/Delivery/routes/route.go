package routes

import (
	bootstrap "aait-backend-group4/Bootstrap"
	infrastructure "aait-backend-group4/Infrastructure"
	"aait-backend-group4/Infrastructure/middlewares"
	repositories "aait-backend-group4/Repositories"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine, rc redis.Client) {
	gin.LoadHTMLGlob("templates/*")

	publicRouter := gin.Group("")

	// Public routes
	NewOtpRoute(env, timeout, db, publicRouter)
	NewLoginRoute(env, timeout, db, publicRouter)

	// Public routes with image interceptor
	publicRouterWithImage := gin.Group("")
	publicRouterWithImage.Use(middlewares.ImageUploadMiddleware())
	NewSignupRoute(env, timeout, db, publicRouterWithImage)

	userRepository := repositories.NewUserRepository(db, env.UserCollection)
	protectedRouter := gin.Group("")
	tokenService := infrastructure.NewTokenService(
		userRepository,
		env,
	)
	protectedRouter.Use(middlewares.AuthMiddleware(env, tokenService))

	// Protected routes
	NewAiRoute(env, timeout, db, protectedRouter)
	NewBlogRoute(env, timeout, db, protectedRouter, rc)
	NewLikeRoute(env, timeout, db, protectedRouter)
	NewForgotPasswordRoute(env, timeout, db, protectedRouter)

	protectedAdminRouter := gin.Group("")
	protectedAdminRouter.Use(middlewares.AdminMiddleware())
	protectedAdminRouter.Use(middlewares.AuthMiddleware(env, tokenService))
	// Admin routes
	NewPromoteRoute(env, timeout, db, publicRouter)
}
