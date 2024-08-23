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

func NewUserRoute(group *gin.RouterGroup  , user_collection database.CollectionInterface) {
	repo := repository.NewUserRepository(user_collection)
	usecase := usecase.NewUserUseCase(repo)
	ctrl := controller.NewUserController(usecase)
	

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
	mustOwn := middleware.RoleBasedAuth(false)
	mustBeAdmin := middleware.RoleBasedAuth(true)


	group.GET("api/user/:id", LoggedInmiddleWare, ctrl.GetOneUser())
	group.GET("api/user/",LoggedInmiddleWare ,ctrl.GetUsers())

	group.PUT("api/user/:id", LoggedInmiddleWare , mustOwn , ctrl.UpdateUser())
	group.DELETE("api/user/:id",LoggedInmiddleWare ,mustOwn, ctrl.DeleteUser())

	group.PUT("api/demote/:id" , LoggedInmiddleWare , mustBeAdmin , ctrl.DemoteUser())
	group.PUT("api/promote/:id" , LoggedInmiddleWare , mustBeAdmin , ctrl.PromoteUser())
}