package router

import (
	"blog_api/delivery/controllers"
	"blog_api/delivery/env"
	ai_service "blog_api/infrastructure/ai"
	"blog_api/repository"
	"blog_api/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewBlogRouter initalizes the controllers, usecases and repositories before setting up the blog routes
func NewBlogRouter(collection *mongo.Collection, blogGroup *gin.RouterGroup) {
	br := repository.NewBlogRepository(collection)
	bu := usecase.NewBlogUseCase(br, time.Second*100)
	bu := usecase.NewBlogUseCase(br, time.Second*100, ai_service.NewAIService(env.ENV.GEMINI_API_KEY))

	bc := controllers.NewBlogController(bu)

	blogGroup.POST("/create", bc.CreateBlogHandler)
	blogGroup.PUT("/:id", bc.UpdateBlogHandler)
	blogGroup.DELETE("/:id", bc.DeleteBlogHandler)
	blogGroup.POST("/", bc.GetBlogHandler)
	blogGroup.GET("/:id", bc.GetBlogByIDHandler)
	blogGroup.POST("/update-popularity", bc.TrackBlogPopularityHandler)
}
