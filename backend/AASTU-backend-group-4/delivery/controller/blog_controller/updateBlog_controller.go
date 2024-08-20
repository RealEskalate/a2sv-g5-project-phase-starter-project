package blog_controller

import (
	"blog-api/domain"
	"net/http"
	"time"

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

	existingBlog, err := bc.usecase.GetBlogByID(c.Request.Context(), blogID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		return
	}
	if existingBlog.AuthorID != authorID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the author of this blog post"})
		return
	}

	var updatedBlog domain.Blog
	if err := c.ShouldBindJSON(&updatedBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	updatedBlog.UpdatedAt = time.Now()

	err = bc.usecase.UpdateBlog(c.Request.Context(), blogID, &updatedBlog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog post"})
		return
	}

	// Return the updated blog data
	c.JSON(http.StatusOK, gin.H{
		"message": "Blog updated successfully",
		"blog":    updatedBlog,
	})
}
