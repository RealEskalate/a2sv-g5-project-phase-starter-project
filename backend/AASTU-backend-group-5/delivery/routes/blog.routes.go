package routes

import (
	"log"
	"os"

	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/infrastructure/middleware"
	tokenservice "github.com/RealEskalate/blogpost/infrastructure/token_service"
	"github.com/RealEskalate/blogpost/repository"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func NewBlogRoutes(group *gin.RouterGroup, blog_collection database.CollectionInterface, user_collection database.CollectionInterface) {
	repo := repository.NewBlogRepository(blog_collection)
	blog_usecase := usecase.NewBlogUsecase(repo)

	user_repo := repository.NewUserRepository(user_collection)
	user_usecase := usecase.NewUserUseCase(user_repo)
	ctrl := controller.NewBlogController(blog_usecase, user_usecase)

	//load middlewares
	err := godotenv.Load()
	if err != nil {
		log.Panic(err.Error())
	}
	access_secret := os.Getenv("ACCESSTOKENSECRET")
	if access_secret == "" {
		log.Panic("No accesstoken")
	}

	refresh_secret := os.Getenv("REFRESHTOKENSECRET")
	if refresh_secret == "" {
		log.Panic("No refreshtoken")
	}
	TokenSvc := *tokenservice.NewTokenService(access_secret, refresh_secret)

	LoggedInmiddleWare := middleware.LoggedIn(TokenSvc)
	mustOwn := middleware.RoleBasedAuth(false)

	group.GET("api/allblog", ctrl.GetAllBlogs())
	group.GET("api/search-blog/", ctrl.FilterBlogs())

	group.GET("api/blog/:id", LoggedInmiddleWare, ctrl.GetOneBlog())
	group.POST("api/blog/", LoggedInmiddleWare, ctrl.CreateBlog())
	group.GET("api/UniqueBlog", LoggedInmiddleWare, ctrl.GetUniqueBlog())

	group.GET("api/my-blog", LoggedInmiddleWare, mustOwn, ctrl.GetMyBlogs())
	group.PUT("api/blog/:id", LoggedInmiddleWare, mustOwn, ctrl.UpdateBlog())
	group.DELETE("api/blog/:id", LoggedInmiddleWare, mustOwn, ctrl.DeleteBlog())

}
