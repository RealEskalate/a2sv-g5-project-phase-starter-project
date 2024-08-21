package routers

import (
	"context"
	"log"

	config "github.com/aait.backend.g5.main/backend/Config"
	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	repository "github.com/aait.backend.g5.main/backend/Repository"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func NewAISuggestionRouter(env *config.Env, group *gin.RouterGroup) {
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(env.GEMINI_API_KEY))
	blog_repo := repository.NewBlogRepository()

	if err != nil {
		log.Fatal(err)
	}

	// instantiate Logout controller
	AISuggestionController := &controllers.ContentSuggestionController{
		AISuggestionUsecase: usecases.NewAISuggestionUsecase(client, blog_repo),
	}

	group.GET("/suggestContent", AISuggestionController.HandleSuggestion)
}
