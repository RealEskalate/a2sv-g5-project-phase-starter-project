package router

import (
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/gin-gonic/gin"
)

func SetupRouter(blogController domain.BlogController) *gin.Engine {
	router := gin.Default()

	// Blog routes
	blogRoutes := router.Group("/blogs")
	{
		blogRoutes.POST("/", blogController.CreateBlog)
		blogRoutes.GET("/:id", blogController.GetBlog)
		blogRoutes.GET("/", blogController.GetBlogs)
		blogRoutes.PUT("/:id", blogController.UpdateBlog)
		blogRoutes.DELETE("/:id", blogController.DeleteBlog)
		blogRoutes.GET("/search/title", blogController.SearchBlogsByTitle)
		blogRoutes.GET("/search/author", blogController.SearchBlogsByAuthor)
		blogRoutes.GET("/filter", blogController.FilterBlogs)
		blogRoutes.POST("/:id/like", blogController.LikeBlog)
		blogRoutes.POST("/:id/dislike", blogController.DislikeBlog)
		blogRoutes.POST("/:id/comments", blogController.AddComment)
		blogRoutes.DELETE("/:id/comments/:comment_id", blogController.DeleteComment)
		blogRoutes.PUT("/:id/comments/:comment_id", blogController.EditComment)
	}

	return router
}