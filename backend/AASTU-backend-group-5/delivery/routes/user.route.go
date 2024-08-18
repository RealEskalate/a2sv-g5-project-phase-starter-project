package routes

import (
	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/repository"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
)

func NewUserRoute(group *gin.RouterGroup  , user_collection database.CollectionInterface) {
	repo := repository.NewUserRepository(user_collection)
	usecase := usecase.NewUserUseCase(repo)
	ctrl := controller.NewUserController(usecase)
	
	group.GET("api/user/:id", ctrl.GetOneUser())
	group.GET("api/user/", ctrl.GetUsers())
	group.POST("api/user/", ctrl.Register())
	group.PUT("api/user/:id", ctrl.UpdateUser())
	group.DELETE("api/user/:id", ctrl.DeleteUser())
}