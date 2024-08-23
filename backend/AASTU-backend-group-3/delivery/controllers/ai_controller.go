package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"group3-blogApi/usecase"
)

type AIController struct {
	aiUseCase *usecase.AIUseCase
}

func NewAIController(aiUseCase *usecase.AIUseCase) *AIController {
	return &AIController{aiUseCase: aiUseCase}
}

func (c *AIController) GenerateContent(ctx *gin.Context) {
	// Define the request body structure
	var requestBody struct {
		Title     string   `json:"title"`
		WordCount int      `json:"word_count"`
		Tags      []string `json:"tags"`
		Keywords  []string `json:"keywords"`
	}
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Construct the prompt
	prompt := fmt.Sprintf("Create a blog post titled '%s' with a word count of %d. Include the following tags: %s. Use the following keywords: %s.",
		requestBody.Title,
		requestBody.WordCount,
		strings.Join(requestBody.Tags, ", "),
		strings.Join(requestBody.Keywords, ", "),
	)

	// Generate content using the use case
	response, err := c.aiUseCase.GenerateContent(prompt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the response
	ctx.JSON(http.StatusOK, gin.H{"content": response})
}
