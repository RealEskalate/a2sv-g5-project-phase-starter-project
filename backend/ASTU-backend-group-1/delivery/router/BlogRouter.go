package router

import (
	"astu-backend-g1/delivery/controllers"

	"github.com/gin-gonic/gin"
)

type BlogRoute struct {
	usecase controllers.BlogController
}

func NewBlogRoute(usecase controllers.BlogController) *BlogRoute {
	return &BlogRoute{usecase: usecase}
}

func (gr *BlogRoute) GinBlogRouter() {
	router := gin.Default()
	blogRouter := router.Group("/blog")
	{
		blogRouter.POST("/new", gr.usecase.HandleCreateBlog)
		blogRouter.GET("/all", gr.usecase.HandleGetAllBlogs)
		blogRouter.GET("/filter", gr.usecase.HandleFilterBlogs)
		blogRouter.PATCH("/:blogId", gr.usecase.HandleUpdate)
		blogRouter.DELETE("/:blogId", gr.usecase.HandleDelete)
		blogRouter.DELETE("/like/:blogId", gr.usecase.HandleBlogLikeOrDislike)
		blogRouter.DELETE("/dislike/:blogId", gr.usecase.HandleBlogLikeOrDislike)
		blogRouter.DELETE("/view/:blogId", gr.usecase.HandleBlogLikeOrDislike)
	}
	commentRouter := router.Group("/comment")
	{
		commentRouter.POST("/:blogId", gr.usecase.HandleCommentOnBlog)
		commentRouter.GET("/:blogId/like:commentId", gr.usecase.HandleCommentLikeOrDislike)
		commentRouter.GET("/:blogId/dislike:commentId", gr.usecase.HandleCommentLikeOrDislike)
		commentRouter.GET("/:blogId/View:commentId", gr.usecase.HandleCommentLikeOrDislike)
		// commentRouter.GET("/:blogId/reply:commentId",gr.usecase.Handle)
	}
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Blog API get"})
		ctx.Abort()
		return
	})
	router.POST("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Blog API create"})
		ctx.Abort()
		return
	})
	router.DELETE("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Blog API delete"})
		ctx.Abort()
		return
	})
	router.PATCH("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Blog API patch"})
		ctx.Abort()
		return
	})
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "Not such route"})
		ctx.Abort()
		return
	})
	router.Run(":9090")
}
