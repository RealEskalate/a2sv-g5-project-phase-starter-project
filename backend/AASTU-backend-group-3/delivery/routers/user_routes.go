package routers

import (
	"group3-blogApi/config/db"
	"group3-blogApi/delivery/controllers"
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

	{
		// user.GET("/profile", authMiddleware, userController.Profile)
		// user.PUT("/update", authMiddleware, userController.Update)
		// user.POST("/upload-image", authMiddleware, userController.UploadImage)
		// user.POST("/logout", authMiddleware, authController.Logout)
		// user.POST("/reset-password", authMiddleware, authController.ResetPassword)
		user.POST("/refresh-token", authController.RefreshToken)
		user.POST("/logout", authController.Logout)
		user.GET("logout-all", authController.LogoutAll)
		user.GET("/devices/logout", authController.LogoutDevice)
		user.GET("/devices", authController.GetDevices)

	}
}
