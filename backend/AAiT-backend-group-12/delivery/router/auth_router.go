package router

import (
	"blog_api/delivery/controllers"
	"blog_api/repository"
	"blog_api/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthRouter(collection *mongo.Collection, authGroup *gin.RouterGroup) {
	repository := repository.NewUserRepository(collection)
	usecase := usecase.NewUserUsecase(repository)
	controller := controllers.NewAuthController(usecase)

	authGroup.POST("/signup", controller.HandleSignup)
	authGroup.POST("/login", controller.HandleLogin)
	authGroup.POST("/renew-token", controller.HandleRenewAccessToken)
}
