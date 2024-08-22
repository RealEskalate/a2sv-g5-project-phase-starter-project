package routes

import (
	"log"
	"os"
	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/infrastructure/middleware"
	tokenservice "github.com/RealEskalate/blogpost/infrastructure/token_service"
	"github.com/RealEskalate/blogpost/repository"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func NewPopularityRoutes(group *gin.RouterGroup, blog_collection database.CollectionInterface) {
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

	mongoCollection := blog_collection.(*database.MongoCollection).Collection

	popularityRepo := repository.NewBlogPopularityRepository(mongoCollection)

	popularityUsecase := usecase.NewBlogPopularityUsecase(popularityRepo)

	popularityCtrl := controller.NewPopularityController(popularityUsecase)

	group.GET("/popular-blogs", LoggedInmiddleWare,popularityCtrl.GetPopularBlogs())
}
