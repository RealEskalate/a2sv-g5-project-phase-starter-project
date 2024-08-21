package routers

import (
	"group3-blogApi/config/db"
	"group3-blogApi/delivery/controllers"
	"group3-blogApi/infrastracture"
	"group3-blogApi/repository"
	"group3-blogApi/usecase"

	"github.com/gin-gonic/gin"
)

func SetUpLike(router *gin.Engine) {
	


	likeRepo := repository.NewLikeRepositoryImpl(db.LikeCollection)
	likeUsecase := usecase.NewLikeUsecase(likeRepo)
	likeController := controllers.NewLikeController(likeUsecase)



	like := router.Group("/blogs")
	like.Use(infrastracture.AuthMiddleware())
	{
			
		like.POST("/:id/like",  likeController.LikeBlog)
		like.POST("/:id/dislike",  likeController.DisLikeBlog)
	
	}
}