package router

import (
	"blog_api/repository"
	"blog_api/delivery/controllers"
	"blog_api/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBlogRouter(collection *mongo.Collection, taskGroup *gin.RouterGroup) {
	br := repository.NewBlogRepository(collection)
	bu := usecase.NewBlogUseCase(br, time.Second * 100)

	bc := controllers.NewBlogController(bu)

	taskGroup.POST("/create", bc.CreateBlogHandler)
	taskGroup.PUT("/:id", bc.UpdateBlogHandler)
	taskGroup.DELETE("/:id", bc.DeleteBlogHandler)
	taskGroup.POST("/", bc.GetBlogHandler)
	taskGroup.GET("/:id", bc.GetBlogByIDHandler)
}
