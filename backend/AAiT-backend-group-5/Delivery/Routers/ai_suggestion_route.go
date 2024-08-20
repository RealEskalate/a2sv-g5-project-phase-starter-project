package routers

import (
	config "github.com/aait.backend.g5.main/backend/Config"
	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	"github.com/gin-gonic/gin"
)

func NewAISuggestionRouter(env *config.Env, group *gin.RouterGroup) {

	// instantiate Logout controller
	AISuggestionController := &controllers.ContentSuggestionController{
		AI_Service: infrastructure.NewAIContentSuggester(env.GEMINI_API_KEY),
	}

	group.GET("/suggestContent", AISuggestionController.HandleSuggestion)
}
