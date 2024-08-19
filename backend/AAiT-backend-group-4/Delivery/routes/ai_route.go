package routes

import (
	bootstrap "aait-backend-group4/Bootstrap"
	controllers "aait-backend-group4/Delivery/Controllers"
	infrastructure "aait-backend-group4/Infrastructure"
	usecases "aait-backend-group4/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAiRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ar := infrastructure.NewAiService(env)
	ac := controllers.AIController{
		AiUsecase: usecases.NewAiUsecase(
			timeout, ar,
		),
		Env: env,
	}

	group.POST("/blog/generateWithPrompt", ac.GenerateTextWithPrompt)
	group.POST("/blog/generateWithTags", ac.GenerateTextWithTags)
	group.POST("/blog/generateSuggestions", ac.GenerateSuggestions)
}
