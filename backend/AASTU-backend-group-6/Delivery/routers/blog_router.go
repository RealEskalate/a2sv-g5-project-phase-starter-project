package routers

import (
	"blogs/Delivery/controllers"

	"github.com/gin-gonic/gin"
)

func NewBlogrouter(blogRouter *gin.RouterGroup, controller controllers.BlogController) {
	// unprotected
	blogRouter.GET("/")
	blogRouter.GET("/:id")

	blogRouter.GET("/search/:title&:author")
	blogRouter.GET("/tag/:tag")

	// protected
	blogRouter.GET("/my")
	blogRouter.GET("/my/:id")

	blogRouter.POST("/create")
	blogRouter.PUT("/update/:id")
	blogRouter.DELETE("/delete/:id")
	blogRouter.POST("/comment/create")
}
