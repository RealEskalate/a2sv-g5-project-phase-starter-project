package routers

import (
	"astu-backend-g1/config"
	"astu-backend-g1/infrastructure"

	"github.com/gin-gonic/gin"
)

func (gr *MainRouter) AddAIRoutes(r *gin.Engine, config config.Config, prompts infrastructure.Prompts) *gin.RouterGroup {

	aiRouteGroup := r.Group("/ai")
	{
		aiRouteGroup.POST("/recommendTitles", gr.aiController.RecommendTitle)
		aiRouteGroup.POST("/recommendContent", gr.aiController.RecommendContent)
		aiRouteGroup.POST("/recommendTags", gr.aiController.Recommendtags)
		aiRouteGroup.POST("/summarize", gr.aiController.Sumarize)
		aiRouteGroup.POST("/chat", gr.aiController.Chat)
	}
	return aiRouteGroup
}
