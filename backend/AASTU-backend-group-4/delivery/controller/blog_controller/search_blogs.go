package blog_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (bc *BlogController) SearchBlogs(c *gin.Context) {
	author := c.Query("author")
	title := c.Query("title")
	tags := c.QueryArray("tags")
	dateFrom := c.Query("dateFrom")
	dateTo := c.Query("dateTo")

	blogs, err := bc.usecase.SearchBlogs(context.Background(), author, title, tags, dateFrom, dateTo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}
