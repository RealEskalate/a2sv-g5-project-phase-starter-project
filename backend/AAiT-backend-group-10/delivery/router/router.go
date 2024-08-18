package router

import (
	"aait.backend.g10/delivery/controllers"
	"aait.backend.g10/repositories"
	"aait.backend.g10/usecases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(db *mongo.Database) {
	router := gin.Default()
	blogRepo := repositories.NewBlogRepository(db, "blogs")
	blogUseCase := usecases.NewBlogUseCase(blogRepo)
	blogController := controllers.NewBlogController(blogUseCase)

	router.POST("/blogs", blogController.CreateBlog)
	router.GET("/blogs", blogController.GetAllBlogs)
	router.GET("/blogs/:id", blogController.GetBlogByID)
	router.PUT("/blogs/:id", blogController.UpdateBlog)
	router.DELETE("/blogs/:id", blogController.DeleteBlog)
	router.PATCH("/blogs/:id/view", blogController.AddView)
	router.GET("/blogs/search", blogController.SearchBlogs)

	router.Run(":8080")
}