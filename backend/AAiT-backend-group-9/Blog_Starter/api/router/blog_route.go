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
	lr := repository.NewLikeRepository(database, domain.CollectionLike)
	cr := repository.NewCommentRepository(database, domain.CollectionComment)
	rtr := repository.NewBlogRatingRepository(database, domain.CollectionRating)

	bu := usecase.NewBlogUseCase(br, ur, lr, cr, rtr, timeout, cacheServic)
	bc := controller.NewBlogController(bu)

	group.POST("/", bc.CreateBlog)
	group.GET("/", bc.GetAllBlog)
	group.GET("/:blog_id", bc.GetBlogByID)
	group.PATCH("/:blog_id", bc.UpdateBlog)
	group.DELETE("/:blog_id", bc.DeleteBlog)
	group.GET("/search", bc.SearchBlog)
	group.GET("/filter", bc.FilterBlog)
}
