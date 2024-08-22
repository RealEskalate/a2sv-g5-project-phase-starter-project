package blog_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (bc *BlogController) SearchBlog(c *gin.Context) {
	filter := make(map[string]string)
	filter["title"] = c.Query("title")
	filter["author"] = c.Query("author")
	if filter["title"] == "" && filter["author"] == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "at least one search criterion (title or author) must be provided"})
		return
	}

	blogs, err := bc.usecase.SearchBlog(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search blogs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blogs": blogs})
}
