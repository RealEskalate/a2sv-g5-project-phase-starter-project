package router

import (
	"Blog_Starter/config"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *config.Environment, timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewSignupRouter(timeout, db, publicRouter)

	privateRouter := gin.Group("/api/blog")
	NewBlogRouter(timeout, db, privateRouter)
	NewBlogRatingRouter(timeout, db, privateRouter)
	NewBlogCommentRouter(timeout, db, privateRouter)
}
