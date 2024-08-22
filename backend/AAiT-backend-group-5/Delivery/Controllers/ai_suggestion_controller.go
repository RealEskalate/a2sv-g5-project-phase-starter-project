package controllers

import (
	"net/http"
	"strings"

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

	blog_title := strings.TrimSpace(suggestion[0])
	blog_content := strings.TrimSpace(suggestion[1])

	blog_tag_list := strings.TrimSpace(suggestion[2])
	blog_tags := strings.Split(blog_tag_list, "&")

	ctx.JSON(http.StatusOK, gin.H{"title": blog_title, "content": blog_content, "tags": blog_tags})
}

func (suggestionController *ContentSuggestionController) HandleContentImprovement(ctx *gin.Context) {
	blogID := ctx.Param("id")

	suggestion, err := suggestionController.AISuggestionUsecase.ImproveBlogContent(blogID)
	if err != nil {
		ctx.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	blog_title := strings.TrimSpace(suggestion[0])
	blog_content := strings.TrimSpace(suggestion[1])

	blog_tag_list := strings.TrimSpace(suggestion[2])
	blog_tags := strings.Split(blog_tag_list, "&")

	ctx.JSON(http.StatusOK, gin.H{"title": blog_title, "content": blog_content, "tags": blog_tags})
}
