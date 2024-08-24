package route

import (
	"backend-starter-project/bootstrap"
	"backend-starter-project/infrastructure/middleware"
	"backend-starter-project/mongo"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
)


func Setup(env *bootstrap.Env, db mongo.Database, gin *gin.Engine, auth middleware.AuthMiddleware, model *genai.GenerativeModel) {
	
	publicRouter := gin.Group("")

	privateRouter := gin.Group("")
	privateRouter.Use(auth.AuthMiddleware(""))
	
	adminRouter := gin.Group("")
	adminRouter.Use(auth.AuthMiddleware("admin"))
	_ = publicRouter


	NewBlogRouter(&db, privateRouter.Group("/blogs"), model)
	NewCommmentRouter(&db, privateRouter.Group("/comments"))	
	NewAuthRouter(env,db, publicRouter.Group("/auth"))
	NewProfileRouter(&db, privateRouter.Group("/user"))

	gin.Run(":8080")
}