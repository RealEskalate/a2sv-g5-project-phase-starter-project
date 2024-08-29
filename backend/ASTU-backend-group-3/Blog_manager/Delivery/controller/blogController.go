package controller

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/Usecases"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogController struct {
	blogUsecase Usecases.BlogUsecase
	userUsecase Usecases.UserUsecase
}

func NewBlogController(bu Usecases.BlogUsecase, uu Usecases.UserUsecase) *BlogController {
	return &BlogController{
		blogUsecase: bu,
		userUsecase: uu,
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
		Comments:  []Domain.Comment{},
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
	userRole, _ := c.Get("role") // Extract the role from the header

	// Get the current user's username from the context
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}
	fmt.Println("Username from context:", username)
	// Check if the user is an admin or the author of the blog
	isAdmin := userRole == "admin"
	if !isAdmin {
		// Find the blog by ID to check if the requesting user is the author
		blog, err := h.blogUsecase.FindByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
			return
		}
		if blog.Author != username {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this blog"})
			return
		}
	}

	// Delete the blog
	err := h.blogUsecase.DeleteBlogByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
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

func (h *BlogController) UpdateBlog(c *gin.Context) {
	blogID := c.Param("id")

	var input Domain.UpdateBlogInput // Use the UpdateBlogInput struct

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract the author ID from the context (assuming it's set by authentication middleware)
	author := c.GetString("username")
	if author == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	// Call the use case to update the blog
	updatedBlog, err := h.blogUsecase.UpdateBlog(blogID, input, author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedBlog})
}

func (h *BlogController) IncrementViewCount(c *gin.Context) {
	blogID := c.Param("id")

	err := h.blogUsecase.IncrementViewCount(blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to increment view count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "View count updated"})
}

func (h *BlogController) ToggleLike(c *gin.Context) {
	blogID := c.Param("id")
	username := c.GetString("username")

	err := h.blogUsecase.ToggleLike(blogID, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to toggle like"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like toggled"})
}

func (h *BlogController) ToggleDislike(c *gin.Context) {
	blogID := c.Param("id")
	username := c.GetString("username")

	err := h.blogUsecase.ToggleDislike(blogID, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to toggle dislike"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dislike toggled"})
}

func (h *BlogController) AddComment(c *gin.Context) {
	blogID := c.Param("id")

	var input struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment := Domain.Comment{
		Id:        primitive.NewObjectID().Hex(),
		Content:   input.Content,
		PostID:    blogID,
		UserID:    c.GetString("username"),
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	err := h.blogUsecase.AddComment(blogID, comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment added"})
}

func (c *BlogController) FilterBlogs(ctx *gin.Context) {
	var tags []string
	if ctx.Query("tags") != "" {
		tags = strings.Split(ctx.Query("tags"), ",")
	}

	// Define the date format to include only year, month, and day
	const layout = "2006-01-02"

	startDateStr := ctx.Query("startDate")
	endDateStr := ctx.Query("endDate")

	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse(layout, startDateStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid startDate format. Use YYYY-MM-DD"})
			return
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse(layout, endDateStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid endDate format. Use YYYY-MM-DD"})
			return
		}
	}

	fmt.Print(startDate, endDate)
	sortBy := ctx.Query("sortBy")

	blogs, err := c.blogUsecase.FilterBlogs(tags, startDate, endDate, sortBy)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}
