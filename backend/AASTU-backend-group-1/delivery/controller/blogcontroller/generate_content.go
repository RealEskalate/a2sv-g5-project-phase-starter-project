package blogcontroller

import (
	"net/http"
	"strings"


	"github.com/gin-gonic/gin"
)

func (l *BlogController) GenerateContent(ctx *gin.Context) {
	var prompts struct {
		Title string `json:"title"`
		Tags []string `json:"tags"`
	}

	if err := ctx.ShouldBindJSON(&prompts); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}



	var content string
	
	content = "Generate content for this title :" + prompts.Title + " by considering this tags :" + strings.Join(prompts.Tags, ", ")

	result,err := l.BlogUsecase.GenerateAiContent(content)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, result)

}


