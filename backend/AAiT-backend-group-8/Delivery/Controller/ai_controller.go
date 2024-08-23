package controller

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"AAiT-backend-group-8/Domain"
)


func (c *Controller) GenerateBlog(ctx *gin.Context) {
	var req Domain.BlogRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := c.AiUseCase.GenerateBlogContent(req.UserInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate content"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) SuggestImprovements(ctx *gin.Context) {
	var req Domain.BlogSuggestionRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := c.AiUseCase.SuggestImprovements(req.Title, req.Body, req.Tags)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to suggest improvements"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
