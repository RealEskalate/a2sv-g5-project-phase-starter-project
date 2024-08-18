package controllers

import (
	"blogapp/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type profile_controller struct {
	userUseCase Domain.ProfileUseCases
}

func NewProfileController(service_reference Domain.ProfileUseCases) *profile_controller {
	return &profile_controller{
		userUseCase: service_reference,
	}
}

func (uc *profile_controller) GetProfile(c *gin.Context) {
	cur_user, err := Getclaim(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}

	user, err, statusCode := uc.userUseCase.GetProfile(c, cur_user.ID, *cur_user)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(statusCode, gin.H{"Profile": user})
	}
}

func (uc *profile_controller) UpdateProfile(c *gin.Context) {
	logeduser, err := Getclaim(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}

	var user Domain.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	OmitedUser, err, statusCode := uc.userUseCase.UpdateProfile(c, logeduser.ID, user, *logeduser)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(statusCode, gin.H{"Profile": OmitedUser})
	}
}

func (uc *profile_controller) DeleteProfile(c *gin.Context) {
	user, err := Getclaim(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}

	err, statusCode := uc.userUseCase.DeleteProfile(c, user.ID, *user)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(statusCode, gin.H{"message": "Profile deleted successfully"})
	}
}
