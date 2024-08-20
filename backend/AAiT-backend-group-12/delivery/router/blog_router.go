package router

import (
	"blog_api/delivery/controllers"
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

	bc := controllers.NewBlogController(bu)

	blogGroup.POST("/create", bc.CreateBlogHandler)
	blogGroup.PUT("/:id", bc.UpdateBlogHandler)
	blogGroup.DELETE("/:id", bc.DeleteBlogHandler)
	blogGroup.POST("/", bc.GetBlogHandler)
	blogGroup.GET("/:id", bc.GetBlogByIDHandler)
	blogGroup.POST("/update-popularity", bc.TrackBlogPopularityHandler)
}
