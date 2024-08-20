package controller

import (
	"Blog_Starter/usecase"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AIController struct {
	aiUseCase usecase.AIUseCase
	ctx       context.Context
}

func NewAIController(aiUseCase usecase.AIUseCase, ctx context.Context) *AIController {
	return &AIController{
		aiUseCase: aiUseCase,
		ctx:       ctx,
	}
}

func (ac *AIController) GenerateAIContent(c *gin.Context) {
	// implementation
	var content string
	err := c.ShouldBindJSON(&content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	aiContent, err := ac.aiUseCase.GenerateAIContent(ac.ctx, content)
	if err != nil {
		if err.Error() == "content is empty" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, aiContent)
}
