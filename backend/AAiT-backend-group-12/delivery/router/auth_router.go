package router

import (
	"blog_api/delivery/controllers"
	"blog_api/delivery/env"
	"blog_api/infrastructure/cryptography"
	jwt_service "blog_api/infrastructure/jwt"
	mail_service "blog_api/infrastructure/mail"
	"blog_api/infrastructure/middleware"
	"blog_api/infrastructure/utils"
	"blog_api/repository"
	"blog_api/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthRouter(collection *mongo.Collection, authGroup *gin.RouterGroup, cacheClient *redis.Client) {
	userRepository := repository.NewUserRepository(collection)
	cacheRepoistory := repository.NewCacheRepository(cacheClient)
	usecase := usecase.NewUserUsecase(
		userRepository,
		cacheRepoistory,
		utils.GenerateToken,
		mail_service.EmailVerificationTemplate,
		mail_service.PasswordResetTemplate,
		mail_service.SendMail,
		jwt_service.SignJWTWithPayload,
		jwt_service.GetTokenType,
		jwt_service.GetUsername,
		jwt_service.GetExpiryDate,
		jwt_service.ValidateAndParseToken,
		cryptography.HashString,
		cryptography.ValidateHashedString,
		env.ENV,
	)

	controller := controllers.NewAuthController(usecase)

	authGroup.POST("/signup", controller.HandleSignup)
	authGroup.GET("/verify/email/:username/:token", controller.HandleVerifyEmail)

	authGroup.POST("/login", controller.HandleLogin)
	authGroup.POST("/renew-token", controller.HandleRenewAccessToken)
	authGroup.POST("/logout", middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, cacheRepoistory, "user", "admin", "root"), controller.HandleLogout)

	authGroup.PATCH("/:username", middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, cacheRepoistory, "user"), controller.HandleUpdateUser)

	authGroup.PATCH("/promote/:username", middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, cacheRepoistory, "admin", "root"), controller.HandlePromoteUser)
	authGroup.PATCH("/demote/:username", middleware.AuthMiddlewareWithRoles(env.ENV.JWT_SECRET_TOKEN, jwt_service.ValidateAndParseToken, cacheRepoistory, "root"), controller.HandleDemoteUser)

	authGroup.POST("/forgot-password", controller.HandleInitResetPassword)
	authGroup.POST("/reset-password", controller.HandleResetPassword)
}
