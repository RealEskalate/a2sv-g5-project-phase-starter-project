package controllers

import (
	"net/http"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
)

type ContentSuggestionController struct {
	AISuggestionUsecase interfaces.AIContentSuggestionUsecase
}

func NewContentSuggestionController(suggestionUsecase interfaces.AIContentSuggestionUsecase) *ContentSuggestionController {
	return &ContentSuggestionController{
		AISuggestionUsecase: suggestionUsecase,
	}
}

func (suggestionController *ContentSuggestionController) HandleSuggestion(ctx *gin.Context) {
	query := ctx.Query("query")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	suggestion, err := suggestionController.AISuggestionUsecase.SuggestContent(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"suggestion": suggestion})
}
