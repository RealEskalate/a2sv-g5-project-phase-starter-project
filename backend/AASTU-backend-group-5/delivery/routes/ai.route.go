package routes

import (
	"log"
	"os"
	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/infrastructure"
	"github.com/RealEskalate/blogpost/infrastructure/middleware"
	tokenservice "github.com/RealEskalate/blogpost/infrastructure/token_service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func NewAiRequestRoute(group *gin.RouterGroup) {
	var Ai infrastructure.AI
	ctrl := controller.AI_controller{
		Ai_func: Ai,
	}

	//load middlewares
	err := godotenv.Load()
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
	TokenSvc := *tokenservice.NewTokenService(access_secret, refresh_secret)

	LoggedInmiddleWare := middleware.LoggedIn(TokenSvc)


	group.POST("api/generate-blog/", LoggedInmiddleWare ,ctrl.GenerateBlog())
}