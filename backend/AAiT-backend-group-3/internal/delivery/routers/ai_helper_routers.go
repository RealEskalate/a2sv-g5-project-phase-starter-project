// create a new router for the ai_helper service

package routers

import (
	"AAIT-backend-group-3/internal/delivery/controllers"
	"github.com/gin-gonic/gin"
)

func CreateAiHelperRouter(router *gin.Engine, AiHelperController controllers.AiHelperControllerInterface) {
	router.POST("/ai_helper/generate_blog", AiHelperController.GenerateBlog)
	router.POST("/ai_helper/enhance_blog", AiHelperController.EnhanceBlog)
	router.POST("/ai_helper/generate_summary", AiHelperController.GenerateSummary)
	router.POST("/ai_helper/generate_tags", AiHelperController.GenerateTags)
}
