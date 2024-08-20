package route

import (
	"backend-starter-project/bootstrap"
	"backend-starter-project/infrastructure/middleware"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"go.mongodb.org/mongo-driver/mongo"
)


func Setup(env *bootstrap.Env, db *mongo.Database, gin *gin.Engine, auth middleware.AuthMiddleware, model *genai.GenerativeModel) {
	
	publicRouter := gin.Group("")

	privateRouter := gin.Group("")
	privateRouter.Use(auth.AuthMiddleware(""))
	
	adminRouter := gin.Group("")
	adminRouter.Use(auth.AuthMiddleware("admin"))
	_ = publicRouter


	NewBlogRouter(db, publicRouter.Group("/blogs"), model)
	NewCommmentRouter(db, publicRouter.Group("/comments"))	

	gin.Run(":8080")
}