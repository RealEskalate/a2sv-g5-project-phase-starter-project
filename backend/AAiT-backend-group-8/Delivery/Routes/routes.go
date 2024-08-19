package Router

import (
	"AAiT-backend-group-8/Delivery/Controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(userHandler *controllers.UserHandler, controller controllers.Controller) *gin.Engine {
	router := gin.Default()

	router.POST("/register", userHandler.RegisterUser)
	router.GET("/verify", userHandler.VerifyEmail)
	router.POST("/comment/:blogID", controller.CreateComment)
	router.POST("/login", userHandler.Login)

	return router
}
