package routers

import (
	"group3-blogApi/config/db"
	"group3-blogApi/delivery/controllers"
	"group3-blogApi/repository"
	"group3-blogApi/usecase"

	"github.com/gin-gonic/gin"
)

func SetUpPublicRoutes(router *gin.Engine) {	

	db.CreateTextIndex(db.BlogCollection)
	blogRepo := repository.NewBlogRepositoryImpl(db.BlogCollection, db.CommentCollection, db.LikeCollection)
    blogUsecase := usecase.NewBlogUsecase(blogRepo)
    blogController := controllers.NewBlogController(blogUsecase)
	
	public := router.Group("/")
	{

		public.GET("/blogs", blogController.GetBlogs)
		public.GET("/blogs/:id", blogController.GetBlogByID)
		public.GET("/users/:id/blogs", blogController.GetUserBlogs)   
	
	}
}