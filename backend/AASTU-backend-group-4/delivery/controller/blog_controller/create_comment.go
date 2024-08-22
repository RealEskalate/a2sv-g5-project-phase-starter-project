package blog_controller

import (
	"blog-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bc *BlogController) CreateComment(c *gin.Context) {
	var comment domain.CommentRequest

	// Bind the incoming JSON request to the blog struct
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the AuthorID (assuming it's extracted from JWT or another source)
	authorID := c.GetString("user_id")

	ID, err := primitive.ObjectIDFromHex(authorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	comment.UserID = ID

	// Create the blog post using the usecase
	createdBlog, err := bc.usecase.CreateComment(c, &comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created blog post with StatusCreated
	c.JSON(http.StatusCreated, gin.H{"comment": createdBlog})
}
