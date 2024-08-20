package routers


import (
	"github.com/gin-gonic/gin"
	"AAIT-backend-group-3/internal/delivery/controllers"
)

func CreateUserRouter(router *gin.Engine, userController *controllers.UserController, otpController controllers.IOTPController){
	router.POST("/auth/sign-up", userController.Register)
	router.POST("/auth/sign-in", userController.Login)
	router.POST("/auth/refersh-token", userController.RefreshToken)
	router.POST("/auth/forgot-password", otpController.ForgotPassword)
	router.POST("/auth/reset-password", otpController.ResetPassword)
}