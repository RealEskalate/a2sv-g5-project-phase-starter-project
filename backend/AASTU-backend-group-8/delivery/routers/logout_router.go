package routers

import (
	"meleket/delivery/controllers"
	"meleket/domain"
	"meleket/infrastructure"

	"github.com/gin-gonic/gin"
)

func NewLogoutRouter(r *gin.Engine, userUsecase domain.UserUsecaseInterface, refreshTokenUsecase domain.RefreshTokenUsecaseInterface, jwtService infrastructure.JWTService) {
	logoutController := controllers.NewLogoutController(refreshTokenUsecase)
	// Authenticated routes
	auth := r.Group("/api")
	auth.Use(infrastructure.AuthMiddleware(jwtService))
	{
		auth.POST("/logout", logoutController.Logout)
	}

}
