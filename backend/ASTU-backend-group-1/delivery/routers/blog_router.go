package routers

import "github.com/gin-gonic/gin"

func (gr *MainRouter) addBlogRouter(generalRouter *gin.Engine) *gin.RouterGroup {
	generalRouter.GET("blogs/", gr.blogController.HandleGetAllBlogs)
	generalRouter.GET("blogs/popular", gr.blogController.HandleGetPopularBlog)
	generalRouter.GET("blogs/filter", gr.blogController.HandleFilterBlogs)
	blogRouter := generalRouter.Group("/blogs")
	{
		blogRouter.GET("/:blogId/summarize", gr.aiController.SumarizeBlog)
		blogRouter.GET("/:blogId/refine", gr.aiController.RefineBlog)
		blogRouter.GET("/:blogId/", gr.blogController.HandleGetBlogById)
	}
	blogRouter.Use(gr.authController.AuthenticationMiddleware())
	{
		blogRouter.POST("/", gr.authController.USERMiddleware(), gr.blogController.HandleCreateBlog)
		blogRouter.PATCH("/:blogId", gr.authController.OWNERMiddleware(), gr.blogController.HandleBlogUpdate)
		blogRouter.DELETE("/:blogId", gr.authController.OWNERMiddleware(), gr.blogController.HandleBlogDelete)
		blogRouter.POST("/:blogId/:type", gr.authController.USERMiddleware(), gr.blogController.HandleBlogLikeOrDislike)
	}
	return blogRouter
}
