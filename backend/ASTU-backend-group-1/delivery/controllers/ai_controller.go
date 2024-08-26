package controllers

import (
	"astu-backend-g1/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AIController struct {
	model infrastructure.AIModel
}

func NewAIController(model infrastructure.AIModel) *AIController {
	return &AIController{model: model}
}
func (c *AIController) RecommendTitle(ctx *gin.Context) {
	data := infrastructure.Data{}
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "invalid data format"})
		return
	}
	resp, err := c.model.Recommend(data, "title")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no response from the ai model"})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c *AIController) RecommendContent(ctx *gin.Context) {
	data := infrastructure.Data{}
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "invalid data format"})
		return
	}
	resp, err := c.model.Recommend(data, "content")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no response from the ai model"})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
func (c *AIController) Recommendtags(ctx *gin.Context) {
	data := infrastructure.Data{}
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "invalid data format"})
		return
	}
	resp, err := c.model.Recommend(data, "tags")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no response from the ai model"})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
