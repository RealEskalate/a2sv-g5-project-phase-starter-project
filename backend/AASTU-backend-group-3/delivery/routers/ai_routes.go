package routers

import (
	"github.com/gin-gonic/gin"

	"group3-blogApi/config"
	"group3-blogApi/delivery/controllers"
	"group3-blogApi/usecase"
)

func SetUpAi(router *gin.Engine) {

	geminiApiKey := config.EnvConfigs.GEMINI_API_KEY

	// Initialize the use case
	useCase := usecase.NewAIUseCase(geminiApiKey)

	// Initialize the controller
	aiController := controllers.NewAIController(useCase)

	// Define the single route
	router.POST("/generate", func(c *gin.Context) {
		aiController.GenerateContent(c)
	})
}
