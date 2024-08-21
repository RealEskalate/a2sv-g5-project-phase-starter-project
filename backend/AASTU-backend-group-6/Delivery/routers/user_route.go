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

func NewUserrouter(config *infrastructure.Config, DB mongo.Database, userRouter *gin.RouterGroup) {

	userRepo := repositories.NewUserRepository(DB, config.UserCollection)
	passwordService := &infrastructure.DefaultPasswordService{}

	userService := usecases.NewUserUseCase(userRepo, time.Duration(config.ContextTimeout)*time.Second , passwordService)
	userController := controllers.NewUserController{
		UserUsecase: userService,
	}

	userRouter.PUT("/update/:id", userController.UpdateUser)
	userRouter.POST("/promote/:id", userController.PromoteUser)

}
