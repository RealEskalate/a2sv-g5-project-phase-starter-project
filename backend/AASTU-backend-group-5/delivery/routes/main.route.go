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
		Collection: clinect.Client.Database("blogPost").Collection("Users"),
	}
	aiRoute := router.Group("")
	userRoute := router.Group("")

	NewAiRequestRoute(aiRoute)
	NewUserRoute(userRoute, userCollection)
}
