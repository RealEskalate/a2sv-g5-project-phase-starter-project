package controller

import (
	"backend-starter-project/domain/dto"
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

	var response dto.Response


	if err := c.ShouldBindJSON(&req); err != nil {
		response.Success = false
		response.Error = "Invalid request payload"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	contentSuggestion, err := acc.aiContentService.GenerateContentSuggestions( req.Keywords)
	if err != nil {
		response.Success = false
		response.Error = "Error generating content suggestions"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Success = true
	response.Data = gin.H{
		"content": contentSuggestion,
	}
	c.JSON(http.StatusOK, response)
}

func (acc *AIContentController) SuggestContentImprovements(c *gin.Context) {
	var req struct {
		BlogPostID string `json:"blogPostId"`
		Instruction string `json:"instruction"`
	}

	var response dto.Response

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Success = false
		response.Error = "Invalid request payload"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	contentSuggestion, err := acc.aiContentService.SuggestContentImprovements(req.BlogPostID, req.Instruction)
	if err != nil {
		response.Success = false
		response.Error = "Error suggesting content improvements"
		c.JSON(http.StatusBadRequest, response)
		return
	}


	response.Success = true
	response.Data = gin.H{
		"suggestion result": contentSuggestion,
	}
	c.JSON(http.StatusOK, response)
}
