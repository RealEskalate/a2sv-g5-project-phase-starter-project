package controllers

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AIController struct {
	AiUsecase domain.AiUsecase
	Env       *bootstrap.Env
}

func (ac *AIController) GenerateTextWithTags(c *gin.Context) {
	var request []domain.Tag

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	text, err := ac.AiUsecase.GenerateTextWithTags(c, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg := map[string]string{
		"result": text,
	}

	c.JSON(http.StatusOK, msg)
}

func (ac *AIController) GenerateTextWithPrompt(c *gin.Context) {
	var prompt map[string]string

	err := c.ShouldBindJSON(&prompt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	text, err := ac.AiUsecase.GenerateTextWithPrompt(c, prompt["prompt"])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg := map[string]string{
		"result": text,
	}

	c.JSON(http.StatusOK, msg)

}

func (ac *AIController) GenerateSuggestions(c *gin.Context) {
	var textContent map[string]string

	err := c.ShouldBindJSON(&textContent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	text, err := ac.AiUsecase.GenerateSuggestions(c, textContent["textContent"])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg := map[string]string{
		"suggestions": text,
	}

	c.JSON(http.StatusOK, msg)

}
