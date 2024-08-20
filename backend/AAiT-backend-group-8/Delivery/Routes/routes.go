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
	router.POST("/refresh",userHandler.RefreshToken)
    router.POST("/forgot-password", userHandler.ForgotPassword)
	router.GET("/store-token", userHandler.StoreToken) 
	router.POST("/reset-password", userHandler.ResetPassword)

	return router
}
