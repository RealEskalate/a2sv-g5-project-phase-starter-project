package blog_controller

import (
	"blog-api/domain/blog"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bc *BlogController) CreateBlog(c *gin.Context) {
	var blog blog.Blog

	// Bind the incoming JSON request to the blog struct
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the AuthorID (assuming it's extracted from JWT or another source)
	authorID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Cast the authorID to primitive.ObjectID
	if oid, ok := authorID.(string); ok {
		objID, err := primitive.ObjectIDFromHex(oid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
			return
		}
		blog.AuthorID = objID
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID format"})
		return
	}

	// Create the blog post using the usecase
	createdBlog, err := bc.usecase.CreateBlog(context.Background(), &blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created blog post with StatusCreated
	c.JSON(http.StatusCreated, createdBlog)
}
