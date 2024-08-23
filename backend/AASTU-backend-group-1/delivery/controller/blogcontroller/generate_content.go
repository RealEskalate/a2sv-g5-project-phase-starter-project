package blogcontroller

import (
	"blogs/domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (l *BlogController) GenerateContent(ctx *gin.Context) {
	var prompts struct {
		Title string   `json:"title"`
		Tags  []string `json:"tags"`
	}

	if err := ctx.ShouldBindJSON(&prompts); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	if prompts.Title == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "title is required",
		})
		return
	}

	content := "Generate content for this title :" + prompts.Title + " by considering this tags :" + strings.Join(prompts.Tags, ", ") + ". If the content contains any inappropriate content, please remove it, and state the reason for the removal."
	result, err := l.BlogUsecase.GenerateAiContent(content)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    result,
	})
}
