package blog_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (bc *BlogController) GenerateContent(c *gin.Context) {
	prompt := c.Query("prompt")
	if prompt == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Prompt is required"})
		return
	}

	content, err := bc.usecase.GenerateAIContent(context.Background(), prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"content": content})
}
