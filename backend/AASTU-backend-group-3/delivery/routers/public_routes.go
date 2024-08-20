package routers

import (
	"group3-blogApi/config/db"
	"group3-blogApi/delivery/controllers"
	"group3-blogApi/repository"
	"group3-blogApi/usecase"

	"github.com/gin-gonic/gin"
)

func SetUpPublicRoutes(router *gin.Engine) {	


	blogRepo := repository.NewBlogRepositoryImpl(db.BlogCollection)
    blogUsecase := usecase.NewBlogUsecase(blogRepo)
    blogController := controllers.NewBlogController(blogUsecase)

	public := router.Group("/")
	{

	
		public.GET("/blogs", blogController.GetBlogs)
		public.GET("/blogs/:id", blogController.GetBlogByID)
		public.GET("/users/:id/blogs", blogController.GetUserBlogs)   
		// desired page number and the number of posts per page, and any sorting preferences
		
		



		// public.GET("/blogs", blogController.GetBlogs)                    // Get all blogs with pagination and sorting
		// public.GET("/blog/:id", blogController.GetBlog)                  // Get blog by ID
		// public.GET("/blogs/search", blogController.SearchBlogs)       // Search blogs
		// public.GET("/blogs/filter", blogController.FilterBlogs)       // Filter blogs
	}
}