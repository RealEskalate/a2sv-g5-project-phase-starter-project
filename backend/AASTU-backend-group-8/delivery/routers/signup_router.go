package routers

import (
	"meleket/delivery/controllers"
	"meleket/usecases"

	"github.com/gin-gonic/gin"
)

func NewSignupRouter(r *gin.Engine, userUsecase *usecases.UserUsecase, otpUsecase *usecases.OTPUsecase) {
	// Initialize controllers
	signupController := controllers.NewSignupController(userUsecase, otpUsecase)
	// Public routes
	r.POST("/signup", signupController.Signup)
	r.POST("/verify", signupController.VerifyOTP)
}
