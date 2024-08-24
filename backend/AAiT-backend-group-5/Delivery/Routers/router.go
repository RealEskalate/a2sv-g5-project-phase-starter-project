package routers

import (
	"context"
	"log"
	"path/filepath"

	config "github.com/aait.backend.g5.main/backend/Config"
	"github.com/aait.backend.g5.main/backend/Delivery/middlewares"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	repository "github.com/aait.backend.g5.main/backend/Repository"
	"github.com/gin-gonic/gin"
)

func Setup(env *config.Env, db interfaces.Database, gin *gin.Engine) {
	user_repo := repository.NewUserRepository(db)
	jwt_service := infrastructure.NewJwtService(env)
	oAuthService := infrastructure.NewOAuthService(*env, user_repo)
	session_repo := repository.NewSessionRepository(db)
	jwtMiddleware := middlewares.NewJwtAuthMiddleware(jwt_service, session_repo, oAuthService)
	redisClient := config.NewRedisClient(*env, context.Background())

	projectRoot, err := filepath.Abs(filepath.Join(""))
	if err != nil {
		log.Fatalf("Error getting project root path: %v", err)
	}

	templatesDir := filepath.Join(projectRoot, "Infrastructure", "web", "templates", "*.html")
	staticDir := filepath.Join(projectRoot, "Infrastructure", "web", "static")

	gin.LoadHTMLGlob(templatesDir)
	gin.Static("/static", staticDir)

	publicRoute := gin.Group("")
	protectedRoute := gin.Group("")
	adminRoute := gin.Group("")
	refreshRoute := publicRoute.Group("")

	publicRoute.Use(middlewares.NewSecureMiddleware())
	publicRoute.Use(middlewares.RateLimitMiddleware(redisClient, context.Background()))

	refreshRoute.Use(jwtMiddleware.JWTRefreshAuthMiddelware())
	protectedRoute.Use(jwtMiddleware.JWTAuthMiddelware())

	adminRoute.Use(
		jwtMiddleware.JWTAuthMiddelware(),
		middlewares.AuthenticateAdmin(),
	)

	NewAuthenticationRouter(env, db, publicRoute)
	NewOAuthRouter(env, db, *publicRoute)
	NewForgotPasswordRouter(env, db, publicRoute)
	NewLogoutRouter(env, db, protectedRoute)
	NewRefreshRouter(env, db, refreshRoute)
	NewUserProfileRouter(db, *env, protectedRoute)
	NewPromoteDemoteRouter(db, adminRoute)
	NewBlogRouter(env, db, protectedRoute, redisClient)
	NewBlogCommentRouter(env, db, protectedRoute, redisClient)
	NewAISuggestionRouter(db, env, protectedRoute)
}
