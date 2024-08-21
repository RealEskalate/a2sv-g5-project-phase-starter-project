package routers

import (
	"time"

	config "github.com/aait.backend.g5.main/backend/Config"
	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	repository "github.com/aait.backend.g5.main/backend/Repository"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	utils "github.com/aait.backend.g5.main/backend/Utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBlogRouter(env *config.Env, database mongo.Database, group *gin.RouterGroup, redisClient *redis.Client) {

	popularity_repository := repository.NewBlogPopularityActionRepository(&database)
	user_repository := repository.NewUserRepository(&database)
	blog_repository := repository.NewBlogRepository(&database)
	cacheService := infrastructure.NewRedisCache(redisClient)
	helper := utils.NewBlogHelper()
	blogUsecase := usecases.NewblogUsecase(blog_repository, cacheService, *env, time.Hour*24, helper, user_repository, popularity_repository)

	blogController := controllers.NewBlogController(blogUsecase)

	group.POST("/createBlog", blogController.CreateBlogController)
	group.GET("/getBlog/:id", blogController.GetBlogController)
	group.GET("/getBlogs", blogController.GetBlogsController)
	group.PUT("/updateBlog/:id", blogController.UpdateBlogController)
	group.DELETE("/deleteBlog/:id", blogController.DeleteBlogController)
	group.POST("/action", blogController.TrackPopularityController)
	group.POST("/filter", blogController.SearchBlogsController)
	group.POST("/comment", blogController.AddCommentController)
}
