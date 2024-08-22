package routes

import (
	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/repository"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
)

func NewPopularityRoutes(group *gin.RouterGroup, blog_collection database.CollectionInterface) {

	mongoCollection := blog_collection.(*database.MongoCollection).Collection

	popularityRepo := repository.NewBlogPopularityRepository(mongoCollection)

	popularityUsecase := usecase.NewBlogPopularityUsecase(popularityRepo)

	popularityCtrl := controller.NewPopularityController(popularityUsecase)

	group.GET("/popular-blogs", popularityCtrl.GetPopularBlogs())
}
