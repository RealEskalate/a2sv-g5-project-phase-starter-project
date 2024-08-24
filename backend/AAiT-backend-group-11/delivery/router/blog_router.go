package route

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/mongo"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
)

func NewBlogRouter(db *mongo.Database, group *gin.RouterGroup, model *genai.GenerativeModel)  {
	blogcollection := (*db).Collection("blogs")
	br := repository.NewBlogRepository(&blogcollection, context.TODO())
	bs := service.NewBlogService(br)
	ais := service.NewAIContentService(context.TODO(), model, br)

	commentcollection := (*db).Collection("comments")
	cr := repository.NewCommentRepository(&commentcollection, context.TODO())

	pts := service.NewPopularityTrackingService(br,cr)
	

	aic := controller.NewAIContentController(ais)
	ac := controller.NewBlogController(bs)
	ptc := controller.NewPopularityTrackingController(pts)


	group.POST("", ac.CreateBlogPost)
	group.GET("", ac.GetBlogPosts)
	group.GET(":id", ac.GetBlogPost)
	group.PUT(":id", ac.UpdateBlogPost)
	group.DELETE(":id", ac.DeleteBlogPost)

	group.POST("generate", aic.GenerateContentSuggestions)
	group.POST("enhance/:id", aic.SuggestContentImprovements)

	group.POST("like/:id", ptc.LikeBlogPost)
	group.POST("dislike/:id", ptc.DislikeBlogPost)
}
