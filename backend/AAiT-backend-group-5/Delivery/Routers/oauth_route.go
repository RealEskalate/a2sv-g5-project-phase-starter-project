package routers

import (
	config "github.com/aait.backend.g5.main/backend/Config"
	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	repository "github.com/aait.backend.g5.main/backend/Repository"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/gin-gonic/gin"
)

func NewOAuthRouter(env *config.Env, db interfaces.Database, router gin.RouterGroup) {
	config := config.NewOAuthConfig(*env)
	sessionRepo := repository.NewSessionRepository(db)
	useRepo := repository.NewUserRepository(db)
	oauth_usecase := usecases.NewOAuthUseCase(useRepo, sessionRepo)
	OAuthController := controllers.NewOAuthController(*config, *env, oauth_usecase)

	router.GET("/oauth/login", OAuthController.LoginHandlerController)
	router.GET("/auth/oauth", OAuthController.OAuthHanderController)
	router.GET("/auth/callback", OAuthController.OAuthCallbackHandler)

}
