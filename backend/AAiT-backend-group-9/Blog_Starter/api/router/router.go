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

    NewSignupRouter(env, timeout, db, publicRouter)
    NewLoginRouter(env, timeout, db, publicRouter)


	privateRouterBlog := gin.Group("/api/blog")
	privateRouterBlog.Use(middleware.AuthMiddleWare(env.AccessTokenSecret))
	
	NewBlogRouter(env, timeout, db, privateRouterBlog)
	NewBlogRatingRouter(env, timeout, db, privateRouterBlog)
	NewBlogCommentRouter(env, timeout, db, privateRouterBlog)
	NewBlogLikeRouter(env, timeout, db, privateRouterBlog)

	privateRouterUser := gin.Group("/api/user")
	privateRouterUser.Use(middleware.AuthMiddleWare(env.AccessTokenSecret))
	NewUserRouter(env, timeout, db, privateRouterUser)

	privateRouterRefreshToken := gin.Group("/api")
	privateRouterRefreshToken.Use(middleware.RefreshMiddleware(env.RefreshTokenSecret))
	NewRefreshTokenRouter(env, timeout, db, privateRouterRefreshToken)
}