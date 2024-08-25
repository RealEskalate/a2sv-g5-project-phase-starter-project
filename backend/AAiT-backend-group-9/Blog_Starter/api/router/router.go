package router

import (
	//"Blog_Starter/api/middleware"
	"Blog_Starter/api/middleware"
	"Blog_Starter/config"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// Setup initializes the routes for the application.
func Setup(env *config.Env, timeout time.Duration, db *mongo.Client, gin *gin.Engine) {
    publicRouter := gin.Group("")
	publicRouter.Use(middleware.CaptureResponseMiddleware())
    NewSignupRouter(env, timeout, db, publicRouter)
    NewLoginRouter(env, timeout, db, publicRouter)


	privateRouterBlog := gin.Group("/api/blog")
	privateRouterBlog.Use(middleware.AuthMiddleWare(env.AccessTokenSecret))
	privateRouterBlog.Use(middleware.CaptureResponseMiddleware())
	NewBlogRouter(env, timeout, db, privateRouterBlog)

	popularityRouter := gin.Group("/api/blog/:blog_id")
	privateRouterBlog.Use(middleware.AuthMiddleWare(env.AccessTokenSecret))
	privateRouterBlog.Use(middleware.CaptureResponseMiddleware())
	NewBlogRatingRouter(env, timeout, db, popularityRouter)
	NewBlogCommentRouter(env, timeout, db, popularityRouter)
	NewBlogLikeRouter(env, timeout, db, popularityRouter)

	privateRouterUser := gin.Group("/api/user")
	privateRouterUser.Use(middleware.AuthMiddleWare(env.AccessTokenSecret))
	privateRouterUser.Use(middleware.CaptureResponseMiddleware())
	NewUserRouter(env, timeout, db, privateRouterUser)

	privateRouterRefreshToken := gin.Group("/api")
	privateRouterRefreshToken.Use(middleware.RefreshMiddleware(env.RefreshTokenSecret))
	privateRouterRefreshToken.Use(middleware.CaptureResponseMiddleware())
	NewRefreshTokenRouter(env, timeout, db, privateRouterRefreshToken)
}