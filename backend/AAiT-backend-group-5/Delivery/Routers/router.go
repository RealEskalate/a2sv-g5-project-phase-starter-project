package routers

import (
	"context"

	config "github.com/aait.backend.g5.main/backend/Config"
	middlewares "github.com/aait.backend.g5.main/backend/Delivery/middlewares"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	repository "github.com/aait.backend.g5.main/backend/Repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *config.Env, db mongo.Database, gin *gin.Engine) {
	jwt_service := infrastructure.NewJwtService(env)
	session_repo := repository.NewSessionRepository(&db)
	jwtMiddleware := middlewares.NewJwtAuthMiddleware(jwt_service, session_repo)

	publicRouter := gin.Group("")
	protectedRouter := gin.Group("")
	adminRouter := gin.Group("")
	redisClient := config.NewRedisClient(*env, context.Background())

	protectedRouter.Use(jwtMiddleware.JWTAuthMiddelware())

	adminRouter.Use(
		jwtMiddleware.JWTAuthMiddelware(),
		middlewares.AuthenticateAdmin(),
	)

	NewAuthenticationRouter(env, db, publicRouter)
	NewForgotPasswordRouter(env, db, protectedRouter)
	NewLogoutRouter(env, db, protectedRouter)
	NewPromoteDemoteRouter(db, adminRouter)
	NewBlogRouter(env, db, protectedRouter, redisClient)
	NewBlogCommentRouter(env, db, protectedRouter, redisClient)
}
