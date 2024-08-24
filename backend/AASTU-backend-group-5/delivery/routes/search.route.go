package routes

import (
	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/repository"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
)

func NewBookmarkRoutes(group *gin.RouterGroup, bookmarkCollection database.CollectionInterface, BlogCollection database.CollectionInterface) {
	repo := repository.NewBookmarkRepository(bookmarkCollection)
	blog := repository.NewBlogRepository(BlogCollection)
	usecase := usecase.NewBookmarkUseCase(repo, blog)
	ctrl := controller.NewBookmarkController(usecase)

	group.POST("/add/:userID/:blogID", ctrl.AddBookmark())
	group.DELETE("/remove/:userID/:blogID", ctrl.RemoveBookmark())
	group.GET("/getbook/:userID", ctrl.GetUserBookmarks())
}
