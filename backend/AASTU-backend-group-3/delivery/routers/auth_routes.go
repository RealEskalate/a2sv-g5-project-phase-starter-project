package routers

import (
	"group3-blogApi/config/db"
	"group3-blogApi/delivery/controllers"
	"group3-blogApi/repository"
	"group3-blogApi/usecase"

	"github.com/gin-gonic/gin"
)

func SetUpAuth(router *gin.Engine) {	


	userRepo := repository.NewUserRepositoryImpl(db.UserCollection)
    userUsecase := usecase.NewUserUsecase(userRepo)
    authController := controllers.NewUserController(userUsecase)

	auth := router.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
		auth.GET("/activate", authController.ActivateAccount)
		
		// auth.POST("/continueWithGoogle", authController.continueWithGoogle)


	}
}