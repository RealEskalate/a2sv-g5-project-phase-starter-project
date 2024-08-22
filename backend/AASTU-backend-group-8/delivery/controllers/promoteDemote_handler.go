package controllers

import (
	"meleket/domain"
	// "meleket/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PromoteDemoteController struct {
	userUsecase domain.UserUsecaseInterface
}

func NewPromoteDemoteController(userUsecase domain.UserUsecaseInterface) *PromoteDemoteController {
	return &PromoteDemoteController{userUsecase: userUsecase}
}

func (pdc *PromoteDemoteController) PromoteToAdmin(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pdc.userUsecase.PromoteToAdmin(req.Username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}

func (pdc *PromoteDemoteController) DemoteToUser(c *gin.Context) {
	demoterRole := c.MustGet("role")
	if demoterRole != "root" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Only root can demote admins"})
		return
	}

	var req struct {
		Username string `json:"username" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pdc.userUsecase.DemoteToUser(req.Username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User demoted to regular user"})
}
