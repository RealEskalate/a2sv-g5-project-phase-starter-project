package routers

import (
	infrastructure "blogs/Infrastructure"
	"blogs/mongo"

	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine, config *infrastructure.Config, DB mongo.Database) {
	blogRouter := server.Group("blogs")
	NewBlogrouter(config, DB, blogRouter)
	userRouter := server.Group("")
	NewUserrouter(config, DB, userRouter)

}
