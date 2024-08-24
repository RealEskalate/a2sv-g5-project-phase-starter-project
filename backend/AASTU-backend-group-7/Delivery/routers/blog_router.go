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
		blogrepo := Repositories.NewBlogrepository(BlogCollections)
		blogusecase := usecases.NewBlogUseCase(blogrepo)
		blogcontroller := controllers.NewBlogController(blogusecase)

		postRouter.POST("/create", blogcontroller.CreateBlog)
		postRouter.GET("/get", blogcontroller.GetUserPosts)
		postRouter.GET("/get/:slug", blogcontroller.GetPostBySlug)
		postRouter.GET("/getbyid/:id", blogcontroller.GetPostByID)
		postRouter.GET("/getbyauthor/:authorID", blogcontroller.GetPostByAuthorID)
		postRouter.PUT("/update/:id", blogcontroller.UpdatePostByID)
		// delete post
		postRouter.DELETE("/delete/:id", blogcontroller.DeletePost)
		
		postRouter.PUT("/tags/:id", blogcontroller.AddTagToPost)
		

		// get tags
		postRouter.GET("/tags/:id", blogcontroller.GetTags)
		//get comments
		postRouter.GET("/comments/:id", blogcontroller.GetComments)
		// get all posts
		postRouter.GET("/all", blogcontroller.GetAllPosts)
		// like post
		postRouter.POST("/like/:id", blogcontroller.LikePost)
		// dislike post
		postRouter.POST("/dislike/:id", blogcontroller.DislikePost)
		// search posts
		postRouter.GET("/search", blogcontroller.SearchPosts)


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
		// comment on comment
		commentRouter.POST("/comment/:id", commentcontroller.CommentOnComment)
	}

	tagRouter := Router.Group("/tags", auth_middleware.AuthMiddleware())
	{
		tagRepo := Repositories.NewTagRepository(BlogCollections)
		tagUsecase := usecases.NewTagsUseCase(tagRepo)
		tagController := controllers.NewTagsController(tagUsecase)
		//get all tags
		// admin routes
		adminRouter := tagRouter.Group("/admin",  auth_middleware.IsAdminMiddleware())
		{
			adminRouter.POST("/create", tagController.CreateTag)
			//delete tag
			adminRouter.DELETE("/delete/:slug", tagController.DeleteTag)
		}

		tagRouter.GET("/all", tagController.GetAllTags)
		// get tags by slug
		tagRouter.GET("/get/:slug", tagController.GetTagBySlug)
		// get posts of a tag by slug
		tagRouter.GET("/posts/:slug", tagController.GetPostsByTag)

	}
	
}
