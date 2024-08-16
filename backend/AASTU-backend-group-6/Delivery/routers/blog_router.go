package routers

import (
	infrastructure "blogs/Infrastructure"
	"blogs/mongo"

	"github.com/gin-gonic/gin"
)

func NewBlogrouter(config *infrastructure.Config, DB mongo.Database, blogRouter *gin.RouterGroup) {
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
