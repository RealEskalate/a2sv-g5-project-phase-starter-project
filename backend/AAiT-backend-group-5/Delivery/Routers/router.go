package routers

import (
	"context"
	"log"
	"path/filepath"

	"github.com/aait.backend.g5.main/backend/Config"
	"github.com/aait.backend.g5.main/backend/Delivery/middlewares"
	"github.com/aait.backend.g5.main/backend/Infrastructure"
	"github.com/aait.backend.g5.main/backend/Repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *config.Env, db mongo.Database, gin *gin.Engine) {
	user_repo := repository.NewUserRepository(&db)
	jwt_service := infrastructure.NewJwtService(env)
	oAuthService := infrastructure.NewOAuthService(*env, user_repo)
	session_repo := repository.NewSessionRepository(&db)
	jwtMiddleware := middlewares.NewJwtAuthMiddleware(jwt_service, session_repo, oAuthService)

	projectRoot, err := filepath.Abs(filepath.Join(""))
	if err != nil {
		log.Fatalf("Error getting project root path: %v", err)
	}

	templatesDir := filepath.Join(projectRoot, "Infrastructure", "web", "templates", "*.html")
	staticDir := filepath.Join(projectRoot, "Infrastructure", "web", "static")

	gin.LoadHTMLGlob(templatesDir)
	gin.Static("/static", staticDir)

	publicRouter := gin.Group("")
	protectedRouter := gin.Group("")
	adminRouter := gin.Group("")

	redisClient := config.NewRedisClient(*env, context.Background())

	refreshGroup := publicRouter.Group("")
	refreshGroup.Use(jwtMiddleware.JWTRefreshAuthMiddelware())

	protectedRouter.Use(jwtMiddleware.JWTAuthMiddelware())

	adminRouter.Use(
		jwtMiddleware.JWTAuthMiddelware(),
		middlewares.AuthenticateAdmin(),
	)

	// Set up routes
	NewOAuthRouter(env, db, *publicRouter)
	NewAuthenticationRouter(env, db, publicRouter)
	NewForgotPasswordRouter(env, db, publicRouter)
	NewLogoutRouter(env, db, protectedRouter)
	NewRefreshRouter(env, db, refreshGroup)
	NewPromoteDemoteRouter(db, adminRouter)
	NewBlogRouter(env, db, protectedRouter, redisClient)
	NewBlogCommentRouter(env, db, protectedRouter, redisClient)
	NewAISuggestionRouter(db, env, protectedRouter)
}
