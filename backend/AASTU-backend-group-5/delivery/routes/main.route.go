package routes

import (
	"github.com/RealEskalate/blogpost/config"
	"github.com/RealEskalate/blogpost/database"
	"github.com/gin-gonic/gin"
)

func SetUp(router *gin.Engine) {
	var clinect config.ServerConnection
	clinect.Connect_could()

	// Initialize collections
	userCollection := &database.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("Users"),
	}
	blogCollection := &database.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("Blogs"),
	}

	stateCollection := &database.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("States"),
	}
	aiRoute := router.Group("")
	userRoute := router.Group("")
	verifiRoute := router.Group("")
	uplaodRoute := router.Group("")
	authRoute := router.Group("")
	blogRot := router.Group("")

	NewBlogRoutes(blogRoute, blogCollection, userCollection)
	NewAiRequestRoute(aiRoute)
	NewUserRoute(userRoute, userCollection)
	NewAuthRoute(authRoute, userCollection, stateCollection)
}
