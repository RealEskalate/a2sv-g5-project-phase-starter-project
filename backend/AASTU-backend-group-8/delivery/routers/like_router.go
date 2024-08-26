package routers

import (
	"meleket/delivery/controllers"
	"meleket/domain"
	"meleket/infrastructure"

	"github.com/gin-gonic/gin"
)

func NewLikeRouter(r *gin.Engine,
	likeUsecase domain.LikeUsecaseInterface,
	otpUsecase domain.OTPUsecaseInterface,
	jwtService infrastructure.JWTService) {
	likeController := controllers.NewLikeController(likeUsecase)

	// Authenticated routes
	auth := r.Group("/api")
	auth.Use(infrastructure.AdminMiddleware(jwtService))
	{
		// Like routes
		auth.POST("/blogs/:id/likes", likeController.AddLike)
		auth.GET("/blogs/:id/likes", likeController.GetLikesByBlogID)
		auth.DELETE("/likes/:id", likeController.RemoveLike)
	}

}
