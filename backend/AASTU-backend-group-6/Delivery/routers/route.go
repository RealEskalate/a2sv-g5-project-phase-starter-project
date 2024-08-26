package routers

import (
	domain "blogs/Domain"
	infrastructure "blogs/Infrastructure"
	repositories "blogs/Repositories"
	usecases "blogs/Usecases"
	"blogs/mongo"
	"time"

	"blogs/Delivery/controllers"

	"github.com/gin-gonic/gin"
)

func Router(server *gin.RouterGroup, config *infrastructure.Config, DB mongo.Database) {

	blog_repo := repositories.NewBlogRepository(DB.Collection(config.BlogCollection), DB.Collection(config.UserCollection), DB.Collection(config.CommentCollection), *config)
	validator := domain.NewValidator()
	idConverter := domain.NewIdConverter()
	blog_usecase := usecases.NewBlogUsecase(blog_repo, idConverter)
	blog_controller := controllers.NewBlogController(blog_usecase, validator)
	blogRouter := server.Group("blogs")
	authHandler := infrastructure.NewAuthMiddleware(*config).AuthenticationMiddleware()
	redisClient :=infrastructure.NewCacheService(*config)

	NewBlogrouter(blogRouter, blog_controller, authHandler, redisClient)
	authRoute := server.Group("auth")
	// NewUserrouter(config, DB , userRouter)
	exp := time.Duration(time.Now().Add(time.Duration(config.RefreshTokenExpiryHour)).Unix())
	NewSignupRoute(config, DB, authRoute)
	NewRefreshTokenRouter(config, exp, DB, authRoute)
	NewLoginRouter(config, exp, DB, authRoute)

	logout := server.Group("")
	logout.Use(authHandler)
	NewLogoutRouter(config, exp, DB, logout)

	oauthRoute := server.Group("")
	NewOauthRoute(config, DB, oauthRoute)

	aiRoute := server.Group("")
	aiRoute.Use(authHandler)
	NewAIRoute(config, DB, aiRoute)

	userroute := server.Group("/user")
	userroute.Use(authHandler)
	NewUserrouter(config, DB, userroute)

	

}
