package router

import (
	"blog_api/delivery/controllers"
	"blog_api/delivery/env"
	jwt_service "blog_api/infrastructure/jwt"
	"blog_api/infrastructure/middleware"
	"blog_api/repository"
	"blog_api/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthRouter(collection *mongo.Collection, authGroup *gin.RouterGroup, redisClient *redis.Client) {
	userRepository := repository.NewUserRepository(collection)
	redisRepoistory := repository.NewRedisRepository(redisClient)
	usecase := usecase.NewUserUsecase(userRepository, redisRepoistory)
	controller := controllers.NewAuthController(usecase)

	authGroup.POST("/signup", controller.HandleSignup)
	authGroup.POST("/login", controller.HandleLogin)
	authGroup.PATCH("/:username", middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, "user"), controller.HandleUpdateUser)
	authGroup.POST("/renew-token", controller.HandleRenewAccessToken)
	authGroup.PATCH("/promote/:username", middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, "admin", "root"), controller.HandlePromoteUser)
	authGroup.PATCH("/demote/:username", middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, "root"), controller.HandleDemoteUser)
	authGroup.GET("/verify/email/:username/:token", controller.HandleVerifyEmail)

	authGroup.POST("/forgot-password", controller.HandleInitResetPassword)
	authGroup.POST("/reset-password", controller.HandleResetPassword)
}
