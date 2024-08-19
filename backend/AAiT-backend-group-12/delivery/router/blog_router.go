package router

import (
	"blog_api/delivery/controllers"
	"blog_api/delivery/env"
	jwt_service "blog_api/infrastructure/jwt"
	"blog_api/infrastructure/middleware"
	"blog_api/repository"
	"blog_api/usecase"
	"blog_api/domain"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBlogRouter(collection *mongo.Collection, blogGroup *gin.RouterGroup) {
	br := repository.NewBlogRepository(collection)
	bu := usecase.NewBlogUseCase(br, time.Second * 100)

	bc := controllers.NewBlogController(bu)

	blogGroup.POST("/create", middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, domain.RoleRoot,  domain.RoleUser, domain.RoleAdmin), bc.CreateBlogHandler)
	blogGroup.PUT("/:id", middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, domain.RoleRoot,  domain.RoleUser, domain.RoleAdmin), bc.UpdateBlogHandler)
	blogGroup.DELETE("/:id", middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, domain.RoleRoot,  domain.RoleUser, domain.RoleAdmin), bc.DeleteBlogHandler)
	blogGroup.POST("/", bc.GetBlogHandler)
	blogGroup.GET("/:id", bc.GetBlogByIDHandler)
	blogGroup.POST("/update-popularity", bc.TrackBlogPopularityHandler)
}
