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

func NewBlogRouter(db *mongo.Database, group *gin.RouterGroup, model *genai.GenerativeModel)  {
	br := repository.NewBlogRepository(db.Collection("blogs"), context.TODO())
	bs := service.NewBlogService(br)
	ais := service.NewAIContentService(context.TODO(), model, br)
	

	aic := controller.NewAIContentController(ais)

	ac := controller.NewBlogController(bs)


	group.POST("", ac.CreateBlogPost)
	group.GET("", ac.GetBlogPosts)
	group.GET(":id", ac.GetBlogPost)
	group.PUT(":id", ac.UpdateBlogPost)
	group.DELETE(":id", ac.DeleteBlogPost)

	group.POST("generate", aic.GenerateContentSuggestions)
	group.POST("enhance/:id", aic.SuggestContentImprovements)
}
