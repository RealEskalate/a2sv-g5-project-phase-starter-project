package controllers

import (
	domain "aait-backend-group4/Domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogController struct {
	BlogUsecase domain.BlogUsecase
}

func (bc *BlogController) CreateBlog(c *gin.Context) {
	var blog domain.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog.ID = primitive.NewObjectID()
	blog.Created_At = time.Now()
	blog.Updated_At = time.Now()

	if err := bc.BlogUsecase.CreateBlog(c, &blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blog)
}

func (bc *BlogController) UpdateBlog(c *gin.Context) {
	blogID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}
	userID := c.Request.Header.Get("userID")
	var blog domain.BlogUpdate // Adjusted to match usecase

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog.Updated_At = time.Now()
	if err := bc.BlogUsecase.UpdateBlog(c, objectID, blog, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully"})
}

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

func (bc *BlogController) FetchAll(c *gin.Context) {
	blogs, err := bc.BlogUsecase.FetchAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func (bc *BlogController) FetchByBlogAuthor(c *gin.Context) {
	authorID := c.Param("author_id")
	blogs, err := bc.BlogUsecase.FetchByBlogAuthor(c, authorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func (bc *BlogController) FetchByBlogTitle(c *gin.Context) {
	title := c.Param("title")
	blogs, err := bc.BlogUsecase.FetchByBlogTitle(c, title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func (bc *BlogController) UpdateFeedback(c *gin.Context) {
	blogID := c.Param("id")
	var feedback domain.Feedback

	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateFunc := func(fb *domain.Feedback) error {
		*fb = feedback
		return nil
	}

	if err := bc.BlogUsecase.UpdateFeedback(c, blogID, updateFunc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Feedback updated successfully"})
}

func (bc *BlogController) SearchBlogs(c *gin.Context) {
	var filter domain.Filter

	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blogs, err := bc.BlogUsecase.SearchBlogs(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}
