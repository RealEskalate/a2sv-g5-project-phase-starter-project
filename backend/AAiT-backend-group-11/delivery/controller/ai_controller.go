package controllers

import (
	"backend-starter-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AIContentController struct {
	aiContentService service.AIContentService
}

func NewAIContentController(aiContentService service.AIContentService) *AIContentController {
	return &AIContentController{
		aiContentService: aiContentService,

	}}

func (acc *AIContentController) GenerateContentSuggestions(c *gin.Context) {
	var req struct {
		Keywords []string `json:"keywords"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contentSuggestion, err := acc.aiContentService.GenerateContentSuggestions( req.Keywords)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contentSuggestion)
}

func (acc *AIContentController) SuggestContentImprovements(c *gin.Context) {
	var req struct {
		BlogPostID string `json:"blogPostId"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contentSuggestion, err := acc.aiContentService.SuggestContentImprovements(req.BlogPostID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contentSuggestion)
}
