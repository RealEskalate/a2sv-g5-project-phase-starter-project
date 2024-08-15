package routes

import (
	"blogApp/internal/http/handlers"
	"blogApp/internal/usecase/user"

	"github.com/gin-gonic/gin"
)

func RegisterVerificationRoutes(router *gin.Engine) {
	userUsecase := user.NewUserUsecase()
	userHandler := handlers.NewUserHandler(userUsecase)

	router.POST("/request-verification-email", userHandler.RequestVerifyEmail)
	router.GET("/verify-email", userHandler.VerifyEmail)

	router.POST("/reset-password-request", userHandler.ResetPasswordRequest)
	router.GET("/reset-password", userHandler.ResetPassword)
}
