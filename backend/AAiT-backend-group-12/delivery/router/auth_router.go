package router

import (
	"blog_api/delivery/controllers"
	"blog_api/delivery/env"
	"blog_api/domain"
	"blog_api/infrastructure/cryptography"
	"blog_api/infrastructure/fs"
	jwt_service "blog_api/infrastructure/jwt"
	mail_service "blog_api/infrastructure/mail"
	"blog_api/infrastructure/middleware"
	google_auth "blog_api/infrastructure/oauth"
	"blog_api/infrastructure/utils"
	"blog_api/repository"
	"blog_api/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewAuthRouter initalizes the controllers, usecases and repositories before setting up the auth routes
func NewAuthRouter(collection *mongo.Collection, authGroup *gin.RouterGroup, cacheClient *redis.Client) {
	userRepository := repository.NewUserRepository(collection)
	cacheRepoistory := repository.NewCacheRepository(cacheClient)
	jwtService := jwt_service.NewJWTService(env.ENV.JWT_SECRET_TOKEN)
	mailService := mail_service.NewMailService(env.ENV.SMTP_GMAIL, env.ENV.SMTP_PASSWORD)
	hashService := cryptography.NewHashingService()
	usecase := usecase.NewUserUsecase(
		userRepository,
		cacheRepoistory,
		utils.GenerateToken,
		mailService,
		jwtService,
		hashService,
		google_auth.VerifyIdToken,
		fs.DeleteFile,
		env.ENV,
	)

	controller := controllers.NewAuthController(usecase, fs.DeleteFile)

	authGroup.POST("/signup", controller.HandleSignup)
	authGroup.GET("/verify/email/:username/:token", controller.HandleVerifyEmail)

	authGroup.POST("/login", controller.HandleLogin)
	authGroup.POST("/renew-token", controller.HandleRenewAccessToken)
	authGroup.POST("/logout", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleUser, domain.RoleAdmin, domain.RoleRoot), controller.HandleLogout)

	authGroup.PATCH("/:username", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleUser), controller.HandleUpdateUser)

	authGroup.PATCH("/promote/:username", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleAdmin, domain.RoleRoot), controller.HandlePromoteUser)
	authGroup.PATCH("/demote/:username", middleware.AuthMiddlewareWithRoles(jwtService, cacheRepoistory, domain.RoleRoot), controller.HandleDemoteUser)

	authGroup.POST("/forgot-password", controller.HandleInitResetPassword)
	authGroup.POST("/reset-password", controller.HandleResetPassword)

	authGroup.POST("/google/login", controller.HandleGoogleLogin)
	authGroup.POST("/google/signup", controller.HandleGoogleSignup)
}
