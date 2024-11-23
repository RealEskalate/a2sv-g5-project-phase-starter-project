package blog_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bc *BlogController) Removelike(c *gin.Context) {
	blogid := c.Param("id")
	userID := c.GetString("user_id")
	isAdmin := c.GetBool("is_admin")

	blogID, err := primitive.ObjectIDFromHex(blogid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}
	userId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	err = bc.usecase.RemoveLike(context.Background(), userId, blogID, isAdmin)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like removed successfully"})
}
