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
		

		// get tags
		postRouter.GET("/tags/:id", blogcontroller.GetTags)
		//get comments
		postRouter.GET("/comments/:id", blogcontroller.GetComments)
		// get all posts
		postRouter.GET("/all", blogcontroller.GetAllPosts)


	}

	commentRouter := Router.Group("/comment", auth_middleware.AuthMiddleware())
	{
		commentrepo := Repositories.NewCommentRepository(BlogCollections)
		commentusecase := usecases.NewCommentUseCase(commentrepo)
		commentcontroller := controllers.NewCommentController(commentusecase)

		commentRouter.POST("/:id", commentcontroller.CommentOnPost)
		// comment by id
		commentRouter.GET("/get/:id", commentcontroller.GetCommentByID)
		//	edit comment
		commentRouter.PUT("/edit/:id", commentcontroller.EditComment)
		// get user comments
		commentRouter.GET("/getauthorcomments/:id", commentcontroller.GetUserComments)
		// get user's comments
		commentRouter.GET("/getmycomments", commentcontroller.GetMyComments)
		// delete comment
		commentRouter.DELETE("/delete/:id", commentcontroller.DeleteComment)
	}

	tagRouter := Router.Group("/tags", auth_middleware.AuthMiddleware(), auth_middleware.IsAdminMiddleware())
	{
		tagRepo := Repositories.NewTagsRepository(BlogCollections)
		tagUsecase := usecases.NewTagsUseCase(tagRepo)
		tagController := controllers.NewTagsController(tagUsecase)

		tagRouter.POST("/create", tagController.CreateTag)
		//delete tag
		tagRouter.DELETE("/delete/:id", tagController.DeleteTag)
		// add tag to post
		
	}
}
