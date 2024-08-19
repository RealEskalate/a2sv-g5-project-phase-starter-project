package route

import (
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/middleware"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")

	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)
	NewVerificationRouter(env, timeout, db, publicRouter)
	NewPublicBlogsRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	// All Private APIs
	NewPrivateBlogsRouter(env, timeout, db, protectedRouter)
	NewProfileRouter(env, timeout, db, protectedRouter)
}
