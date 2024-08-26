package routers

import (
	"meleket/delivery/controllers"
	"meleket/domain"
	"meleket/infrastructure"

	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(r *gin.Engine, userUsecase domain.UserUsecaseInterface, refreshTokenUsecase domain.RefreshTokenUsecaseInterface, jwtService infrastructure.JWTService) {
	refreshTokenController := controllers.NewRefreshTokenController(userUsecase, refreshTokenUsecase, jwtService)
	r.POST("/refreshtoken", refreshTokenController.RefreshToken)
}
