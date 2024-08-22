package router

import (
	"Blog_Starter/api/controller"
	"Blog_Starter/config"
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"Blog_Starter/usecase"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBlogRouter(env *config.Env, timeout time.Duration, db *mongo.Client, group *gin.RouterGroup) {

	database := db.Database(env.DBName) // Replace with your actual database name
	br := repository.NewBlogRepository(database, domain.CollectionBlog)
	ur := repository.NewUserRepository(database, domain.CollectionUser)
	bu := usecase.NewBlogUseCase(br, ur, timeout)
	bc := controller.NewBlogController(bu)

	group.POST("/create", bc.CreateBlog)
	group.GET("/", bc.GetAllBlog)
	group.GET("/:id", bc.GetBlogByID)
	group.PUT("/:id", bc.UpdateBlog)
	group.DELETE("/:id", bc.DeleteBlog)
	group.GET("/search", bc.SearchBlog)
	group.GET("/filter", bc.FilterBlog)
}