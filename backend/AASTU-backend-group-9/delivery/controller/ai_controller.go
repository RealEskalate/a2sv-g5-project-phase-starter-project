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
    keywords := c.Query("keywords")
    if keywords == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Keywords are required"})
        return
    }

    content, err := ctrl.AIUsecase.GenerateBlogContent(c.Request.Context(), keywords)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"content": content})
}