package route

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"backend-starter-project/mongo"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
)

func NewAiRouter(db *mongo.Database, group *gin.RouterGroup, model *genai.GenerativeModel)  {
	collection := (*db).Collection("blogs")
	br := repository.NewBlogRepository(&collection, context.TODO())
	ais := service.NewAIContentService(context.TODO(), model, br)
	
	aic := controller.NewAIContentController(ais)
	group.POST("generate-blog", aic.GenerateContentSuggestions)
	group.POST("enhance/:id", aic.SuggestContentImprovements)

}
