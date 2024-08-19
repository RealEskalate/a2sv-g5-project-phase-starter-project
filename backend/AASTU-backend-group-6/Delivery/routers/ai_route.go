package routers

import (
	controllers "blogs/Delivery/controllers"
	infrastructure "blogs/Infrastructure"
	usecases "blogs/Usecases" // Renamed "usecases" to "usecase"
	"blogs/mongo"
	"time"

	"github.com/gin-gonic/gin"
)

func NewAIRoute(config *infrastructure.Config, DB mongo.Database, airoute *gin.RouterGroup) {

	aiServcie := infrastructure.NewAIService(config)
	aiUsecase := usecases.NewAIUsecase(aiServcie , time.Duration(config.ContextTimeout) * time.Second)


	AiController := controllers.AiController{
		Config: config,
		AiUsecase: aiUsecase,
	}

	airoute.POST("/ask", AiController.Ask)

}