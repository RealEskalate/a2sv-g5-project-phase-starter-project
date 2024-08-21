package blogcontroller

import (
	"log"
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
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if prompts.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	content := "Generate content for this title :" + prompts.Title + " by considering this tags :" + strings.Join(prompts.Tags, ", ") + ". If the content contains any inappropriate content, please remove it, and state the reason for the removal."
	result, err := l.BlogUsecase.GenerateAiContent(content)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"content": result})
}
