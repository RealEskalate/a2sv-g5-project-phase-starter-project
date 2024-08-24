package controller

import (
	"net/http"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/gin-gonic/gin"
)

type BookmarkController struct {
	BookmarkUsecase domain.BookmarkUseCaseInterface
}

func NewBookmarkController(usecase domain.BookmarkUseCaseInterface) *BookmarkController {
	return &BookmarkController{
		BookmarkUsecase: usecase,
	}
}

func (bc *BookmarkController) AddBookmark() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userID")
		blogID := c.Param("blogID")

		if err := bc.BookmarkUsecase.AddBookmark(userID, blogID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add bookmark"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Bookmark added successfully"})
	}
}

func (bc *BookmarkController) RemoveBookmark() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userID")
		blogID := c.Param("blogID")

		if err := bc.BookmarkUsecase.RemoveBookmark(userID, blogID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove bookmark"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Bookmark removed successfully"})
	}
}

func (bc *BookmarkController) GetUserBookmarks() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userID")

		bookmarks, err := bc.BookmarkUsecase.GetUserBookmarks(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve bookmarks"})
			return
		}

		c.JSON(http.StatusOK, bookmarks)
	}
}
