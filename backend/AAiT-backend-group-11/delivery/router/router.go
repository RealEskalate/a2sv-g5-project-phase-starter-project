package route

import (
	"backend-starter-project/bootstrap"
	"backend-starter-project/infrastructure/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/generative-ai-go/genai"
	"go.mongodb.org/mongo-driver/mongo"
)


func Setup(env *bootstrap.Env, db *mongo.Database, gin *gin.Engine, auth middleware.AuthMiddleware, model *genai.GenerativeModel, redis *redis.Client) {
	
	publicRouter := gin.Group("")

	privateRouter := gin.Group("")
	privateRouter.Use(auth.AuthMiddleware(""))
	
	adminRouter := gin.Group("")
	adminRouter.Use(auth.AuthMiddleware("admin"))

	NewBlogRouter(db, privateRouter.Group("/blogs"), model, redis)
	NewCommmentRouter(db, privateRouter.Group("/comments"))	
	NewAuthRouter(env,db, publicRouter.Group("/auth"))
	NewProfileRouter(env,db, privateRouter.Group("/user"))

	gin.Run(":8080")
}