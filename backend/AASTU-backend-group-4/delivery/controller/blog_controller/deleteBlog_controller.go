package blog_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bc *BlogController) DeleteBlog(c *gin.Context) {
	blogIDParam := c.Param("id")
	userID := c.MustGet("user_id").(primitive.ObjectID)
	isAdmin := c.MustGet("is_admin").(bool)

	blogID, err := primitive.ObjectIDFromHex(blogIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	err = bc.usecase.DeleteBlog(c.Request.Context(), userID, blogID, isAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}
