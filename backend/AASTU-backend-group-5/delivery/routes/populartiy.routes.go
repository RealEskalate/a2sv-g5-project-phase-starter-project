package routes

import (
	"log"
	"os"

	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/infrastructure/middleware"
	tokenservice "github.com/RealEskalate/blogpost/infrastructure/token_service"
	"github.com/RealEskalate/blogpost/repository"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func NewLikeRoutes(group *gin.RouterGroup, collection database.CollectionInterface, blog database.CollectionInterface) {

	//load middlewares
	err := godotenv.Load()
	if err != nil {
        log.Panic(err.Error())
    }
	access_secret := os.Getenv("ACCESSTOKENSECRET")
	if access_secret == ""{
		log.Panic("No accesstoken")
	}
	
	refresh_secret := os.Getenv("REFRESHTOKENSECRET")
	if refresh_secret == ""{
		log.Panic("No refreshtoken")
	}
	TokenSvc := *tokenservice.NewTokenService(access_secret, refresh_secret)

	LoggedInmiddleWare := middleware.LoggedIn(TokenSvc)

	likeRepo := repository.NewLikeRepository(collection, blog)
	dislikeRepo := repository.NewDislikeRepository(collection, blog)
	likeUseCase := usecase.NewLikeUseCase(likeRepo, dislikeRepo)
	likeController := controller.NewLikeController(likeUseCase)

	group.GET("/like/:post_id", LoggedInmiddleWare,likeController.GetLikes)
	group.POST("/like/:post_id", LoggedInmiddleWare,likeController.CreateLike)
	group.PUT("/like/toggle/:user_id/:post_id", LoggedInmiddleWare,likeController.ToggleLike)
	group.DELETE("/like/:user_id/:post_id", LoggedInmiddleWare,likeController.RemoveLike)
}

func NewDislikeRoutes(group *gin.RouterGroup, collection database.CollectionInterface, blog database.CollectionInterface) {
	dislikeRepo := repository.NewDislikeRepository(collection, blog)
	likeRepo := repository.NewLikeRepository(collection, blog)
	dislikeUseCase := usecase.NewDislikeUseCase(dislikeRepo, likeRepo)
	dislikeController := controller.NewDislikeController(dislikeUseCase)

	//load middlewares
	err := godotenv.Load()
	if err != nil {
        log.Panic(err.Error())
    }
	access_secret := os.Getenv("ACCESSTOKENSECRET")
	if access_secret == ""{
		log.Panic("No accesstoken")
	}
	
	refresh_secret := os.Getenv("REFRESHTOKENSECRET")
	if refresh_secret == ""{
		log.Panic("No refreshtoken")
	}
	TokenSvc := *tokenservice.NewTokenService(access_secret, refresh_secret)

	LoggedInmiddleWare := middleware.LoggedIn(TokenSvc)

	group.GET("/dislike/:post_id",LoggedInmiddleWare, dislikeController.GetDislikes)
	group.POST("/dislike/:post_id", LoggedInmiddleWare,dislikeController.CreateDislike)
	group.PUT("/dislike/toggle/:post_id",LoggedInmiddleWare, dislikeController.ToggleDislike)
	group.DELETE("/dislike/:post_id", LoggedInmiddleWare,dislikeController.RemoveDislike)
}

func NewCommentRoutes(group *gin.RouterGroup, collection database.CollectionInterface, blog database.CollectionInterface) {
	commentRepo := repository.NewCommentRepository(collection, blog)
	commentUseCase := usecase.NewCommentUseCase(commentRepo)
	commentController := controller.NewCommentController(commentUseCase)

	//load middlewares
	err := godotenv.Load()
	if err != nil {
        log.Panic(err.Error())
    }
	access_secret := os.Getenv("ACCESSTOKENSECRET")
	if access_secret == ""{
		log.Panic("No accesstoken")
	}
	
	refresh_secret := os.Getenv("REFRESHTOKENSECRET")
	if refresh_secret == ""{
		log.Panic("No refreshtoken")
	}
	TokenSvc := *tokenservice.NewTokenService(access_secret, refresh_secret)

	LoggedInmiddleWare := middleware.LoggedIn(TokenSvc)

	group.GET("/comment/:post_id",LoggedInmiddleWare, commentController.GetComments())
	group.POST("/comment/:post_id", LoggedInmiddleWare,commentController.CreateComment())
	group.PUT("/comment/:comment_id", LoggedInmiddleWare,commentController.UpdateComment())
	group.DELETE("/comment/:comment_id", LoggedInmiddleWare,commentController.DeleteComment())
}
