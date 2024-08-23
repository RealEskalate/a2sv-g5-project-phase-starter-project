package routers

import (
	config "github.com/aait.backend.g5.main/backend/Config"
	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	repository "github.com/aait.backend.g5.main/backend/Repository"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/gin-gonic/gin"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
)

func NewLogoutRouter(env *config.Env, database interfaces.Database, group *gin.RouterGroup) {
	session_repo := repository.NewSessionRepository(database)
	jwt_service := infrastructure.NewJwtService(env)

	// instantiate Logout controller
	LogoutController := &controllers.LogoutController{
		LogoutUsecase: usecases.NewLogoutUsecase(jwt_service, session_repo),
		JwtService:    infrastructure.NewJwtService(env),
	}

	group.POST("/logout", LogoutController.Logout)
}
