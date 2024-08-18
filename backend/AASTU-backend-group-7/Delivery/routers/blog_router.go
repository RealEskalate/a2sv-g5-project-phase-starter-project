package routers

import (
	"blogapp/Delivery/controllers"
	"blogapp/Infrastructure/auth_middleware"
	"blogapp/Repositories"
	usecases "blogapp/UseCases"
)

func BlogRouter() {
	postRouter := Router.Group("/blog", auth_middleware.AuthMiddleware())
	{
		blogrepo := Repositories.NewBlogRepository(BlogCollections)
		blogusecase := usecases.NewBlogUseCase(blogrepo)
		blogcontroller := controllers.NewBlogController(blogusecase)

		postRouter.POST("/create", blogcontroller.CreateBlog)
		postRouter.GET("/get", blogcontroller.GetUserPosts)
		postRouter.GET("/get/:slug", blogcontroller.GetPostBySlug)
		postRouter.GET("/getbyid/:id", blogcontroller.GetPostByID)
		postRouter.GET("/getbyauthor/:authorID", blogcontroller.GetPostByAuthorID)
		postRouter.PUT("/update/:id", blogcontroller.UpdatePostByID)


	}

	commentRouter := Router.Group("/comment", auth_middleware.AuthMiddleware())
	{
		commentrepo := Repositories.NewCommentRepository(BlogCollections)
		commentusecase := usecases.NewCommentUseCase(commentrepo)
		commentcontroller := controllers.NewCommentController(commentusecase)

		commentRouter.POST("/:id", commentcontroller.CommentOnPost)
	}
}
