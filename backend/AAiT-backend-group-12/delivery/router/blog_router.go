package router

import (
	"blog_api/delivery/controllers"
	"blog_api/delivery/env"
	"blog_api/domain"
	ai_service "blog_api/infrastructure/ai"
	jwt_service "blog_api/infrastructure/jwt"
	"blog_api/infrastructure/middleware"
	"blog_api/repository"
	"blog_api/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewBlogRouter initializes the controllers, usecases and repositories before setting up the blog and comment routes
func NewBlogRouter(collection *mongo.Collection, cacheClient *redis.Client, blogGroup *gin.RouterGroup) {
	br := repository.NewBlogRepository(collection)
	cacheRepoistory := repository.NewCacheRepository(cacheClient)
	jwtService := jwt_service.NewJWTService(env.ENV.JWT_SECRET_TOKEN)
	bu := usecase.NewBlogUseCase(br, time.Second*100, ai_service.NewAIService(env.ENV.GEMINI_API_KEY), cacheRepoistory, env.ENV)
	bc := controllers.NewBlogController(bu)

	blogGroup.POST("/create", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleRoot, domain.RoleUser, domain.RoleAdmin), bc.CreateBlogHandler)
	blogGroup.PUT("/:id", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleRoot, domain.RoleUser, domain.RoleAdmin), bc.UpdateBlogHandler)
	blogGroup.DELETE("/:id", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleRoot, domain.RoleUser, domain.RoleAdmin), bc.DeleteBlogHandler)
	blogGroup.POST("/", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleRoot, domain.RoleUser, domain.RoleAdmin), bc.GetBlogHandler)
	blogGroup.GET("/:id", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleRoot, domain.RoleUser, domain.RoleAdmin), bc.GetBlogByIDHandler)
	blogGroup.POST("/like", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleRoot, domain.RoleUser, domain.RoleAdmin), bc.BlogLikeHandler)
	blogGroup.POST("/dislike", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleRoot, domain.RoleUser, domain.RoleAdmin), bc.BlogDisLikeHandler)
	blogGroup.POST("/generate-content", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleRoot, domain.RoleUser, domain.RoleAdmin), bc.GenerateContentHandler)
	blogGroup.POST("/review-content", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleRoot, domain.RoleUser, domain.RoleAdmin), bc.ReviewContentHandler)
	blogGroup.POST("/generate-topic", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleRoot, domain.RoleUser, domain.RoleAdmin), bc.GenerateTopicHandler)

	// router for comment
	blogGroup.POST("/comment/:blogId", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleRoot, domain.RoleUser, domain.RoleAdmin), bc.HandleCreateComment)
	blogGroup.PUT("/comment/:blogId/:commentId", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleRoot, domain.RoleUser, domain.RoleAdmin), bc.HandleUpdateComment)
	blogGroup.DELETE("/comment/:blogId/:commentId", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleRoot, domain.RoleUser, domain.RoleAdmin), bc.HandleDeleteComment)
}
