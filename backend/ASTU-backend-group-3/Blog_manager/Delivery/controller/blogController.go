package controller

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/Usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	blogUsecase Usecases.BlogUsecase
}

func NewBlogController(bu Usecases.BlogUsecase) *BlogController {
	return &BlogController{
		blogUsecase: bu,
	}
}

func (h *BlogController) CreateBlog(c *gin.Context) {
	var input Domain.Blog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := c.GetString("username") // Extracted from the context

	// Set the user ID in the blog details
	input.Author = username

	createdBlog, err := h.blogUsecase.CreateBlog(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": createdBlog})
}
