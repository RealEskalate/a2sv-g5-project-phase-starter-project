package routers

import (
	"group3-blogApi/delivery/controllers/authController"

	"github.com/gin-gonic/gin"
)

func SetUpAuth(router *gin.Engine) {

	
	auth := router.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
		auth.POST("/forgot-password", authController.ForgotPassword)
		auth.POST("/reset-password", authController.ResetPassword)
	}
}