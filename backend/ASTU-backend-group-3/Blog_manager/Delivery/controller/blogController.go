package controller

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/Usecases"
	"net/http"
	"strconv"
	"time"

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
	var input struct {
		Title   string   `json:"title" binding:"required"`
		Content string   `json:"content" binding:"required"`
		Tags    []string `json:"tags" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the Blog struct with additional fields
	blog := &Domain.Blog{
		Title:     input.Title,
		Content:   input.Content,
		Tags:      input.Tags,
		Author:    c.GetString("username"), // Extracted from the context
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
		ViewCount: 0,
		Likes:     []string{},
		Dislikes:  []string{},
		Comments:  []string{},
	}

	createdBlog, err := h.blogUsecase.CreateBlog(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": createdBlog})
}

func (h *BlogController) RetrieveBlogs(c *gin.Context) {
	// Get query parameters for pagination and sorting
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	sortBy := c.DefaultQuery("sortBy", "date") // Default sort by date

	blogs, totalPosts, err := h.blogUsecase.RetrieveBlogs(page, pageSize, sortBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (totalPosts + int64(pageSize) - 1) / int64(pageSize) // Calculate total pages

	c.JSON(http.StatusOK, gin.H{
		"data":        blogs,
		"totalPages":  totalPages,
		"currentPage": page,
		"totalPosts":  totalPosts,
	})
}
