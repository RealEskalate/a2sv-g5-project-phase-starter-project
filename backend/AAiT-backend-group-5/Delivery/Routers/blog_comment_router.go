package routers

import (
	config "github.com/aait.backend.g5.main/backend/Config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"	
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"


	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	repository "github.com/aait.backend.g5.main/backend/Repository"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
)

func NewBlogCommentRouter(env *config.Env, database interfaces.Database, group *gin.RouterGroup, redisClient *redis.Client) {
	blog_repository := repository.NewBlogRepository(database)
	blog_comment_repository := repository.NewBlogCommentRepository(database)
	cacheService := infrastructure.NewRedisCache(redisClient)
	blogCommentUsecase := usecases.NewCommentUsecase(blog_comment_repository, blog_repository, cacheService)
	blogCommentController := controllers.NewBlogCommentController(blogCommentUsecase)

	group.POST("/comment/:blogID", blogCommentController.AddCommentController)
	group.GET("/comments/:blogID", blogCommentController.GetCommentsController)
	group.GET("/comment/:commentID", blogCommentController.GetCommentController)
	group.PUT("/comment/:commentID", blogCommentController.UpdateCommentController)
	group.DELETE("/comment/:commentID", blogCommentController.DeleteCommentController)
}
