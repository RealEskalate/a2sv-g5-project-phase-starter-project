package controllers

import (
	"strings"

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
	keywords := strings.Split(c.Query("keywords"), " ")
	tone := c.DefaultQuery("tone", "neutral")
	audience := c.DefaultQuery("audience", "general")
	if len(keywords) == 0 {
		c.JSON(http.StatusBadRequest, "keywords are required")
		return
	}

	resp, err := controller.usecase.GenerateBlog(keywords, tone, audience)
	if err != nil {
		c.JSON(err.StatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"generated blog": resp})
}

func (controller *blogAssistantController) EnhanceBlog(c *gin.Context) {
	content := c.Query("content")
	command := c.Query("command")
	if content == "" {
		c.JSON(http.StatusBadRequest, "blog content is required")
	}
	resp, err := controller.usecase.EnhanceBlog(content, command)
	if err != nil {
		c.JSON(err.StatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"enhanced blog": resp})
}

func (controller *blogAssistantController) SuggestBlog(c *gin.Context) {
	niche := c.Query("industry")

	resp, err := controller.usecase.SuggestBlog(niche)
	if err != nil {
		c.JSON(err.StatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"suggested topics": resp})
}
