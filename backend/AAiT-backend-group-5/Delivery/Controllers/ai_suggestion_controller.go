package controllers

import (
	"net/http"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
)

type ContentSuggestionController struct {
	AI_Service interfaces.ContentSuggester
}

func NewContentSuggestionController(AI_Service interfaces.ContentSuggester) *ContentSuggestionController {
	return &ContentSuggestionController{
		AI_Service: AI_Service,
	}
}

func (suggestionController *ContentSuggestionController) HandleSuggestion(ctx *gin.Context) {
	query := ctx.Query("query")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	suggestion, err := suggestionController.AI_Service.SuggestContent(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"suggestion": suggestion})
}
