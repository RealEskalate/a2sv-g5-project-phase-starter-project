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

	blogRepo := repositories.NewBlogRepository(db, os.Getenv("BLOG_COLLECTION"))
	blogUseCase := usecases.NewBlogUseCase(blogRepo)
	blogController := controllers.NewBlogController(blogUseCase)

	commentRepo := repositories.NewCommentRepository(db, os.Getenv("COMMENT_COLLECTION_NAME"))
	commentController := controllers.CommentController{
		CommentUsecase: usecases.NewCommentUsecase(commentRepo),
	}

	router.POST("/blogs", blogController.CreateBlog)
	router.GET("/blogs", blogController.GetAllBlogs)
	router.GET("/blogs/:id", blogController.GetBlogByID)
	router.PUT("/blogs/:id", blogController.UpdateBlog)
	router.DELETE("/blogs/:id", blogController.DeleteBlog)
	router.PATCH("/blogs/:id/view", blogController.AddView)
	router.GET("/blogs/search", blogController.SearchBlogs)

	router.GET("/comment/:blog_id", commentController.GetComments)
	router.GET("/comment_count/:blog_id", commentController.GetCommentsCount)
	router.POST("/comment", commentController.AddComment)
	router.PUT("/comment/:id", commentController.UpdateComment)
	router.DELETE("/comment/:id", commentController.DelelteComment)
	port := os.Getenv("PORT")
	router.Run(":" + port)
}
