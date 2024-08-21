package controllers

import (
	"meleket/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AIBlog struct {
	Title string `json:"title" binding:"required"`
	Tags  []string `json:"tags"`
}

type AIHandler struct {
	aiUsecase domain.AIUsecaseInterface
}

func NewAIHandler(ai domain.AIUsecaseInterface) *AIHandler {
	return &AIHandler{
		aiUsecase: ai,
	}
}

func (s *AIHandler) GenerateBlogWithAI(c *gin.Context) {
	aiInput := domain.AIBlog
	if err := c.ShouldBindJSON(&aiInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate blog content using AI
	generatedContent, err := s.aiUsecase.GenerateBlogContent(aiInput.Title, aiInput.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blog": generatedContent})
}
