package routers

import (
	controllers "blogs/Delivery/controllers"
	infrastructure "blogs/Infrastructure"
	repositories "blogs/Repositories"
	usecases "blogs/Usecases"
	"blogs/mongo"
	"time"

	"github.com/gin-gonic/gin"
)

func NewOauthRoute(config *infrastructure.Config, DB mongo.Database, OauthRoute *gin.RouterGroup) {

	oauth:= infrastructure.NewOauthConfig(config)

	repo := repositories.NewSignupRepository(DB , config.UserCollection)
	
	userrepo := repositories.NewUserRepository(DB, config.UserCollection)
	aur := repositories.NewActiveUserRepository(DB, config.ActiveUserCollection)
	usecase := usecases.NewOauthUsecase(repo , time.Duration(config.ContextTimeout) * time.Second , oauth )
	loginusecase := usecases.NewLoginUsecase(userrepo, aur , time.Duration(config.ContextTimeout) * time.Second)

	oauthcontroller := controllers.OauthController{
		OauthUsecase : usecase,
		Login : loginusecase,
		Config : config,
	}

	OauthRoute.GET("/auth/google" , oauthcontroller.GoogleAuth)
	OauthRoute.GET("/auth/callback" , oauthcontroller.GoogleCallback)
	
}