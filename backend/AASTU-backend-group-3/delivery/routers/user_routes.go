package routers

import (
	"group3-blogApi/config/db"
	"group3-blogApi/delivery/controllers"
	"group3-blogApi/infrastracture"
	"group3-blogApi/repository"
	"group3-blogApi/usecase"

	"github.com/gin-gonic/gin"
)

func SetUpUser(router *gin.Engine) {
	//user routes
	userRepo := repository.NewUserRepositoryImpl(db.UserCollection)
    userUsecase := usecase.NewUserUsecase(userRepo)
    authController := controllers.NewUserController(userUsecase)
	user := router.Group("/user")
	user.Use(infrastracture.AuthMiddleware())

	{
		user.GET("/me",  authController.GetMyProfile)
		user.PUT("/update", authController.UpdateMyProfile)
		user.POST("/upload-image", authController.UploadImage)
		user.DELETE("/me",  authController.DeleteMyAccount)
	


		// Logout Routes
	
		user.POST("/refresh-token", authController.RefreshToken)
		user.POST("/logout", authController.Logout)
		user.GET("logout-all", authController.LogoutAll)
		user.GET("/devices/logout", authController.LogoutDevice)
		user.GET("/devices", authController.GetDevices)

	}
}
