package controller

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/Usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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

func (h *BlogController) DeleteBlogByID(c *gin.Context) {
	id := c.Param("id")
	err := h.blogUsecase.DeleteBlogByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog not post not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog post deleted successfully"})
}

func (h *BlogController) SearchBlogs(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	tags := c.QueryArray("tags") // Assuming tags are provided as a query array like ?tags=tag1&tags=tag2

	blogs, err := h.blogUsecase.SearchBlogs(title, author, tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": blogs})
}
