package routers

import (
	"group3-blogApi/config/db"
	"group3-blogApi/delivery/controllers"
	"group3-blogApi/infrastracture"
	"group3-blogApi/repository"
	"group3-blogApi/usecase"

	"github.com/gin-gonic/gin"
)

func SetUpBlog(router *gin.Engine) {

	blogRepo := repository.NewBlogRepositoryImpl(db.BlogCollection, db.CommentCollection, db.LikeCollection)
    blogUsecase := usecase.NewBlogUsecase(blogRepo)
    blogController := controllers.NewBlogController(blogUsecase)

	blogs := router.Group("/blogs")
	blogs.Use(infrastracture.AuthMiddleware())
	{
		blogs.POST("/new",infrastracture.EligibilityMiddleware(), blogController.CreateBlog)
		blogs.DELETE("/:id", infrastracture.EligibilityMiddleware(), blogController.DeleteBlog)
		blogs.PUT("/:id", infrastracture.EligibilityMiddleware(),blogController.UpdateBlog)		
	}
		

}


