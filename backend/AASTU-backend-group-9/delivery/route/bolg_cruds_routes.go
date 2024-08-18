package route

import (
	"blog/config"
	"blog/database"
	"blog/delivery/controller"

	// "blog/domain"
	"blog/repository"
	"blog/usecase"
	"time"

	"blog/delivery/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(env *config.Env, timeout time.Duration, db database.Database, router *gin.RouterGroup) {

	blogRepo := repository.NewBlogRepository(db, "blogs")
	blogUse := usecase.NewBlogUsecase(blogRepo, timeout)
	blogController := &controller.BlogController{
		BlogUsecase: blogUse,
		Env:         env,
	}
	blogRoutes := router.Group("/blogs")
	{
		blogRoutes.Use(middleware.AuthMidd) // Protect the routes with authentication middleware
		blogRoutes.POST("/", blogController.CreateBlog)
		blogRoutes.GET("/:id", blogController.GetBlogByID)
		blogRoutes.GET("/", blogController.GetAllBlogs)
		blogRoutes.PUT("/:id", blogController.UpdateBlog)
		blogRoutes.DELETE("/:id", blogController.DeleteBlog)
	}
}
