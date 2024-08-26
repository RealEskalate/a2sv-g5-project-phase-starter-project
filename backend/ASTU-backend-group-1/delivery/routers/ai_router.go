package routers

import (
	"astu-backend-g1/config"
	"astu-backend-g1/delivery/controllers"
	"astu-backend-g1/gemini"
	"astu-backend-g1/infrastructure"
	"log"

	"github.com/gin-gonic/gin"
)

func  (gr *MainRouter) AddAIRoutes(r *gin.Engine, config config.Config, prompts infrastructure.Prompts)  *gin.RouterGroup{
	model, err := gemini.NewGeminiModel(config.Gemini.ApiKey, config.Gemini.Model, prompts)
	if err != nil {
		log.Fatal(err)
	}
	aiController := controllers.NewAIController(model)
	aiRouteGroup := r.Group("/ai")
	{
		aiRouteGroup.POST("/recommendTitle", aiController.RecommendTitle)
		aiRouteGroup.POST("/recommendContent", aiController.RecommendContent)
		aiRouteGroup.POST("/recommendTags", aiController.Recommendtags)
		aiRouteGroup.POST("/summarize", aiController.Sumarize)
		aiRouteGroup.POST("/chat", aiController.Chat)
	}
	return aiRouteGroup
}
