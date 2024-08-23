package routes

import (
	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/repository"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
)

// In `routes/routes.go`
func NewBookmarkRoutes(group *gin.RouterGroup, bookmarkCollection database.CollectionInterface, userCollection database.CollectionInterface) {
	repo := repository.NewBookmarkRepository(bookmarkCollection)
	usecase := usecase.NewBookmarkUseCase(repo)
	ctrl := controller.NewBookmarkController(usecase)

	group.POST("/add/:userID/:blogID", ctrl.AddBookmark())
	group.DELETE("/remove/:userID/:blogID", ctrl.RemoveBookmark())
	group.GET("/getbook/:userID", ctrl.GetUserBookmarks())
}

// In `routes/routes.go`
// func NewSearchHistoryRoutes(group *gin.RouterGroup, searchHistoryCollection database.CollectionInterface) {
// 	//repo:= repository.NewSearchHistoryRepository(searchHistoryCollection)
// 	usecase := usecase.NewSearchHistoryUseCase(nil)
// 	ctrl := controller.NewSearchHistoryController(usecase)

// 	group.POST("/log/:userID", ctrl.LogSearch())
// 	group.GET("/history/:userID", ctrl.GetSearchHistory())
// 	group.DELETE("/clear/:userID", ctrl.ClearSearchHistory())
// }
