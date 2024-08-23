package controllers

import (
	"blogapp/Config"
	"blogapp/Dtos"
	groqservice "blogapp/Infrastructure/groq_services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type chatController struct {
	groqservice *groqservice.GroqAI
}

func NewChatController() *chatController {

	groqAI := groqservice.NewGroqAI(Config.GROQ_API_KEY)
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

	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

func (cc *chatController) GetChatCompletionByTags(c *gin.Context) {

	var json struct {
		Tags []string `json:"tags"`
	}

	if err := c.ShouldBindJSON(&json); err != nil || len(json.Tags) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	response, err := cc.groqservice.GetChatCompletionByTags(json.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

func (cc *chatController) GetChatCompletionEnhancements(c *gin.Context) {

	var json struct {
		Prompt string `json:"prompt"`
	}

	if err := c.ShouldBindJSON(&json); err != nil || json.Prompt == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	response, err := cc.groqservice.GetChatCompletionEnhancements(json.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

func (cc *chatController) HandleGeneratePostRequest(c *gin.Context) {
	var dto Dtos.PostDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	response, err := cc.groqservice.GenerateBlog(dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}
