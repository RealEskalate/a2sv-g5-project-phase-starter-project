package routes

import (
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
)


func RegisterBlogRoutes(blogsCollection, commentsCollection, likesCollection, viewsCollection, tagsCollection *mongo.Collection, router *gin.Engine) {
	blogHandler := InstantaiteBlogHandler(blogsCollection, commentsCollection, likesCollection, viewsCollection, tagsCollection)
	blogRoutes := router.Group("/api/v1/blog")
	tagRoutes := router.Group("/api/v1/tag")

	{
		blogRoutes.POST("/", blogHandler.CreateBlogHandler)
		blogRoutes.GET("/:id", blogHandler.GetBlogByIDHandler)
		blogRoutes.PUT("/:id", blogHandler.UpdateBlogHandler)
		blogRoutes.DELETE("/:id", blogHandler.DeleteBlogHandler)
		blogRoutes.GET("/", blogHandler.GetAllBlogsHandler)
		blogRoutes.GET("/filter", blogHandler.FilterBlogsHandler)
		blogRoutes.GET("/paginate", blogHandler.PaginateBlogsHandler)
		blogRoutes.POST("/:id/tags", blogHandler.AddTagToBlogHandler)
		blogRoutes.DELETE("/:id/tags/:tagId", blogHandler.RemoveTagFromBlogHandler)
		blogRoutes.POST("/:id/comments", blogHandler.AddCommentHandler)
		blogRoutes.GET("/:id/comments", blogHandler.GetCommentsByBlogIDHandler)
		blogRoutes.POST("/:id/likes", blogHandler.AddLikeHandler)
		blogRoutes.GET("/:id/likes", blogHandler.GetLikesByBlogIDHandler)
		blogRoutes.POST("/:id/views", blogHandler.AddViewHandler)
		blogRoutes.GET("/:id/views", blogHandler.GetViewsByBlogIDHandler)
		
	}
	{
		tagRoutes.POST("/", blogHandler.CreateTagHandler)
		tagRoutes.GET("/:id", blogHandler.GetTagByIDHandler)
		tagRoutes.PUT("/:id", blogHandler.UpdateTagHandler)
		tagRoutes.DELETE("/:id", blogHandler.DeleteTagHandler)
		tagRoutes.GET("/", blogHandler.GetAllTagsHandler)
	}
}