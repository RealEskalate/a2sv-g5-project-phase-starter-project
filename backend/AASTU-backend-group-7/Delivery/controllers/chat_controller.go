package controllers

import (
	groqservice "blogapp/Infrastructure/groq_services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type chatController struct {
	groqservice *groqservice.GroqAI
}

func NewChatController() *chatController {

	groqAI := groqservice.NewGroqAI("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")
	return &chatController{
		groqservice: groqAI,
	}
}

// complete chat
func (cc *chatController) GetChatCompletion(c *gin.Context) {

	var json struct {
		Prompt string `json:"prompt"`
	}

	if err := c.ShouldBindJSON(&json); err != nil || json.Prompt == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	response, err := cc.groqservice.GetChatCompletion(json.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}
