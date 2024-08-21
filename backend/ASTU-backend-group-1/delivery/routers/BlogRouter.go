package router

import (
	"astu-backend-g1/delivery/controllers"

	"github.com/gin-gonic/gin"
)

type BlogRoute struct {
	usecase        controllers.BlogController
	authController controllers.GeneralAuthorizationController
}

func NewBlogRoute(usecase controllers.BlogController, authcontroller controllers.GeneralAuthorizationController) *BlogRoute {
	return &BlogRoute{usecase: usecase, authController: authcontroller}
}

func (gr *BlogRoute) GinBlogRouter() {
	router := gin.Default()
	blogRouter := router.Group("/blogs")
	{
		blogRouter.POST("/", gr.usecase.HandleCreateBlog)
		blogRouter.GET("/", gr.usecase.HandleGetAllBlogs)
		blogRouter.GET("/popular", gr.usecase.HandleGetPopularBlog)
		blogRouter.GET("/filter", gr.usecase.HandleFilterBlogs)
		blogRouter.GET("/:blogId", gr.usecase.HandleGetBlogById)
		blogRouter.PATCH("/:blogId", gr.usecase.HandleBlogUpdate)
		blogRouter.DELETE("/:blogId", gr.authController.AdminMiddlewareGin(), gr.usecase.HandleBlogDelete)
		blogRouter.POST("/:blogId/interact/:type", gr.usecase.HandleBlogLikeOrDislike)

		// TODO: check if there is a blog with such id
		commentRouter := blogRouter.Group("/:blogId/comments")
		{
			commentRouter.GET("/", gr.usecase.HandleGetAllComments)
			commentRouter.POST("/", gr.usecase.HandleCommentOnBlog)
			commentRouter.GET("/:commentId", gr.usecase.HandleGetCommentById)
			commentRouter.POST("/:commentId/interact/:type", gr.usecase.HandleCommentLikeOrDislike)

			repliesRouter := commentRouter.Group("/:commentId/replies")
			{
				repliesRouter.GET("/", gr.usecase.HandleGetAllRepliesForComment)
				repliesRouter.POST("/", gr.usecase.HandleReplyOnComment)
				repliesRouter.GET("/:replyId", gr.usecase.HandleGetReplyById)
				// todo: test the below functions
				repliesRouter.POST("/:replyId/interact/:type", gr.usecase.HandleReplyLikeOrDislike)
			}
		}

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
