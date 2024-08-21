package controllers

import (
	"blogapp/Domain"
	"blogapp/Utils"
	"fmt"
	"log"
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
	fmt.Println("hehehe")
	fmt.Println("hehehe")
	log.Println("hehehe")
	log.Println("hehehe")
	logeduser, err := Getclaim(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}
	updateUser := Domain.UpdateUser{}	
	
	if err := c.BindJSON(&updateUser); err != nil {
		fmt.Println("i am at thr top")
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	// get profile picture image from request
	file, _ := c.FormFile("profilepicture")
	
	var user Domain.User
	user.ID = logeduser.ID
	user.Name = updateUser.Name
	user.UserName = updateUser.UserName
	user.Email = updateUser.Email
	user.Password = updateUser.Password
	user.Bio = updateUser.Bio
	if file != nil {
		profilePicture, err:= Utils.SetProfilePicture(file)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user.ProfilePicture = profilePicture
	}
	
	OmitedUser, err, statusCode := uc.userUseCase.UpdateProfile(c, logeduser.ID, user, *logeduser)
	if err != nil {
		fmt.Println("i was here all along ")
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
