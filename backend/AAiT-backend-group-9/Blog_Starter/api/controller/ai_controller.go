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

	preText := "Generate a blog post in plain text. Be sure the response is plain text response. The topic is: "
	aiContent, err := ac.aiUseCase.GenerateAIContent(ac.ctx, content, preText)
	if err != nil {
		if err.Error() == "content is empty" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, aiContent)
}

func (ac *AIController) SuggestAIContent(c *gin.Context) {
	// implementation
	var content string
	err := c.ShouldBindJSON(&content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	preText := "refine the following blog post and return it in plain text.Be sure the response is plain text response. the blog is:"
	aiContent, err := ac.aiUseCase.GenerateAIContent(ac.ctx, content, preText)
	if err != nil {
		if err.Error() == "content is empty" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, aiContent)
}
