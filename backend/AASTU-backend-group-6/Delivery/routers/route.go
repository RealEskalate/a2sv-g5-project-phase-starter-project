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

func Router(server *gin.Engine, config *infrastructure.Config, DB mongo.Database) {

	blog_repo := repositories.NewBlogRepository(DB.Collection(config.BlogCollection), DB.Collection(config.UserCollection), *config)
	validator := domain.NewValidator()
	blog_usecase := usecases.NewBlogUsecase(blog_repo)
	blog_controller := controllers.NewBlogController(blog_usecase,  validator)
	blogRouter := server.Group("blogs")
	NewBlogrouter(blogRouter, blog_controller)
	userRouter := server.Group("")
	// NewUserrouter(config, DB , userRouter)
	exp := time.Duration(time.Now().Add(time.Duration(config.RefreshTokenExpiryHour)).Unix())
	NewSignupRoute(config, DB, userRouter)
	NewRefreshTokenRouter(config, exp, DB, userRouter)
	NewLoginRouter(config, exp, DB, userRouter)
}
