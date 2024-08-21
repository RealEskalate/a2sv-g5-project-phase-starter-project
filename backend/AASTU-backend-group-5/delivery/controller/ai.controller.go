package controller

import (
	"net/http"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/gin-gonic/gin"
)

type AI_controller struct {
	Ai_func domain.AI_interface
}

func (ai *AI_controller) GenerateBlog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input domain.AI_Input
		var output domain.AI_Output

		err := ctx.BindJSON(&input)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		generated_content, err := ai.Ai_func.GenerateContentFromGemini(input.Title, input.Description)
		if err != nil {
			ctx.IndentedJSON(http.StatusRequestTimeout, gin.H{"error": err.Error()})
			return
		}
		output.Title = input.Title
		output.Content = generated_content
		ctx.IndentedJSON(http.StatusOK, gin.H{"data": output})
	}
}
