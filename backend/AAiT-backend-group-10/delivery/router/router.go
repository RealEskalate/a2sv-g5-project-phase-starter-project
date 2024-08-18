package router

import (
	"os"

	"aait.backend.g10/delivery/controllers"
	"aait.backend.g10/repositories"
	"aait.backend.g10/usecases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(db mongo.Database) {
	router := gin.Default()

	likeRepo := repositories.NewLikeRepository(db, os.Getenv("LIKE_COLLECTION"))
	likeController := controllers.LikeCOntroller{
		LikeUseCase: usecases.NewLikeUseCase(likeRepo),
	}
	router.PUT("/like", likeController.LikeBlog)
	router.DELETE("/like", likeController.DeleteLike)
	port := os.Getenv("PORT")
	router.Run(":" + port)
}
