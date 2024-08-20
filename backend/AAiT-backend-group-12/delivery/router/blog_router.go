package router

import (
	"blog_api/repository"
	"blog_api/delivery/controllers"
	"blog_api/delivery/env"
	jwt_service "blog_api/infrastructure/jwt"
	"blog_api/infrastructure/middleware"
	ai_service "blog_api/infrastructure/ai"
	"blog_api/usecase"
	"blog_api/domain"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBlogRouter(collection *mongo.Collection, blogGroup *gin.RouterGroup) {
	br := repository.NewBlogRepository(collection)
	bu := usecase.NewBlogUseCase(br, time.Second*100, ai_service.NewAIService(env.ENV.GEMINI_API_KEY))

	bc := controllers.NewBlogController(bu)

	blogGroup.POST("/create", middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, domain.RoleRoot,  domain.RoleUser, domain.RoleAdmin), bc.CreateBlogHandler)
	blogGroup.PUT("/:id", middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, domain.RoleRoot,  domain.RoleUser, domain.RoleAdmin), bc.UpdateBlogHandler)
	blogGroup.DELETE("/:id", middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, domain.RoleRoot,  domain.RoleUser, domain.RoleAdmin), bc.DeleteBlogHandler)
	blogGroup.POST("/",middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, domain.RoleRoot,  domain.RoleUser, domain.RoleAdmin),bc.GetBlogHandler)
	blogGroup.GET("/:id",middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, domain.RoleRoot,  domain.RoleUser, domain.RoleAdmin), bc.GetBlogByIDHandler)
	blogGroup.POST("/like", middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, domain.RoleRoot,  domain.RoleUser, domain.RoleAdmin), bc.BlogLikeHandler)
	blogGroup.POST("/dislike", middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, domain.RoleRoot,  domain.RoleUser, domain.RoleAdmin), bc.BlogDisLikeHandler)
	blogGroup.POST("/generate-content",middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, domain.RoleRoot,  domain.RoleUser, domain.RoleAdmin), bc.GenerateContentHandler)
	blogGroup.POST("/review-content",middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, domain.RoleRoot,  domain.RoleUser, domain.RoleAdmin), bc.ReviewContentHandler)
	blogGroup.POST("/generate-topic",middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, domain.RoleRoot,  domain.RoleUser, domain.RoleAdmin), bc.GenerateTopicHandler)
}
