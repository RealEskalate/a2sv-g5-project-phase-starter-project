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
		blogrepo := Repositories.NewBlogRepository(BlogCollections.Posts, BlogCollections.Comments, BlogCollections.Tags, BlogCollections.Likes, BlogCollections.Users)
		blogusecase := usecases.NewBlogUseCase(blogrepo)
		blogcontroller := controllers.NewBlogController(blogusecase)

		blogRouter.POST("/create", blogcontroller.CreateBlog)
		blogRouter.GET("/get", blogcontroller.GetUserPosts)
		blogRouter.GET("/get/:slug", blogcontroller.GetPostBySlug)
		blogRouter.GET("/get/:id", blogcontroller.GetPostByID)
		blogRouter.GET("/get/:authorID", blogcontroller.GetPostByAuthorID)


	}
}
