package blog_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bc *BlogController) AddLike(c *gin.Context) {
	blogIDParam := c.Param("id")
	userID := c.MustGet("user_id").(primitive.ObjectID)

	ID, err := primitive.ObjectIDFromHex(blogIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = bc.usecase.AddLike(c, userID, ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "liked successfuly"})
}
