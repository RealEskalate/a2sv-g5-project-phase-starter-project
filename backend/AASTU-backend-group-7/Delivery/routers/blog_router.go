package routers

import (
	"blogapp/Delivery/controllers"
	"blogapp/Infrastructure/auth_middleware"
	"blogapp/Repositories"
	usecases "blogapp/UseCases"
)

func BlogRouter() {
	blogRouter := Router.Group("/blog", auth_middleware.AuthMiddleware())
	{
		blogrepo := Repositories.NewBlogRepository(BlogCollections.Blogs)
		blogusecase := usecases.NewBlogUseCase(blogrepo)
		blogcontroller := controllers.NewBlogController(blogusecase)

		blogRouter.POST("/create", blogcontroller.CreateBlog)
	}
}
