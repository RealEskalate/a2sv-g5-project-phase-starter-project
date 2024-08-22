package routers

import (
	"meleket/delivery/controllers"
	"meleket/infrastructure"

	"github.com/gin-gonic/gin"
)

func PromoteDemoteRouter(r *gin.Engine, promoteDemoteController *controllers.PromoteDemoteController, jwtService infrastructure.JWTService) {
	api := r.Group("/api")
	api.Use(infrastructure.AuthMiddleware(jwtService))
	api.Use(infrastructure.AdminMiddleware(jwtService))
	{
		api.POST("/promote", promoteDemoteController.PromoteToAdmin)
		api.POST("/demote", promoteDemoteController.DemoteToUser)
	}
}
