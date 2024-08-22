package controllers

import (
	"AAIT-backend-group-3/internal/usecases"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AiHelperControllerInterface interface {
	GenerateBlog(c *gin.Context)
	EnhanceBlog(c *gin.Context)
	GenerateSummary(c *gin.Context)
	GenerateTags(c *gin.Context)
}

type AiHelperController struct {
	ai_helper_usecase usecases.AiHelperUsecaseInterface
}

func NewAiHelperController(u usecases.AiHelperUsecaseInterface) AiHelperControllerInterface {
	return &AiHelperController{
		ai_helper_usecase: u,
	}
}

func (ahc *AiHelperController) GenerateBlog(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if _, ok := req["description"]; !ok {
		c.JSON(400, gin.H{"error": "description is required"})
		return
	}

	generatedBlog, err := ahc.ai_helper_usecase.GenerateBlog(req["description"].(string))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate blog"})
		return
	}

	c.JSON(200, gin.H{"blog": generatedBlog})
}

func (ahc *AiHelperController) EnhanceBlog(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if _, ok := req["content"]; !ok {
		c.JSON(400, gin.H{"error": "content is required"})
		return
	}

	generatedBlog, err := ahc.ai_helper_usecase.EnhanceBlog(req["content"].(string))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to enhance blog"})
		return
	}

	c.JSON(200, gin.H{"blog": generatedBlog})
}

func (ahc *AiHelperController) GenerateSummary(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if _, ok := req["content"]; !ok {
		c.JSON(400, gin.H{"error": "content is required"})
		return
	}


	summary, err := ahc.ai_helper_usecase.GenerateSummary(req["content"].(string))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate summary"})
		return
	}

	c.JSON(200, gin.H{"summary": summary})
}

func (ahc *AiHelperController) GenerateTags(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if _, ok := req["content"]; !ok {
		c.JSON(400, gin.H{"error": "content is required"})
		return
	}
	tags, err := ahc.ai_helper_usecase.GenerateTags(req["content"].(string))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"tags": tags})
}


