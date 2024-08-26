package route

import (
	"blog/config"
	"blog/database"
	"blog/delivery/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
)

func Setup(env *config.Env, timeout time.Duration, db database.Database, gin *gin.Engine, ai *genai.Client) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewForgotPasswordRouter(env, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.AuthMidd)
	// All Private APIs
	NewUserRouter(env, timeout, db, protectedRouter)
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewLogoutRouter(env, timeout, db, protectedRouter)
	RegisterAIRoutes(env, timeout, protectedRouter, ai)
	RegisterBlogRoutes(env, timeout, db, protectedRouter)
}
