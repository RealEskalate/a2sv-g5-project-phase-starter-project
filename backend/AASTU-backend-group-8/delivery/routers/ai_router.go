package routers

import (
	"meleket/delivery/controllers"
	// "meleket/domain"
	"meleket/infrastructure"

	"github.com/gin-gonic/gin"
)

func AIRouter(r *gin.Engine, ai *controllers.AIHandler, jwtService infrastructure.JWTService) {
	api := r.Group("/api")
	api.Use(infrastructure.AuthMiddleware(jwtService))
	{
		api.POST("/blogwithai", ai.GenerateBlogWithAI)
	}
}
