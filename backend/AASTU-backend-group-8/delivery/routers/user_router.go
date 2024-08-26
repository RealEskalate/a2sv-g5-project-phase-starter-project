package routers

import (
	"github.com/gin-gonic/gin"
	"meleket/delivery/controllers"
	"meleket/infrastructure"
	"meleket/usecases"
)

func NewUserRouter(r *gin.Engine, userUsecase *usecases.UserUsecase, jwtService infrastructure.JWTService, otpUsecase *usecases.OTPUsecase) {
	userController := controllers.NewUserController(userUsecase)
	r.POST("/login", userController.Login)
	forgotPasswordController := controllers.NewForgotPasswordController(userUsecase, otpUsecase)
	r.POST("/forgotpassword", forgotPasswordController.ForgotPassword)
	r.POST("/verfiyforgotpassword", forgotPasswordController.VerifyForgotOTP)
	auth := r.Group("/api")
	auth.Use(infrastructure.AdminMiddleware(jwtService))
	{
		// Admin-specific routes
		auth.GET("/getallusers", userController.GetAllUsers)
		auth.DELETE("/deleteuser/:id", userController.DeleteUser)
		// Admin-specific routes
		auth.POST("/getallusers", userController.GetAllUsers)
		auth.PUT("/deleteusers/:id", userController.DeleteUser)
	}
}
