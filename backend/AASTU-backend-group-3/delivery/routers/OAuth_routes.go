package routers

import (
	"group3-blogApi/config"
	"group3-blogApi/delivery/controllers"
	"group3-blogApi/repository"
	"group3-blogApi/usecase"

	"github.com/gin-gonic/gin"
)


func SetupOauthRouter(router *gin.Engine) {
	oauthConfig := config.NewOAuthConfig()
	oauthRepo := repository.NewOAuthRepository(oauthConfig)
	oauthUsecase := usecase.NewOAuthUsecase(oauthRepo)
	oauthController := controllers.NewOAuthController(oauthUsecase)

	oauth := router.Group("/oauth")
	{
		oauth.GET("/login", oauthController.HandleGoogleLogin)
		oauth.GET("/callback/", oauthController.HandleGoogleCallback)
	}
}