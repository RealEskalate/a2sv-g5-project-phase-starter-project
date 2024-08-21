package controllers

import (
	domain "aait-backend-group4/Domain"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogController struct {
	BlogUsecase domain.BlogUsecase
}

// CreateBlog handles the creation of a new blog
func (bc *BlogController) CreateBlog(c *gin.Context) {
	var blog domain.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog.ID = primitive.NewObjectID()
	blog.Popularity = 0
	blog.Feedbacks = domain.Feedback{}
	blog.Created_At = time.Now()
	blog.Updated_At = time.Now()

	if err := bc.BlogUsecase.CreateBlog(c, &blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blog)
}

// UpdateBlog handles updating an existing blog
func (bc *BlogController) UpdateBlog(c *gin.Context) {
	blogID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}
	userID := c.Request.Header.Get("userID")
	var blog domain.BlogUpdate

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTime := time.Now()
	blog.Updated_At = &updatedTime
	if err := bc.BlogUsecase.UpdateBlog(c, objectID, blog, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully"})
}

// DeleteBlog handles deleting a blog
func (bc *BlogController) DeleteBlog(c *gin.Context) {
	blogID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}
	userID := c.Request.Header.Get("userID")

	if err := bc.BlogUsecase.DeleteBlog(c, objectID, userID); err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		} else if err.Error() == "unauthorized" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized to delete this blog post"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

// FetchAll handles fetching all blogs with pagination
func (bc *BlogController) FetchAll(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	blogs, err := bc.BlogUsecase.FetchAll(c, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// FetchByPageAndPopularity handles fetching blogs by page and popularity
func (bc *BlogController) FetchByPageAndPopularity(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	blogs, err := bc.BlogUsecase.FetchByPageAndPopularity(c, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// FetchByBlogAuthor handles fetching blogs by author ID
func (bc *BlogController) FetchByBlogAuthor(c *gin.Context) {
	authorID := c.Param("author_id")
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	blogs, err := bc.BlogUsecase.FetchByBlogAuthor(c, authorID, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// FetchByTags handles fetching blogs by tags
func (bc *BlogController) FetchByTags(c *gin.Context) {
	var tags []domain.Tag
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	if err := c.ShouldBindJSON(&tags); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blogs, err := bc.BlogUsecase.FetchByTags(c, tags, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// FetchByBlogTitle handles fetching blogs by title
func (bc *BlogController) FetchByBlogTitle(c *gin.Context) {
	title := c.Query("title")
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	blogs, err := bc.BlogUsecase.FetchByBlogTitle(c, title, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// AddComment handles adding a comment to a blog
func (bc *BlogController) AddComment(c *gin.Context) {
	var comment domain.Comment
	userID := c.Request.Header.Get("userID")

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.BlogUsecase.AddComment(c, userID, comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
}

// UpdateComment handles updating a comment in a blog
func (bc *BlogController) UpdateComment(c *gin.Context) {
	blogID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}
	userID := c.Request.Header.Get("userID")
	var updatedComment domain.Comment

	if err := c.ShouldBindJSON(&updatedComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.BlogUsecase.UpdateComment(c, objectID, userID, updatedComment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

// RemoveComment handles removing a comment from a blog
func (bc *BlogController) RemoveComment(c *gin.Context) {
	blogID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}
	userID := c.Request.Header.Get("userID")

	if err := bc.BlogUsecase.RemoveComment(c, objectID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment removed successfully"})
}

// SearchBlogs handles the search of blogs based on filters
func (bc *BlogController) SearchBlogs(c *gin.Context) {
	var filter domain.Filter
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blogs, err := bc.BlogUsecase.SearchBlogs(c, filter, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// FetchByBlogID handles fetching a single blog by its ID
func (bc *BlogController) FetchByBlogID(c *gin.Context) {
	blogID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	blog, err := bc.BlogUsecase.FetchByBlogID(c, objectID.Hex())
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred"})
		}
		return
	}

	c.JSON(http.StatusOK, blog)
}
