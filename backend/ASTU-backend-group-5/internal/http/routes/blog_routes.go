package routes

import (
	"blogApp/internal/http/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterBlogRoutes(blogsCollection, commentsCollection, likesCollection, viewsCollection, tagsCollection *mongo.Collection, router *gin.Engine) {
	blogHandler := InstantaiteBlogHandler(blogsCollection, commentsCollection, likesCollection, viewsCollection, tagsCollection)
	blogRoutes := router.Group("/api/v1/blog", middleware.AuthMiddleware())
	tagRoutes := router.Group("/api/v1/tag", middleware.AuthMiddleware())
	aiRoutes := router.Group("/api/v1/ai", middleware.AuthMiddleware())

	{
		blogRoutes.POST("/", blogHandler.CreateBlogHandler)
		blogRoutes.GET("/:id", blogHandler.GetBlogByIDHandler)
		blogRoutes.PUT("/:id", blogHandler.UpdateBlogHandler)
		blogRoutes.DELETE("/:id", blogHandler.DeleteBlogHandler)
		blogRoutes.GET("/", blogHandler.GetAllBlogsHandler)
		// blogRoutes.GET("/filter", blogHandler.FilterBlogsHandler)
		blogRoutes.GET("/paginate", blogHandler.PaginateBlogsHandler)
		blogRoutes.POST("/:id/tags", blogHandler.AddTagToBlogHandler)
		blogRoutes.DELETE("/:id/tags/:id", blogHandler.RemoveTagFromBlogHandler)
		blogRoutes.POST("/comments", blogHandler.AddCommentHandler)
		blogRoutes.GET("/:id/comments", blogHandler.GetCommentsByBlogIDHandler)
		blogRoutes.GET("/search", blogHandler.SearchBlogsHandler)

		blogRoutes.POST("/likes", blogHandler.AddLikeHandler)
		blogRoutes.GET("/:id/likes", blogHandler.GetLikesByBlogIDHandler)
		blogRoutes.POST("/views", blogHandler.AddViewHandler)
		blogRoutes.GET("/:id/views", blogHandler.GetViewsByBlogIDHandler)
		blogRoutes.DELETE("/comments/:id", blogHandler.DeleteCommentHandler)
		blogRoutes.DELETE("/likes/:id", blogHandler.DeleteLikeHandler)
		blogRoutes.POST("/uploads", blogHandler.UploadBlogPhotos)
		// blogRoutes.GET("")

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
