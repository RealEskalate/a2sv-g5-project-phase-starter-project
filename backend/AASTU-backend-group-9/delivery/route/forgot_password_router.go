package route

import (
	"blog/config"
	"blog/database"
	"blog/delivery/controller"
	"blog/repository"
	"blog/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewForgotPasswordRouter(env *config.Env, db database.Database, router *gin.RouterGroup) {
    otpRepo := repository.NewOTPRepository(db, "otp_collection")
    userRepo := repository.NewUserRepository(db, "users")
    forgotPasswordUsecase := usecase.NewForgotPasswordUsecase(userRepo, otpRepo, time.Minute*15)

    forgotPasswordController := &controller.ForgotPasswordController{
        ForgotPasswordUsecase: forgotPasswordUsecase,
        // Remove Env if not needed
    }

    router.POST("/forgot_password", forgotPasswordController.ForgotPassword)
    router.POST("/reset_password", forgotPasswordController.ResetPassword)
}
