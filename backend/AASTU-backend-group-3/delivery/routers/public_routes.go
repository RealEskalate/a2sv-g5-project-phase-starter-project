package routers
// import (
//  	"group3-blogApi/config/db"
// 	"github.com/gin-gonic/gin"
// 	"group3-blogApi/repository"
// 	"group3-blogApi/usecase"
// 	"group3-blogApi/delivery/controllers"
	
// )

// func SetUpPublicRoutes(router *gin.Engine) {	


// 	blogRepo := repository.NewBlogRepositoryImpl(db.BlogCollection)
//     blogUsecase := usecase.NewBlogUsecase(blogRepo)
//     blogController := controllers.NewBlogController(blogUsecase)

// 	public := router.Group("/")
// 	{
// 		public.GET("/blogs", blogController.GetBlogs)                    // Get all blogs with pagination and sorting
// 		public.GET("/blogs/:id", blogController.GetBlog)                  // Get blog by ID
// 		public.GET("/users/:id/blogs", blogController.GetUserBlogs)       // Get blogs by user
// 		public.GET("/blogs/search", blogController.SearchBlogs)       // Search blogs
// 		public.GET("/blogs/filter", blogController.FilterBlogs)
// 		       // Filter blogs
// 	}

// }



import (
	"group3-blogApi/config/db"
	"group3-blogApi/delivery/controllers"
	//infrastracture "group3-blogApi/infrastracture"
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
		public.GET("/blogs", blogController.GetBlogs)                    // Get all blogs with pagination and sorting
		public.GET("/blog/:id", blogController.GetBlog)                  // Get blog by ID
		public.GET("/users/:id/blogs", blogController.GetUserBlogs)       // Get blogs by user
		public.GET("/blogs/search", blogController.SearchBlogs)       // Search blogs
		public.GET("/blogs/filter", blogController.FilterBlogs)       // Filter blogs
	}
}