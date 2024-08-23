package routers

import (
	"time"

	config "github.com/aait.backend.g5.main/backend/Config"
	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	repository "github.com/aait.backend.g5.main/backend/Repository"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	utils "github.com/aait.backend.g5.main/backend/Utils"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewBlogRouter(env *config.Env, database interfaces.Database, group *gin.RouterGroup, redisClient *redis.Client) {

	popularity_repository := repository.NewBlogPopularityActionRepository(database)
	user_repository := repository.NewUserRepository(database)
	blog_repository := repository.NewBlogRepository(database)
	blog_comment_repository := repository.NewBlogCommentRepository(database)
	cacheService := infrastructure.NewRedisCache(redisClient)
	helper := utils.NewBlogHelper(blog_repository, cacheService, time.Duration(time.Hour*24), blog_comment_repository)
	blogUsecase := usecases.NewblogUsecase(blog_repository, cacheService, *env, time.Hour*24, helper, user_repository, popularity_repository, blog_comment_repository)

	blogController := controllers.NewBlogController(blogUsecase)

	group.POST("/blog", blogController.CreateBlogController)
	group.GET("/blog/:id", blogController.GetBlogController)
	group.GET("/blogs", blogController.GetBlogsController)
	group.PUT("/blog/:id", blogController.UpdateBlogController)
	group.DELETE("/blog/:id", blogController.DeleteBlogController)

	group.POST("/blog/action/:id", blogController.TrackPopularityController)
	group.GET("/blog/filter", blogController.SearchBlogsController)
}
