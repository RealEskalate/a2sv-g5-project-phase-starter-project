package router

import (
	"Blog_Starter/config"
	"time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo"
)

// Setup initializes the routes for the application.
func Setup(env *config.Env, timeout time.Duration, db *mongo.Client, gin *gin.Engine) {
    publicRouter := gin.Group("")

    NewSignupRouter(env, timeout, db, publicRouter)
	privateRouter := gin.Group("/api/blog")
	NewBlogRouter(env, timeout, db, privateRouter)
	NewBlogRatingRouter(env, timeout, db, privateRouter)
	NewBlogCommentRouter(env, timeout, db, privateRouter)
}