package routers

import (
	"time"

	config "github.com/aait.backend.g5.main/backend/Config"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *config.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	protectedRouter := gin.Group("")
	adminRouter := gin.Group("")

	jwt_service := infrastructure.NewJwtService(env)
	protectedRouter.Use(infrastructure.JWTAuthMiddelware(jwt_service))

	adminRouter.Use(
		infrastructure.JWTAuthMiddelware(jwt_service),
		infrastructure.AuthenticateAdmin(),
	)

	NewAuthenticationRouter(env, db, publicRouter)
	NewForgotPasswordRouter(env, db, protectedRouter)
	NewPromoteDemoteRouter(db, adminRouter)
}
