package route

import (
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/middleware"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, gin *gin.Engine, cloudinary *cloudinary.Cloudinary) {

	// Error handling
	gin.Use(middleware.ErrorHandlerMiddleware())

	publicRouter := gin.Group("")

	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewOAuthRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)
	NewPublicBlogsRouter(env, timeout, db, publicRouter)
	NewPublicResetPasswordRouter(env, timeout, db, publicRouter)

	// Static files
	// NewPublicFileRouter(env, publicRouter)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	// All Protected APIs
	NewProtectedBlogsRouter(env, timeout, db, protectedRouter)
	NewProfileRouter(env, timeout, db, protectedRouter, cloudinary)
	NewChatRouter(env, timeout, protectedRouter)
}
