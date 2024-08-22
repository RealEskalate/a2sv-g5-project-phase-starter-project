package routes

import (
	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/repository"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
)

func NewLikeRoutes(group *gin.RouterGroup, collection database.CollectionInterface) {
	likeRepo := repository.NewLikeRepository(collection)
	dislikeRepo := repository.NewDislikeRepository(collection)
	likeUseCase := usecase.NewLikeUseCase(likeRepo, dislikeRepo)
	likeController := controller.NewLikeController(likeUseCase)

	group.GET("/like/:post_id", likeController.GetLikes)
	group.POST("/like/:user_id/:post_id", likeController.CreateLike)
	group.PUT("/like/toggle/:user_id/:post_id", likeController.ToggleLike)
	group.DELETE("/like/:user_id/:post_id", likeController.RemoveLike)
}

func NewDislikeRoutes(group *gin.RouterGroup, collection database.CollectionInterface) {
	dislikeRepo := repository.NewDislikeRepository(collection)
	likeRepo := repository.NewLikeRepository(collection)
	dislikeUseCase := usecase.NewDislikeUseCase(dislikeRepo, likeRepo)
	dislikeController := controller.NewDislikeController(dislikeUseCase)

	group.GET("/dislike/:post_id", dislikeController.GetDislikes)
	group.POST("/dislike/:user_id/:post_id", dislikeController.CreateDislike)
	group.PUT("/dislike/toggle/:user_id/:post_id", dislikeController.ToggleDislike)
	group.DELETE("/dislike/:user_id/:post_id", dislikeController.RemoveDislike)
}

func NewCommentRoutes(group *gin.RouterGroup, collection database.CollectionInterface) {
	commentRepo := repository.NewCommentRepository(collection)
	commentUseCase := usecase.NewCommentUseCase(commentRepo)
	commentController := controller.NewCommentController(commentUseCase)

	group.GET("/comment/:post_id", commentController.GetComments())
	group.POST("/comment/:post_id", commentController.CreateComment())
	group.PUT("/comment/:comment_id", commentController.UpdateComment())
	group.DELETE("/comment/:comment_id", commentController.DeleteComment())
}
