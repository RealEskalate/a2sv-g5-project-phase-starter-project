package controller

import (
    "blog/domain"
    "net/http"
    
    "github.com/gin-gonic/gin"
)

type AIController struct {
    AIUsecase domain.AIUsecase
}

func (ctrl *AIController) GenerateContent(c *gin.Context) {
    var request struct {
        Prompt string `json:"keywords" binding:"required"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    content, err := ctrl.AIUsecase.GenerateBlogContent(c.Request.Context(), request.Prompt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"content": content})
}