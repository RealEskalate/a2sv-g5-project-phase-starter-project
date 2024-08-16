package routers

import (
	infrastructure "blogs/Infrastructure"
	"blogs/mongo"

	"github.com/gin-gonic/gin"
)

func NewUserrouter(config *infrastructure.Config, DB mongo.Database, userRouter *gin.RouterGroup) {
	// userRouter.POST("/create")
	userRouter.GET("/get")
	userRouter.GET("/get/:id")
	userRouter.PUT("/update/:id")
	userRouter.DELETE("/delete/:id")
	userRouter.POST("/comment/:id")

}
