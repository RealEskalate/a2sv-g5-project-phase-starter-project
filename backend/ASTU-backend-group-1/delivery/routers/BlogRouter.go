package router

import (
	"astu-backend-g1/delivery/controllers"

	"github.com/gin-gonic/gin"
)

type BlogRoute struct {
	usecase controllers.BlogController
	authController controllers.GeneralAuthorizationController
}

func NewBlogRoute(usecase controllers.BlogController,authcontroller controllers.GeneralAuthorizationController) *BlogRoute {
	return &BlogRoute{usecase: usecase,authController: authcontroller}
}

func (gr *BlogRoute) GinBlogRouter() {
	router := gin.Default()
	blogRouter := router.Group("/blog")
	{
		// INFO:PASSED
		blogRouter.POST("/new", gr.authController.UserMiddlewareGin(), gr.usecase.HandleCreateBlog)
		blogRouter.GET("/all", gr.authController.UserMiddlewareGin(), gr.usecase.HandleGetAllBlogs)
		blogRouter.GET("/popular", gr.authController.UserMiddlewareGin(), gr.usecase.HandleGetPopularBlog)
		blogRouter.GET("/filter", gr.authController.UserMiddlewareGin(), gr.usecase.HandleFilterBlogs)
		blogRouter.PATCH("/:blogId", gr.authController.UserMiddlewareGin(), gr.usecase.HandleUpdate)
		blogRouter.DELETE("/:blogId", gr.authController.AdminMiddlewareGin(), gr.usecase.HandleDelete)
		blogRouter.PATCH("/interact/:type", gr.authController.UserMiddlewareGin(), gr.usecase.HandleBlogLikeOrDislike)

	}
	commentRouter := router.Group("/comment")
	{
		// TODO: check if there is a blog with such id
		commentRouter.PATCH("/new", gr.authController.UserMiddlewareGin(), gr.usecase.HandleCommentOnBlog)

		commentRouter.PATCH("/interact/:type", gr.authController.UserMiddlewareGin(), gr.usecase.HandleCommentLikeOrDislike)
		commentRouter.GET("/:blogId", gr.authController.UserMiddlewareGin(), gr.usecase.HandleGetAllComments)
		// INFO:TESTING
		// INFO:TOBE TESTED
		// commentRouter.PATCH("/:comment/interact:commentId", gr.usecase.HandleCommentLikeOrDislike)
		// commentRouter.PATCH("/:comment/interact:commentId", gr.usecase.HandleCommentLikeOrDislike)
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
