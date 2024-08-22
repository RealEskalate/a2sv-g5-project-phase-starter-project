package blog_controller

import (
	"blog-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bc *BlogController) AddLike(c *gin.Context) {
	var like domain.LikeRequest

	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authorID := c.GetString("user_id")

	ID, err := primitive.ObjectIDFromHex(authorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	like.UserID = ID

	// Create the blog post using the usecase
	err = bc.usecase.AddLike(c, &like)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created blog post with StatusCreated
	c.JSON(http.StatusCreated, gin.H{"message": "liked successfuly"})
}
