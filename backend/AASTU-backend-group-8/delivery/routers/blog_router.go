package routers

import (
	"github.com/gin-gonic/gin"
	"meleket/delivery/controllers"
	"meleket/infrastructure"
	"meleket/usecases"
)

func NewBlogRouter(r *gin.Engine, blogUsecase *usecases.BlogUsecase, jwtService infrastructure.JWTService) {
	blogController := controllers.NewBlogController(blogUsecase)
	r.GET("/blogs/:id", blogController.GetBlogByID)
	r.GET("/blogs", blogController.GetAllBlogPosts)
	auth := r.Group("/api")
	// Authenticated routes
	auth.Use(infrastructure.AuthMiddleware(jwtService))
	{
		// Blog routes
		auth.POST("/blogs", blogController.CreateBlogPost)
		auth.PUT("/blogs/:id", blogController.UpdateBlogPost)
		auth.DELETE("/blogs/:id", blogController.DeleteBlogPost)
		// Like routes
		auth.POST("/blogs/:id/likes", blogController.LikeBlogPost)
		auth.DELETE("/likes/:id", blogController.DislikeBlogPost)
	}
}
