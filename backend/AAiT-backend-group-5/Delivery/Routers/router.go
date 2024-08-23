package routers

import (
	"context"

	config "github.com/aait.backend.g5.main/backend/Config"
	middlewares "github.com/aait.backend.g5.main/backend/Delivery/middlewares"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	repository "github.com/aait.backend.g5.main/backend/Repository"

	"github.com/gin-gonic/gin"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"

)

func Setup(env *config.Env, db interfaces.Database, gin *gin.Engine) {
	jwt_service := infrastructure.NewJwtService(env)
	session_repo := repository.NewSessionRepository(db)
	jwtMiddleware := middlewares.NewJwtAuthMiddleware(jwt_service, session_repo)
	redisClient := config.NewRedisClient(*env, context.Background())

	publicRoute := gin.Group("")
	protectedRoute := gin.Group("")
	adminRoute := gin.Group("")
	refreshRoute := publicRoute.Group("")

	refreshRoute.Use(jwtMiddleware.JWTRefreshAuthMiddelware())
	protectedRoute.Use(jwtMiddleware.JWTAuthMiddelware())
	adminRoute.Use(
		jwtMiddleware.JWTAuthMiddelware(),
		middlewares.AuthenticateAdmin(),
	)

	NewAuthenticationRouter(env, db, publicRoute)
	NewForgotPasswordRouter(env, db, protectedRoute)
	NewLogoutRouter(env, db, protectedRoute)
	NewRefreshRouter(env, db, refreshRoute)

	NewUserProfileRouter(db, protectedRoute)
	NewPromoteDemoteRouter(db, adminRoute)

	NewBlogRouter(env, db, protectedRoute, redisClient)
	NewBlogCommentRouter(env, db, protectedRoute, redisClient)

	NewAISuggestionRouter(db, env, protectedRoute)
}
