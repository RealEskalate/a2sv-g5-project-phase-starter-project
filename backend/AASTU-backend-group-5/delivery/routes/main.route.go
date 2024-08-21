package routes

import (
	"github.com/RealEskalate/blogpost/config"
	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/repository"
	"github.com/gin-gonic/gin"
)

func SetUp(router *gin.Engine) {
	var clinect config.ServerConnection
	clinect.Connect_could()

	userCollection := &database.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("Users"),
	}
	aiRoute := router.Group("")
	userRoute := router.Group("")
	verifiRoute := router.Group("")
	uplaodRoute := router.Group("")

	userrepo := repository.NewUserRepository(userCollection)
	NewUploadRoute(uplaodRoute , *userrepo)
	NewVerifyEmialRoute(verifiRoute , userCollection)	
	NewAiRequestRoute(aiRoute)
	NewUserRoute(userRoute, userCollection)
}
