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

// GenerateTextWithTags generates text based on the provided tags.
// It takes a gin.Context object and a slice of domain.Tag as input.
// The function binds the JSON request to the 'request' variable.
// If there is an error in binding the JSON, it returns a JSON response with the error message.
// Otherwise, it calls the AiUsecase.GenerateTextWithTags method passing the context and the request.
// If there is an error in generating the text, it returns a JSON response with the error message.
// Finally, it constructs a JSON response with the generated text and returns it.
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

// GenerateTextWithPrompt generates text based on the given prompt.
// It takes a JSON object containing a prompt as input and returns the generated text as a response.
// If there is an error in binding the JSON object or generating the text, it returns an error message.
// The generated text is returned in the "result" field of the JSON response.
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

// GenerateSuggestions generates suggestions based on the provided text content.
// It takes a gin.Context object and a map[string]string containing the text content as input.
// It returns the generated suggestions as a string and an error, if any.
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
