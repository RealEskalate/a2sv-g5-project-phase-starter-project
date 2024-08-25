package routes

import (
	"blogApp/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(router *gin.Engine) {
	blogHandler := InstantaiteBlogHandler()
	blogRoutes := router.Group("/api/v1/blog", middleware.AuthMiddleware())
	tagRoutes := router.Group("/api/v1/tag", middleware.AuthMiddleware())
	aiRoutes := router.Group("/api/v1/ai", middleware.AuthMiddleware())

	publicBlogRoutes := router.Group("/api/v1/blog")

	{
		blogRoutes.POST("/", blogHandler.CreateBlogHandler)
		blogRoutes.GET("/", blogHandler.GetMyBlogsHandler)
		blogRoutes.PUT("/:id", blogHandler.UpdateBlogHandler)
		blogRoutes.DELETE("/:id", blogHandler.DeleteBlogHandler)

		// blogRoutes.GET("/filter", blogHandler.FilterBlogsHandler)
		blogRoutes.POST("/:id/tags", blogHandler.AddTagToBlogHandler)
		blogRoutes.DELETE("/:id/tags/:id", blogHandler.RemoveTagFromBlogHandler)
		blogRoutes.POST("/comments", blogHandler.AddCommentHandler)

		blogRoutes.POST("/likes", blogHandler.AddLikeHandler)
		blogRoutes.GET("/:id/likes", blogHandler.GetLikesByBlogIDHandler)
		blogRoutes.POST("/views", blogHandler.AddViewHandler)
		blogRoutes.DELETE("/comments/:id", blogHandler.DeleteCommentHandler)
		blogRoutes.DELETE("/likes/:id", blogHandler.DeleteLikeHandler)
		blogRoutes.POST("/uploads", blogHandler.UploadBlogPhotos)
		// blogRoutes.GET("")

	}
	{
		publicBlogRoutes.GET("/:id", blogHandler.GetBlogByIDHandler)
		publicBlogRoutes.GET("/paginate", blogHandler.PaginateBlogsHandler)
		// publicBlogRoutes.GET("/", blogHandler.GetAllBlogsHandler)
		publicBlogRoutes.GET("/:id/comments", blogHandler.GetCommentsByBlogIDHandler)
		publicBlogRoutes.GET("/search", blogHandler.SearchBlogsHandler)
		publicBlogRoutes.GET("/:id/views", blogHandler.GetViewsByBlogIDHandler)
	}

	{
		tagRoutes.POST("/", middleware.AdminMiddleware(), blogHandler.CreateTagHandler)
		tagRoutes.PUT("/:id", middleware.AdminMiddleware(), blogHandler.UpdateTagHandler)
		tagRoutes.DELETE("/:id", middleware.AdminMiddleware(), blogHandler.DeleteTagHandler)

		tagRoutes.GET("/:id", blogHandler.GetTagByIDHandler)
		tagRoutes.GET("/", blogHandler.GetAllTagsHandler)
	}
	{
		aiRoutes.POST("/blog_assistant", blogHandler.GetAiBlog)
		aiRoutes.POST("/moderate_blog", blogHandler.ModerateBlog)
	}
}
