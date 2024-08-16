package route

import (
	"blog/config"
	"blog/database"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *config.Env, timeout time.Duration, db database.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	// NewLoginRouter(env, timeout, db, publicRouter)
	// NewRefreshTokenRouter(env, timeout, db, publicRouter)

	// protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	// protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	// NewProfileRouter(env, timeout, db, protectedRouter)

}
