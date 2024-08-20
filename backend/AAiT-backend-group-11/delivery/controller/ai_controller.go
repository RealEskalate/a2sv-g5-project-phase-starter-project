package controller

import (
	"backend-starter-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AIContentController struct {
	aiContentService service.AIContentServiceInterface
}

func NewAIContentController(aiContentService service.AIContentServiceInterface) *AIContentController {
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
		Instruction string `json:"instruction"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contentSuggestion, err := acc.aiContentService.SuggestContentImprovements(req.BlogPostID, req.Instruction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contentSuggestion)
}
