package routes

import (
	"github.com/RealEskalate/blogpost/config"
	"github.com/RealEskalate/blogpost/database"
	"github.com/gin-gonic/gin"
)

func SetUp(router *gin.Engine) {
	var clinect config.ServerConnection
	clinect.Connect_could()

	userCollection := &database.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("Users"),
	}
	blogCollection := &database.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("Blogs"),
	}
	aiRoute := router.Group("")
	userRoute := router.Group("")
	blogRot := router.Group("")

	NewBlogRoutes(blogRot, blogCollection, userCollection)
	NewAiRequestRoute(aiRoute)
	NewUserRoute(userRoute, userCollection)

}
