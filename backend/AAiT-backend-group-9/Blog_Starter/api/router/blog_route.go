package router

import (
	"Blog_Starter/api/controller"
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"Blog_Starter/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBlogRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {

	br := repository.NewBlogRepository(db, domain.CollectionBlog)
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	bu := usecase.NewBlogUseCase(br, ur, 100 * time.Second)
	bc := controller.NewBlogController(bu)

	group.POST("/create", bc.CreateBlog)
	group.GET("/", bc.GetAllBlog)
	group.GET("/:id", bc.GetBlogByID)
	group.PUT("/:id", bc.UpdateBlog)
	group.DELETE("/:id", bc.DeleteBlog)
	group.GET("/search", bc.SearchBlog)
	group.GET("/filter", bc.FilterBlog)
}