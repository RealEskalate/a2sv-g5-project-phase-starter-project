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

	privateRouter := publicRouter.Group("")
	privateRouter.Use(auth.AuthMiddleware(""))
	
	adminRouter := publicRouter.Group("")
	adminRouter.Use(auth.AuthMiddleware("admin"))


	NewBlogRouter(db, publicRouter.Group("/blogs"), model)
	NewCommmentRouter(db, publicRouter.Group("/comments"))	

	gin.Run(":8080")
}