package routers

import (
	"meleket/delivery/controllers"
	"meleket/domain"
	"meleket/infrastructure"

	"github.com/gin-gonic/gin"
)

func NewCommentRouter(r *gin.Engine,
	commentUsecase domain.CommentUsecaseInterface,
	otpUsecase domain.OTPUsecaseInterface,
	jwtService infrastructure.JWTService,
) {

	commentController := controllers.NewCommentController(commentUsecase)
	auth := r.Group("/api")
	auth.Use(infrastructure.AdminMiddleware(jwtService))
	{
		// Comment routes
		auth.POST("/blogs/:id/comments", commentController.AddComment)
		auth.GET("/blogs/:id/comments", commentController.GetCommentsByBlogID)
		auth.PUT("/comments/:id", commentController.UpdateComment)
		auth.DELETE("/comments/:id", commentController.DeleteComment)
	}

}
