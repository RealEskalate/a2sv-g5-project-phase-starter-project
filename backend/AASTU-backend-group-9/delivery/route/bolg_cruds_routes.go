package route

import (
	"blog/config"
	"blog/database"
	"blog/delivery/controller"

	// "go/doc/comment"

	// "blog/domain"
	"blog/repository"
	"blog/usecase"
	"time"

	"blog/delivery/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(env *config.Env, timeout time.Duration, db database.Database, router *gin.RouterGroup) {

	blogRepo := repository.NewBlogRepository(db, "blogs")
	popuRepo := repository.NewPopularityRepository(db, "popularity")
	commentRepo := repository.NewCommentRepository(db, "comments")
	// blogUse := usecase.NewBlogUsecase(blogRepo, timeout)
	// popuUse := usecase.NewPopularityUsecase(popuRepo, timeout)
	blogController := &controller.BlogController{
		BlogUsecase: usecase.NewBlogUsecase(blogRepo, popuRepo, commentRepo, timeout),
		Env:         env,
	}
	blogRoutes := router.Group("/blogs")
	{
		blogRoutes.Use(middleware.AuthMidd) // Protect the routes with authentication middleware
		blogRoutes.POST("/", blogController.CreateBlog)
		blogRoutes.GET("/:id", blogController.GetBlogByID)
		blogRoutes.GET("/", blogController.GetAllBlogs)
		blogRoutes.PUT("/:id", blogController.UpdateBlog)
		blogRoutes.DELETE("/:id", blogController.DeleteBlog)
		blogRoutes.GET("/search", blogController.SearchBlogs)
		// blogRoutes.GET("/filter/tags", blogController.FilterBlogsByTags)
		// blogRoutes.GET("/filter/date", blogController.FilterBlogsByDate)
		blogRoutes.GET("/filter/popularity", blogController.FilterBlogs)
		blogRoutes.POST("/:id/view", blogController.TrackView)
		blogRoutes.POST("/:id/like", blogController.TrackLike)
		blogRoutes.POST("/:id/dislike", blogController.TrackDislike)

	}
	blogCommentRoutes := router.Group("blogs/:id/comment")

	{
		blogCommentRoutes.Use(middleware.AuthMidd)
		blogCommentRoutes.POST("/", blogController.AddComment)
		blogCommentRoutes.POST("/:comment_id/reply", blogController.AddReply)
		blogCommentRoutes.GET("/", blogController.GetComments)
		blogCommentRoutes.DELETE("/:comment_id", blogController.DeleteComment)
		blogCommentRoutes.PATCH("/:comment_id", blogController.UpdateComment)
		blogCommentRoutes.POST("/:comment_id/popularity", blogController.TrackCommentPopularity)

	}

}
