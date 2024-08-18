package router

import (
	"aait.backend.g10/delivery/controllers"
	"aait.backend.g10/domain"
	"aait.backend.g10/repositories"
	"aait.backend.g10/usecases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(db mongo.Database) {
	router := gin.Default()
	// router.Use(cors.Default())
	// router.Use(middleware.AuthMiddleware())
	likeRepo := repositories.NewLikeRepository(db, domain.CollectionLike)
	likeController := controllers.LikeCOntroller{
		LikeUseCase: usecases.NewLikeUseCase(likeRepo),
	}
	router.PUT("/like", likeController.LikeBlog)
	router.DELETE("/like", likeController.DeleteLike)
	router.Run(":8080")
}
