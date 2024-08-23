package route

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/generative-ai-go/genai"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBlogRouter(db *mongo.Database, group *gin.RouterGroup, model *genai.GenerativeModel, redis *redis.Client)  {
	br := repository.NewBlogRepository(db.Collection("blogs"), context.TODO())
	bs := service.NewBlogService(br, redis, time.Minute * 5)
	ais := service.NewAIContentService(context.TODO(), model, br)

	pts := service.NewPopularityTrackingService(br,redis, time.Minute * 5)
	

	aic := controller.NewAIContentController(ais)
	bc := controller.NewBlogController(bs)
	ptc := controller.NewPopularityTrackingController(pts)


	group.POST("", bc.CreateBlogPost)
	group.GET("", bc.GetBlogPosts)
	group.GET(":id", bc.GetBlogPost)
	group.PUT(":id", bc.UpdateBlogPost)
	group.DELETE(":id", bc.DeleteBlogPost)
	group.POST("search", bc.SearchBlogPosts)
	group.POST("filter", bc.FilterBlogPosts)

	group.POST("generate", aic.GenerateContentSuggestions)
	group.POST("enhance/:id", aic.SuggestContentImprovements)

	group.POST("like/:id", ptc.LikeBlogPost)
	group.POST("dislike/:id", ptc.DislikeBlogPost)
}
