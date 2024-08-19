package router

import (
	"os"

	"aait.backend.g10/delivery/controllers"
	"aait.backend.g10/repositories"
	"aait.backend.g10/usecases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(db *mongo.Database) {
	router := gin.Default()

	commentRepo := repositories.NewCommentRepository(db, os.Getenv("COMMENT_COLLECTION_NAME"))
	commentController := controllers.CommentController{
		CommentUsecase: usecases.NewCommentUsecase(commentRepo),
	}
	router.GET("/comment/:blog_id", commentController.GetComments)
	router.GET("/comment_count/:blog_id", commentController.GetCommentsCount)
	router.POST("/comment", commentController.AddComment)
	router.PUT("/comment/:id", commentController.UpdateComment)
	router.DELETE("/comment/:id", commentController.DelelteComment)
	port := os.Getenv("PORT")
	router.Run(":" + port)
}
