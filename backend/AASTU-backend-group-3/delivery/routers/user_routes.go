package routers

import (
	"github.com/gin-gonic/gin"

	"group3-blogApi/config/db"
	"group3-blogApi/delivery/controllers"
	"group3-blogApi/infrastracture"
	"group3-blogApi/repository"
	"group3-blogApi/usecase"
)

func SetUpUser(router *gin.Engine) {
	// Initialize repository
	userRepo := repository.NewUserRepositoryImpl(db.UserCollection)

	// Initialize token generator and password service
	tokenGen := infrastracture.NewTokenGenerator() 
	passwordService := infrastracture.NewPasswordService()

	// Initialize usecase with token generator and password service
	userUsecase := usecase.NewUserUsecase(userRepo, tokenGen, passwordService)

	// Initialize controller with usecase
	authController := controllers.NewUserController(userUsecase)

	// Set up user routes
	user := router.Group("/user")
	user.POST("/refresh-token", authController.RefreshToken)
	user.Use(infrastracture.AuthMiddleware()) 


	// Protected routes
	{
		user.GET("/me", authController.GetMyProfile)
		user.PUT("/update", authController.UpdateMyProfile)
		user.POST("/upload-image", authController.UploadImage)
		user.DELETE("/me", authController.DeleteMyAccount)


		user.GET("/activate/me", authController.ActivateAccountMe)


		// Logout routes
		user.POST("/logout", authController.Logout)
		user.GET("/logout-all", authController.LogoutAll)
		user.GET("/devices/logout", authController.LogoutDevice)
		user.GET("/devices", authController.GetDevices)
	}
}
