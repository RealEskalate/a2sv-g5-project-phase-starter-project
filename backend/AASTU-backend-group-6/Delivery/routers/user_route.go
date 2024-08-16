package routers

import (
	infrastructure "blogs/Infrastructure"
	"blogs/mongo"

	"github.com/gin-gonic/gin"
)

func NewUserrouter(config *infrastructure.Config, DB mongo.Database, blogRouter *gin.RouterGroup) {
	blogRouter.POST("/create")
	blogRouter.GET("/get")
	blogRouter.GET("/get/:id")
	blogRouter.PUT("/update/:id")
	blogRouter.DELETE("/delete/:id")
	blogRouter.POST("/comment/:id")

}
