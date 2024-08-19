package routes

import (
	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/repository"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
)

func NewBlogRoutes(group *gin.RouterGroup , blog_collection database.CollectionInterface , user_collection database.CollectionInterface) {
	repo := repository.NewBlogRepository(blog_collection)
	blog_usecase := usecase.NewBlogUsecase(repo)

	user_repo := repository.NewUserRepository(user_collection)
	user_usecase := usecase.NewUserUseCase(user_repo)
	ctrl := controller.NewBlogController(blog_usecase , user_usecase)

	group.GET("api/blog" , ctrl.GetAllBlogs())
	group.GET("api/my-blog" , ctrl.GetMyBlogs())
	group.GET("api/blog/:id" , ctrl.GetOneBlog())
	group.POST("api/blog/" , ctrl.CreateBlog())
	group.PUT("api/blog/", ctrl.UpdateBlog())
	group.DELETE("api/blog/" , ctrl.DeleteBlog())
	group.GET("api/search-blog/" , ctrl.FilterBlogs())
}