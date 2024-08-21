package routes

import (
	"log"
	"os"

	"github.com/RealEskalate/blogpost/config"
	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/delivery/controller"
	oauthservice "github.com/RealEskalate/blogpost/infrastructure/oauth_service"
	passwordservice "github.com/RealEskalate/blogpost/infrastructure/password_service"
	tokenservice "github.com/RealEskalate/blogpost/infrastructure/token_service"
	"github.com/RealEskalate/blogpost/repository"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func NewAuthRoute(group *gin.RouterGroup, users, state database.CollectionInterface){
	AuthRepo, err := repository.NewAuthRepo(users)
	if err != nil{
		log.Panic(err.Error())
	}

	StateRepo := repository.NewStateRepo(state)
	err = godotenv.Load()
	if err != nil {
        log.Panic(err.Error())
    }

	access_secret := os.Getenv("ACCESSTOKENSECRET")
	if access_secret == ""{
		log.Panic("No accesstoken")
	}
	
	refresh_secret := os.Getenv("REFRESHTOKENSECRET")
	if refresh_secret == ""{
		log.Panic("No refreshtoken")
	}
	
	TokenSvc := tokenservice.NewTokenService(access_secret, refresh_secret)
	PasswordSvc := &passwordservice.PasswordS{}

	GOC, err := config.InitOauth()
	if err != nil{
		log.Panic("no GOC")
	}
	OauthSvc := oauthservice.NewGoogleOAuthService(GOC)

	AuthUsecase := usecase.NewAuthUsecase(AuthRepo, StateRepo, PasswordSvc, TokenSvc, OauthSvc)
	AuthController := controller.NewAuthController(AuthUsecase)

	group.POST("/signup", AuthController.SignUp())
	group.POST("/login", AuthController.LogIn())
	group.GET("/google", AuthController.GoogleLogIn())
	group.GET("/oauth2/callback/google", AuthController.GoogleCallBack())
	group.POST("/logout", AuthController.LogOut())
	group.POST("/refresh", AuthController.Refresh())

}