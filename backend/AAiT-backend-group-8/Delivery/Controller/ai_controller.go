package controller

import (
	"AAiT-backend-group-8/Domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller *Controller) GenerateBlog(ctx *gin.Context) {
	var req Domain.BlogRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := controller.AiUseCase.GenerateBlogContent(req.UserInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate content"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (controller *Controller) SuggestImprovements(ctx *gin.Context) {
	var req Domain.BlogSuggestionRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := controller.AiUseCase.SuggestImprovements(req.Title, req.Body, req.Tags)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to suggest improvements"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
