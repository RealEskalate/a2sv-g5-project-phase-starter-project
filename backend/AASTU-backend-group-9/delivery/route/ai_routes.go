package route

import (
    "blog/config"
    "blog/delivery/controller"
    "blog/delivery/middleware"
    "blog/usecase"
    "time"

    "github.com/gin-gonic/gin"
)

func RegisterAIRoutes(env *config.Env, timeout time.Duration, router *gin.RouterGroup) {
    aiUse := usecase.NewAIUsecase(env.AIAPIKey, timeout)
    aiController := &controller.AIController{
        AIUsecase: aiUse,
    }

    aiRoutes := router.Group("")
    aiRoutes.Use(middleware.AuthMidd) 
    {
        aiRoutes.GET("/generate-content", aiController.GenerateContent)
    }
}