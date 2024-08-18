package controller

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/Usecases"
	"net/http"
	"strconv"
	"time"

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
