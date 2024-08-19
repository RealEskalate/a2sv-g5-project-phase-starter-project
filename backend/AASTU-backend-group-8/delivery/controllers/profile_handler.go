package controllers

import (
	"fmt"
	"meleket/domain"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileUsecase domain.ProfileUsecase
}

func NewProfileHandler(p domain.ProfileUsecase) domain.ProfileHandler {
	return &ProfileHandler{
		profileUsecase: p,
	}
}

func (p *ProfileHandler) SaveProfile(c *gin.Context) {
	var profile domain.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := p.profileUsecase.SaveProfile(&profile)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Profile saved successfully",
	})
}

func (p *ProfileHandler) FindProfile(c *gin.Context) {
	userID := c.Param("user_id")
	fmt.Println(userID)
	profile, err := p.profileUsecase.FindProfile(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"profile": profile,
	})
}

func (p *ProfileHandler) DeleteProfile(c *gin.Context) {
	userID := c.Param("user_id")
	err := p.profileUsecase.DeleteProfile(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Profile deleted successfully",
	})
}

func (p *ProfileHandler) UpdateProfile(c *gin.Context) {
	var profile domain.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := p.profileUsecase.UpdateProfile(&profile)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Profile updated successfully",
	})
}
