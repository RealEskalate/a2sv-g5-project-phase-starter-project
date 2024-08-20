package controllers

import (
	"net/http"

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
	// Get the user prompt from the request body
	var requestBody struct {
		Prompt string `json:"prompt"`
	}
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Generate content using the use case
	response, err := c.aiUseCase.GenerateContent(requestBody.Prompt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the response
	ctx.JSON(http.StatusOK, gin.H{"content": response})
}
