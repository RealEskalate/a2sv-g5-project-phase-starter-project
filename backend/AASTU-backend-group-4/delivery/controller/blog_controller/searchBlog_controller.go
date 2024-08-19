package blog_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (bc *BlogController) SearchBlog(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")

	blogs, err := bc.usecase.SearchBlog(c.Request.Context(), title, author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search blogs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blogs": blogs})
}
