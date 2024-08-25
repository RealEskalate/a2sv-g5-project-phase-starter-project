package routers

import (
	"group3-blogApi/config/db"
	"group3-blogApi/delivery/controllers"
	"group3-blogApi/infrastracture"
	"group3-blogApi/repository"
	"group3-blogApi/usecase"

	"github.com/gin-gonic/gin"
)

func SetUpComment(router *gin.Engine) {

	commentRepo := repository.NewCommentRepository(db.CommentCollection)
	userRepo := repository.NewUserRepositoryImpl(db.UserCollection)

	commentUseCase := usecase.NewCommentUsecase(commentRepo, userRepo)
	commentController := controllers.NewCommentController(commentUseCase)

	comments := router.Group("/")
	comments.Use(infrastracture.AuthMiddleware())
	{
		comments.POST("/comments", commentController.CreateComment)
		comments.PUT("/comments/:id", commentController.UpdateComment)
		comments.DELETE("/comments/:id", commentController.DeleteComment)
		comments.GET("/comments/:id", commentController.GetCommentByID)
		comments.GET("/posts/:postID/comments", commentController.GetComments)

		comments.POST("/replies", commentController.CreateReply)
		comments.PUT("/replies/:id", commentController.UpdateReply)
		comments.DELETE("/replies/:id", commentController.DeleteReply)
		comments.GET("/replies/comments/:commentID", commentController.GetReplies)

		comments.POST("/comments/:id/like", commentController.LikeComment)
		comments.POST("/comments/:id/unlike" ,commentController.UnlikeComment)
		comments.POST("/replies/:id/like", commentController.LikeReply)
		comments.POST("/replies/:id/unlike", commentController.UnlikeReply)

	}

}
