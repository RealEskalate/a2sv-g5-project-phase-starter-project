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
	likeCollection := &database.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("Likes"),
	}
	dislikeCollection := &database.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("DisLikes"),
	}
	commentCollection := &database.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("Comments"),
	}

	stateCollection := &database.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("States"),
	}
	bookmarkCollection := &database.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("Bookmarks"),
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
	bookmarkRoute := router.Group("")

	NewUploadRoute(uplaodRoute, *userrepo)
	NewVerifyEmialRoute(verifiRoute, userCollection)
	NewBlogRoutes(blogRot, blogCollection, userCollection)
	NewAiRequestRoute(aiRoute)
	NewUserRoute(userRoute, userCollection)
	NewAuthRoute(authRoute, userCollection, stateCollection)
	NewPopularityRoutes(popularityRoute, blogCollection)
	NewLikeRoutes(likeRoute, likeCollection, blogCollection)
	NewDislikeRoutes(dislikeRoute, dislikeCollection, blogCollection)
	NewCommentRoutes(commentRoute, commentCollection, blogCollection)
	NewBookmarkRoutes(bookmarkRoute, bookmarkCollection, userCollection)

}
