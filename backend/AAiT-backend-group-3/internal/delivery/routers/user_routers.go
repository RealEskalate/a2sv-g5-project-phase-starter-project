package routers

import (
	"AAIT-backend-group-3/internal/delivery/controllers"
	"AAIT-backend-group-3/internal/infrastructures/middlewares"

	"github.com/gin-gonic/gin"
)

func CreateUserRouter(router *gin.Engine, userController controllers.UserControllerInterface, otpController controllers.IOTPController, authMiddleware middlewares.IAuthMiddleware) {
	router.POST("/auth/sign-up", userController.Register)
	router.POST("/auth/sign-in", userController.Login)
	router.GET("/auth/sign-out", authMiddleware.Authentication(), userController.Logout)
	router.GET("/auth/verify-email", userController.VerifyEmail)
	router.POST("/auth/refresh-token", userController.RefreshToken)
	router.POST("/auth/forgot-password", otpController.ForgotPassword)
	router.POST("/auth/reset-password", otpController.ResetPassword)

	router.POST("/user/profile_update", authMiddleware.Authentication(), userController.UpdateProfile)
}
