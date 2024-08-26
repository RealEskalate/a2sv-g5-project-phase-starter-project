package blog_controller

import (
	"blog-api/domain"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (bc *BlogController) GenerateContent(c *gin.Context) {
	var req domain.GenerateContentRequest

	// Bind JSON body to the request struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body or missing 'prompt' field"})
		return
	}

	// Use the prompt from the request struct
	content, err := bc.usecase.GenerateAIContent(context.Background(), req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"content": content})
}
