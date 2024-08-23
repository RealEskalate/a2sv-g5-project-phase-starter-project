package blog_controller

import (
	"blog-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bc *BlogController) CreateBlog(c *gin.Context) {
	var blog domain.BlogRequest

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authorID := c.GetString("user_id")

	ID, err := primitive.ObjectIDFromHex(authorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	createdBlog, err := bc.usecase.CreateBlog(c, &blog, ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdBlog)
}
