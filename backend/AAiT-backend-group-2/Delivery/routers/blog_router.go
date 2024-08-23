package routers

import (
	"AAiT-backend-group-2/Delivery/controllers"
	domain "AAiT-backend-group-2/Domain"
	infrastructure "AAiT-backend-group-2/Infrastructure"
	"AAiT-backend-group-2/Infrastructure/cache"
	"AAiT-backend-group-2/Repositories/blog_repository"
	commentrepository "AAiT-backend-group-2/Repositories/comment_repository"
	"AAiT-backend-group-2/Usecases/blog_usecase"
	commentusecase "AAiT-backend-group-2/Usecases/comment_usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBlogRouter(db *mongo.Database, group *gin.RouterGroup, configs *domain.Config) {
	redisCache := cache.NewRedisCache(configs.RedisAdr, configs.RedisPass, 0)
	blogRepo := blog_repository.NewBlogRepository(db, redisCache)
	commentRepo := commentrepository.NewCommentRepository(db)
	blogUsecase := blog_usecase.NewBlogUsecase(blogRepo)
	commentUsecase := commentusecase.NewCommentUsecase(commentRepo)
	blogController := controllers.NewBlogController(blogUsecase, commentUsecase)

	blogRoutes := group.Group("/blogs")
	blogRoutes.GET("", blogController.GetAllBlogs)
	blogRoutes.GET("/:id", blogController.GetBlogByID)
	blogRoutes.GET("/blogFilter", blogController.FilterBlogs)
	blogRoutes.GET("/:id/comments", blogController.GetCommentsByBlogID) // Get all comments for a blog post

	blogRoutes.Use(infrastructure.AuthMiddleWare(configs.SecretKey))
	{
		blogRoutes.POST("", blogController.CreateBlog)
		// blogRoutes.GET("/:id", blogController.GetBlogByID)
		blogRoutes.PUT("/:id", blogController.UpdateBlog)

		// comment route
		blogRoutes.POST("/:id/comments", blogController.CreateComment)        // Add a comment to a blog post
		blogRoutes.PUT("/comments/:comment_id", blogController.UpdateComment) // Update a comment
		blogRoutes.DELETE("/comments/:comment_id", blogController.DeleteComment)
	}
	blogRoutes.Use(infrastructure.AuthMiddleWare(configs.SecretKey), infrastructure.RoleMiddleware())
	{
		blogRoutes.DELETE("/:id", blogController.DeleteBlog)
	}

}
