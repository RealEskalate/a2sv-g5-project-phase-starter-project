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
	commentController := controllers.NewCommentController(commentUseCase)

	comments := router.Group("/")
	comments.Use(infrastracture.AuthMiddleware())
	{
		comments.POST("/comments",infrastracture.EligibilityMiddleware() ,commentController.CreateComment)
		comments.PUT("/comments/:id",infrastracture.EligibilityMiddleware() ,commentController.UpdateComment)
		comments.DELETE("/comments/:id", infrastracture.EligibilityMiddleware(),commentController.DeleteComment)
		comments.GET("/comments/:id", infrastracture.EligibilityMiddleware(),commentController.GetCommentByID)
		comments.GET("/posts/:postID/comments", infrastracture.EligibilityMiddleware(),commentController.GetComments)

		comments.POST("/replies", infrastracture.EligibilityMiddleware(),commentController.CreateReply)
		comments.PUT("/replies/:id", infrastracture.EligibilityMiddleware(),commentController.UpdateReply)
		comments.DELETE("/replies/:id", infrastracture.EligibilityMiddleware(),commentController.DeleteReply)
		comments.GET("/replies/comments/:commentID", infrastracture.EligibilityMiddleware(),commentController.GetReplies)

		comments.POST("/comments/:id/like", infrastracture.EligibilityMiddleware(),commentController.LikeComment)
		comments.POST("/comments/:id/unlike",infrastracture.EligibilityMiddleware() ,commentController.UnlikeComment)
		comments.POST("/replies/:id/like", infrastracture.EligibilityMiddleware(),commentController.LikeReply)
		comments.POST("/replies/:id/unlike", infrastracture.EligibilityMiddleware(),commentController.UnlikeReply)

	}

}
