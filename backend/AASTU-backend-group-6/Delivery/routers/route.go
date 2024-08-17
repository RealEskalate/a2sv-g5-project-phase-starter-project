package routers

import (
	controllers "blogs/Delivery/controllers"
	infrastructure "blogs/Infrastructure"
	repositories "blogs/Repositories"
	usecases "blogs/Usecases"
	"blogs/mongo"

	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine, config *infrastructure.Config, DB mongo.Database) {

	blog_repo := repositories.NewBlogRepository(DB.Collection(config.BlogCollection), config)
	blog_usecase := usecases.NewBlogUsecase(blog_repo)
	blog_controller := controllers.NewBlogController(blog_usecase)
	blogRouter := server.Group("blogs")
	NewBlogrouter(blogRouter,blog_controller)
	userRouter := server.Group("")
	// NewUserrouter(config, DB , userRouter)
	NewSignupRoute(config, DB, userRouter)

}
