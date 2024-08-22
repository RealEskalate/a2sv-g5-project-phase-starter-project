package routes

import (
	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/repository"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
)

func NewBlogRoutes(group *gin.RouterGroup, blog_collection database.CollectionInterface, user_collection database.CollectionInterface) {
	repo := repository.NewBlogRepository(blog_collection)
	blog_usecase := usecase.NewBlogUsecase(repo)

	user_repo := repository.NewUserRepository(user_collection)
	user_usecase := usecase.NewUserUseCase(user_repo)
	ctrl := controller.NewBlogController(blog_usecase, user_usecase)

	group.GET("/blog", ctrl.GetAllBlogs())
	group.GET("/my-blog", ctrl.GetMyBlogs())
	group.GET("/blog/:id", ctrl.GetOneBlog())
	group.POST("/blog/", ctrl.CreateBlog())
	group.PUT("/blog/:id", ctrl.UpdateBlog())
	group.DELETE("/blog/:id", ctrl.DeleteBlog())
	group.GET("/search-blog/", ctrl.FilterBlogs())
}
