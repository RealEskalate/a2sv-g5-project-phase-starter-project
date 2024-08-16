package routers

import (
	"blogs/delivery/controllers"

	"github.com/gin-gonic/gin"
)

func NewBlogrouter(blogRouter *gin.RouterGroup, controller controllers.BlogController) {
	// unprotected
	blogRouter.GET("/", controller.GetBlogs)
	blogRouter.GET("/:id")

	blogRouter.GET("/search", controller.SearchBlogByTitleAndAuthor)
	blogRouter.GET("/filter", controller.FilterBlogsByTag)

	// protected
	blogRouter.GET("/my")
	blogRouter.GET("/my/:id")

	blogRouter.POST("/create")
	blogRouter.PUT("/update/:id")
	blogRouter.DELETE("/delete/:id")
	blogRouter.POST("/comment/create")
}
