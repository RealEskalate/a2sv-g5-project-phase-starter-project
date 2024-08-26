package routers

import "github.com/gin-gonic/gin"

func (gr *MainRouter) addCommentRouter(blogRouter *gin.RouterGroup)  *gin.RouterGroup{
	commentRouter := blogRouter.Group("/:blogId/comments")
	commentRouter.Use(gr.authController.USERMiddleware())
	{
		commentRouter.GET("/", gr.blogController.HandleGetAllComments)
		commentRouter.POST("/", gr.blogController.HandleCommentOnBlog)
		commentRouter.GET("/:commentId", gr.blogController.HandleGetCommentById)
		commentRouter.POST("/:commentId/:type", gr.blogController.HandleCommentLikeOrDislike)
		
		// commentRouter.DELETE("/:commentId", gr.blogController.HandleGetCommentById)
		// commentRouter.PATCH("/:commentId", gr.blogController.HandleGetCommentById)
	}
	return commentRouter
}
