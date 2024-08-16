package route

import (
	"blog/config"
	"blog/database"
	"blog/delivery/controller"
	"blog/repository"
	"blog/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *config.Env, timeout time.Duration, db database.Database, router *gin.RouterGroup) {
	loginRepo := repository.NewUserRepository(db, "users")
	loginUsecase := usecase.NewLoginUsecase(loginRepo, timeout)
	loginController := &controller.LoginController{
		LoginUsecase: loginUsecase,
		Env:          env,
	}

	router.POST("/login", loginController.Login)
}
