package router

import (
	"Blog_Starter/api/controller"
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"Blog_Starter/usecase"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBlogRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() 
	br := repository.NewBlogRepository(db, domain.CollectionBlog, &ctx)
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	bu := usecase.NewBlogUseCase(br, ur)
	bc := controller.NewBlogController(bu, ctx)

	group.POST("/create", bc.CreateBlog)
	group.GET("/", bc.GetAllBlog)
	group.GET("/:id", bc.GetBlogByID)
	group.PUT("/:id", bc.UpdateBlog)
	group.DELETE("/:id", bc.DeleteBlog)
	group.GET("/search", bc.SearchBlog)
	group.GET("/filter", bc.FilterBlog)
}