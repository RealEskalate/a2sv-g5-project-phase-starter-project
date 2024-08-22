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
	blogCollection := &database.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("Blogs"),
	}

  
	stateCollection := &database.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("States"),
	}
	userrepo := repository.NewUserRepository(userCollection)

	aiRoute := router.Group("")
	userRoute := router.Group("")
	verifiRoute := router.Group("")
	uplaodRoute := router.Group("")
	authRoute := router.Group("")
  	blogRot := router.Group("")
  popularityRoute := router.Group("")
	likeRoute := router.Group("")
	dislikeRoute := router.Group("")
	commentRoute := router.Group("")

	NewUploadRoute(uplaodRoute , *userrepo)
	NewVerifyEmialRoute(verifiRoute , userCollection)	
	NewBlogRoutes(blogRot, blogCollection, userCollection)
	NewAiRequestRoute(aiRoute)
	NewUserRoute(userRoute, userCollection)
  	NewAuthRoute(authRoute, userCollection, stateCollection)
  	NewPopularityRoutes(popularityRoute, blogCollection)
	NewLikeRoutes(likeRoute, blogCollection)``
	NewDislikeRoutes(dislikeRoute, blogCollection)
	NewCommentRoutes(commentRoute, blogCollection)
}
