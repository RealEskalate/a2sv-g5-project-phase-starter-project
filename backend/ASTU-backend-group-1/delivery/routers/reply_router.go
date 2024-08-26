package routers

import "github.com/gin-gonic/gin"

func (gr *MainRouter) addReplyRouter(commentRouter *gin.RouterGroup) *gin.RouterGroup {
	repliesRouter := commentRouter.Group("/:commentId/replies")
	repliesRouter.Use(gr.authController.USERMiddleware())
	{
		repliesRouter.GET("/", gr.blogController.HandleGetAllRepliesForComment)
		repliesRouter.POST("/", gr.blogController.HandleReplyOnComment)
		repliesRouter.GET("/:replyId", gr.blogController.HandleGetReplyById)
		repliesRouter.POST("/:replyId/:type", gr.blogController.HandleReplyLikeOrDislike)
	}
	return repliesRouter
}
