package route

import (
	"blog/config"
	"blog/database"
	"blog/delivery/controller"
	"blog/domain"
	"blog/repository"
	"blog/usecase"
	// "go/token"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *config.Env, timeout time.Duration, db database.Database, router *gin.RouterGroup) {
	loginRepo := repository.NewUserRepository(db, domain.CollectionUser)
	tokenRepo := repository.NewMongoTokenRepository(db, domain.TokenCollection)
	// loginUsecase := usecase.NewLoginUsecase(loginRepo, timeout)
	// tokenUsecase := usecase.NewTokenUsecase(tokenRepo, timeout)

	loginController := &controller.LoginController{
		
		LoginUsecase: usecase.NewLoginUsecase(loginRepo,tokenRepo, timeout),
		
		Env:          env,
	}

	router.POST("/login", loginController.Login)
	router.POST("/refresh", loginController.RefreshTokenHandler)
}
