package controllers

import (
	domain "blogs/Domain"
	infrastructure "blogs/Infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AiController struct {
	Config *infrastructure.Config
	AiUsecase domain.AIUsecase
}

func (controller *AiController) Ask(c *gin.Context) {

	var request domain.AiRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	response := controller.AiUsecase.AskAI(c, request)
	
	HandleResponse(c, response)

}