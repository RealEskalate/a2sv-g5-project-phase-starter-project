package router

import (
	infrastructure "astu-backend-g1/Infrastructure"
	"astu-backend-g1/delivery/controllers"

	"github.com/gin-gonic/gin"
)

type MainRouter struct {
	blogController controllers.BlogController
	authController infrastructure.GeneralAuthorizationController
	handler        controllers.UserController
}

func NewMainRouter(uc controllers.UserController, bc controllers.BlogController, authc infrastructure.GeneralAuthorizationController) *MainRouter {
	return &MainRouter{
		blogController: bc,
		authController: authc,
		handler:        uc,
	}
}
func (gr *MainRouter) GinBlogRouter() {
	router := gin.Default()
	userrouter := router.Group("/users")
	{

		userrouter.POST("/register", gr.handler.Register)
		userrouter.GET("/accountVerification", gr.handler.AccountVerification)
		userrouter.POST("/login", gr.handler.LoginUser)
		userrouter.GET("/forgetPassword", gr.handler.ForgetPassword)
		userrouter.POST("/resetPassword", gr.handler.ResetPassword)
	}
	blogRouter := router.Group("/blogs")
	{
		blogRouter.POST("/", gr.authController.UserMiddlewareGin(), gr.blogController.HandleCreateBlog)
		blogRouter.GET("/", gr.authController.UserMiddlewareGin(), gr.blogController.HandleGetAllBlogs)
		blogRouter.GET("/popular", gr.authController.UserMiddlewareGin(), gr.blogController.HandleGetPopularBlog)
		blogRouter.GET("/filter", gr.authController.UserMiddlewareGin(), gr.blogController.HandleFilterBlogs)
		blogRouter.GET("/:blogId", gr.authController.UserMiddlewareGin(), gr.blogController.HandleGetBlogById)
		blogRouter.PATCH("/:blogId", gr.authController.AuthorMiddlewareGin(), gr.blogController.HandleBlogUpdate)
		blogRouter.DELETE("/:blogId", gr.authController.AuthorMiddlewareGin(), gr.authController.AdminMiddlewareGin(), gr.blogController.HandleBlogDelete)
		blogRouter.POST("/:blogId/interact/:type", gr.authController.UserMiddlewareGin(), gr.blogController.HandleBlogLikeOrDislike)

		// TODO: check if there is a blog with such id
		commentRouter := blogRouter.Group("/:blogId/comments")
		{
			commentRouter.GET("/", gr.authController.UserMiddlewareGin(), gr.blogController.HandleGetAllComments)
			commentRouter.POST("/", gr.authController.UserMiddlewareGin(), gr.blogController.HandleCommentOnBlog)
			commentRouter.GET("/:commentId", gr.authController.UserMiddlewareGin(), gr.blogController.HandleGetCommentById)
			commentRouter.POST("/:commentId/interact/:type", gr.authController.UserMiddlewareGin(), gr.blogController.HandleCommentLikeOrDislike)

			repliesRouter := commentRouter.Group("/:commentId/replies")
			{
				repliesRouter.GET("/", gr.authController.UserMiddlewareGin(), gr.blogController.HandleGetAllRepliesForComment)
				repliesRouter.POST("/", gr.authController.UserMiddlewareGin(), gr.blogController.HandleReplyOnComment)
				repliesRouter.GET("/:replyId", gr.authController.UserMiddlewareGin(), gr.blogController.HandleGetReplyById)
				// todo: test the below functions
				repliesRouter.POST("/:replyId/interact/:type", gr.authController.UserMiddlewareGin(), gr.blogController.HandleReplyLikeOrDislike)
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
	router.Run(":8000")
}

// func (r *UserRoute) UserRoutes() *gin.RouterGroup {
// 	ro := gin.Default()
// 	userrouter := ro.Group("/user")
// 	userrouter.POST("/register", r.handler.Register)
// 	userrouter.GET("/verify?email=:email&pwd=:pwd", r.handler.AccountVerification)
// 	userrouter.POST("/login", r.handler.LoginUser)
// 	userrouter.GET("/forgetPassword", r.handler.ForgetPassword)
// 	userrouter.POST("/resetPassword?email=:email&pwd=:pwd", r.handler.ResetPassword)
// 	return userrouter
// }
