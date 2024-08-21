package controllers

import (
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/gin-gonic/gin"

	"net/http"
)

type blogAssistantController struct {
	usecase domain.BlogAssistantUseCase
}

func NewBlogAssistantController(blogAssistantUsecase domain.BlogAssistantUseCase) domain.BlogAssistantController {
	return &blogAssistantController{usecase: blogAssistantUsecase}
}

func (controller *blogAssistantController) GenerateBlog(c *gin.Context) {
	var request struct {
		Keywords []string `json:"keywords" binding:"required"`
		Tone     string   `json:"tone"`
		Audience string   `json:"audience"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(request.Keywords) == 0 {
		c.JSON(http.StatusBadRequest, "keywords are required")
		return
	}

	resp, err := controller.usecase.GenerateBlog(request.Keywords, request.Tone, request.Audience)
	if err != nil {
		c.JSON(err.StatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (controller *blogAssistantController) EnhanceBlog(c *gin.Context) {
	var request struct {
		Content string `json:"content" binding:"required"`
		Command string `json:"command"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := controller.usecase.EnhanceBlog(request.Content, request.Command)
	if err != nil {
		c.JSON(err.StatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (controller *blogAssistantController) SuggestBlog(c *gin.Context) {
	var niche struct {
		Niche string `json:"industry" binding:"required"`
	}
	if err := c.ShouldBindJSON(&niche); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := controller.usecase.SuggestBlog(niche.Niche)
	if err != nil {
		c.JSON(err.StatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}