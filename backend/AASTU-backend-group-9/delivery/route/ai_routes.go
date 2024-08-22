package route

import (
	"blog/config"
	"blog/delivery/controller"
	"blog/delivery/middleware"
	"blog/usecase"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
)

func RegisterAIRoutes(env *config.Env, timeout time.Duration, router *gin.RouterGroup, client *genai.Client) {
	aiUse := usecase.NewAIUsecase(timeout,client)
	aiController := &controller.AIController{
		AIUsecase: aiUse,
	}

	aiRoutes := router.Group("")
	aiRoutes.Use(middleware.AuthMidd)
	{
		aiRoutes.POST("/generate-content", aiController.GenerateContent)
	}
}
