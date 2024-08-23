package controllers

import (
	"blogapp/Domain"
	"blogapp/Utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Profile_controller struct {
	userUseCase Domain.ProfileUseCases
}

func NewProfileController(service_reference Domain.ProfileUseCases) *Profile_controller {
	return &Profile_controller{
		userUseCase: service_reference,
	}
}

func (uc *Profile_controller) GetProfile(c *gin.Context) {
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

func (uc *Profile_controller) UpdateProfile(c *gin.Context) {

	logeduser, err := Getclaim(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}
	// updateUser := Domain.UpdateUser{}

	// if err := c.BindJSON(&updateUser); err != nil {
	// 	fmt.Println("i am at thr top")
	// 	c.IndentedJSON(400, gin.H{"error": err.Error()})
	// 	return
	// }
	// get profile picture image from request
	file, _ := c.FormFile("profilepicture")
	// if err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	var user Domain.User
	user.ID = logeduser.ID
	user.Name = c.PostForm("name")
	user.UserName = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.Bio = c.PostForm("bio")
	if file != nil {
		profilePicture, err := Utils.SetProfilePicture(file)
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

func (uc *Profile_controller) DeleteProfile(c *gin.Context) {
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
