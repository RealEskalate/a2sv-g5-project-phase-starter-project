package controller

import (
	"blog/config"
	"blog/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
	Env            *config.Env
}

// UpdateProfile updates a user's profile
func (pc *ProfileController) UpdateProfile(c *gin.Context) {
	var profile domain.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	claims := c.MustGet("claim").(domain.JwtCustomClaims)
	id, _ := primitive.ObjectIDFromHex(claims.Id)
	resp, err := pc.ProfileUsecase.UpdateProfile(c, &profile, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated", "data": resp})
}