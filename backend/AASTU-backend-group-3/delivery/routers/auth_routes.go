package routers

import (
	"group3-blogApi/config/db"
	"group3-blogApi/delivery/controllers"
	"group3-blogApi/infrastracture" // Corrected import path
	"group3-blogApi/repository"
	"group3-blogApi/usecase"

	"github.com/gin-gonic/gin"
)

func SetUpAuth(router *gin.Engine) {
	// Initialize repository
	userRepo := repository.NewUserRepositoryImpl(db.UserCollection)

	// Initialize token generator and password service
	tokenGen := infrastracture.NewTokenGenerator() 
	passwordSvc := infrastracture.NewPasswordService()

	// Initialize usecase with dependencies
	userUsecase := usecase.NewUserUsecase(userRepo, tokenGen, passwordSvc)

	// Initialize controller with usecase
	authController := controllers.NewUserController(userUsecase)

	// Set up auth routes
	auth := router.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
		auth.POST("/activate", authController.ActivateAccount)

		

		// OAuth routes
		auth.GET("/login/google", authController.HandleGoogleLogin)
		auth.GET("/callback", authController.HandleGoogleCallback)

		// Password reset routes
		auth.POST("/reset-password", authController.SendPasswordResetLink)
		auth.POST("/reset-password/:token", authController.ResetPassword)
	}
}
