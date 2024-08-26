package routers

import (
	"meleket/delivery/controllers"
	"meleket/domain"
	// "meleket/domain"
	"meleket/infrastructure"

	"github.com/gin-gonic/gin"
)

func AIRouter(r *gin.Engine, aiUsecase domain.AIUsecaseInterface, jwtService infrastructure.JWTService) {
	ai := controllers.NewAIHandler(aiUsecase)
	api := r.Group("/api")
	api.Use(infrastructure.AuthMiddleware(jwtService))
	{
		api.POST("/blogwithai", ai.GenerateBlogWithAI)
	}
}
