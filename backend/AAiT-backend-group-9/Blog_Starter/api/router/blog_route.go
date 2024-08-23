package router

import (
	"Blog_Starter/api/controller"
	"Blog_Starter/config"
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"Blog_Starter/usecase"
	"Blog_Starter/utils/infrastructure"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBlogRouter(env *config.Env, timeout time.Duration, db *mongo.Client, group *gin.RouterGroup) {

	database := db.Database(env.DBName) // Replace with your actual database name
	cacheServic := infrastructure.NewcacheServic(env.CacheAddr, "", 0)

	br := repository.NewBlogRepository(database, domain.CollectionBlog)
	ur := repository.NewUserRepository(database, domain.CollectionUser)

	bu := usecase.NewBlogUseCase(br, ur, timeout, cacheServic)
	bc := controller.NewBlogController(bu)

	group.POST("/", bc.CreateBlog)
	group.GET("/", bc.GetAllBlog)
	group.GET("/:id", bc.GetBlogByID)
	group.PATCH("/:id", bc.UpdateBlog)
	group.DELETE("/:id", bc.DeleteBlog)
	group.GET("/search", bc.SearchBlog)
	group.GET("/filter", bc.FilterBlog)
}
