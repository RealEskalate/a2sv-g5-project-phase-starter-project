package blog_controller

import (
	"blog-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *BlogController) FilterBlog(c *gin.Context) {
	var filters domain.FilterRequest
	var posts []domain.Blog
	if err := c.ShouldBind(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
