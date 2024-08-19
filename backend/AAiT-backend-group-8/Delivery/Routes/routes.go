package Router

import (
	"AAiT-backend-group-8/Delivery/Controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(userHandler *Controller.UserHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/register", userHandler.RegisterUser)
	router.GET("/verify", userHandler.VerifyEmail)
	router.POST("/login", userHandler.Login)

	return router
}
