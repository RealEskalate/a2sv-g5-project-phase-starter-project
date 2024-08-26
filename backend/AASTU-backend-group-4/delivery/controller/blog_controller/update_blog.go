package blog_controller

import (
	"blog-api/domain"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bc *BlogController) UpdateBlog(c *gin.Context) {
	blogIDParam := c.Param("id")
	authorID := c.MustGet("user_id").(primitive.ObjectID)

	blogID, err := primitive.ObjectIDFromHex(blogIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	var updatedBlog domain.UpdateBlogRequest
	if err := c.ShouldBindJSON(&updatedBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err = bc.usecase.UpdateBlog(context.Background(), blogID, updatedBlog, authorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Blog updated successfully",
		"blog":    updatedBlog,
	})
}
