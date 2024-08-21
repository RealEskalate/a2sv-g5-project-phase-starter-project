package routes

import (
	bootstrap "aait-backend-group4/Bootstrap"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	gin.LoadHTMLGlob("templates/*")

	publicRouter := gin.Group("")

	NewAiRoute(env, timeout, db, publicRouter)
	NewSignupRoute(env, timeout, db, publicRouter)
	NewLoginRoute(env, timeout, db, publicRouter)
	NewPromoteRoute(env, timeout, db, publicRouter)
	NewOtpRoute(env, timeout, db, publicRouter)
	NewBlogRoute(env, timeout, db, publicRouter)
}
