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
	commentUseCase := usecase.NewCommentUsecase(commentRepo)
	blogController := controllers.NewCommentController(commentUseCase)

	comments := router.Group("/")
	comments.Use(infrastracture.AuthMiddleware())
	{
		comments.POST("/comments", blogController.CreateComment)
		comments.PUT("/comments/:id", blogController.UpdateComment)
		comments.DELETE("/comments/:id", blogController.DeleteComment)
		comments.GET("/comments/:id", blogController.GetCommentByID)
		comments.GET("/posts/:postID/comments", blogController.GetComments)

		comments.POST("/replies", blogController.CreateReply)
		comments.PUT("/replies/:id", blogController.UpdateReply)
		comments.DELETE("/replies/:id", blogController.DeleteReply)
		comments.GET("/replies/comments/:commentID", blogController.GetReplies)

		comments.POST("/comments/:id/like", blogController.LikeComment)
		comments.POST("/comments/:id/unlike", blogController.UnlikeComment)
		comments.POST("/replies/:id/like", blogController.LikeReply)
		comments.POST("/replies/:id/unlike", blogController.UnlikeReply)

	}

}
