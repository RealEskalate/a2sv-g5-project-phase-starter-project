package routers

import (
	"github.com/gin-gonic/gin"
)


func SetupOauthRouter(router *gin.Engine) {

	// userRepo := repository.NewUserRepositoryImpl(db.UserCollection)
    // userUsecase := usecase.NewUserUsecase(userRepo)
    // userController := controllers.NewUserController(userUsecase)

	// oauth := router.Group("/oauth")
	// {
	// 	oauth.GET("/login", userController.HandleGoogleLogin)
	// 	oauth.GET("/callback/", userController.HandleGoogleCallback)
	// }
    // // r.GET("/login/google", userController.HandleGoogleLogin)
    // // r.GET("/callback", userController.HandleGoogleCallback)
}