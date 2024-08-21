package route

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAiRouter(db *mongo.Database, group *gin.RouterGroup, model *genai.GenerativeModel)  {
	br := repository.NewBlogRepository(db.Collection("blogs"), context.TODO())
	ais := service.NewAIContentService(context.TODO(), model, br)
	
	aic := controller.NewAIContentController(ais)
	group.POST("generate-blog", aic.GenerateContentSuggestions)
	group.POST("enhance/:id", aic.SuggestContentImprovements)

}
