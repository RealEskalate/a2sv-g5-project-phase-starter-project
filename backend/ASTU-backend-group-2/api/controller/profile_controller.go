package controller

import (
	"net/http"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/utils"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// interface for blog controllers
type profileController interface {
	GetProfile() gin.HandlerFunc
	UpdateProfile() gin.HandlerFunc
	DeleteProfile() gin.HandlerFunc
	PromoteUser() gin.HandlerFunc
	DemoteUser() gin.HandlerFunc
	ChangePassword() gin.HandlerFunc
}

// ProfileController is a struct to hold the usecase and env
type ProfileController struct {
	UserUsecase domain.UserUsecase
	Env         *bootstrap.Env
}

func (pc *ProfileController) GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		user, err := pc.UserUsecase.GetUserById(c, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}
func (pc *ProfileController) ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var  UpdatedPassword domain.UpdatePassword
		if err := c.ShouldBindJSON(&UpdatedPassword); err != nil {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
			return
		}
		userID, exists := c.Get("x-user-id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
			return
		}
	
		// Now you can use userID which is of type interface{}
		userIDStr, ok := userID.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID is not a valid string"})
			return
		}

		user,err:=pc.UserUsecase.GetUserById(c, userIDStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}
		
		err=bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UpdatedPassword.OldPassword))
		if err!=nil{
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Old password is not correct"})
			return
		}

		err=pc.UserUsecase.UpdateUserPassword(c,userIDStr,&UpdatedPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
}
func (pc *ProfileController) UpdateProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		var user domain.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
			return
		}

		err := pc.UserUsecase.UpdateUser(c, userID, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
	}
}

func (pc *ProfileController) DeleteProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		err := pc.UserUsecase.DeleteUser(c, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Profile deleted successfully"})
	}
}

func (pc *ProfileController) PromoteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		err := pc.UserUsecase.PromoteUserToAdmin(c, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin successfully"})
	}
}

func (pc *ProfileController) DemoteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		err := pc.UserUsecase.DemoteAdminToUser(c, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Admin demoted to user successfully"})
	}
}
func (pc *ProfileController) UploadProfilePicture() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("x-user-id").(string)
		uploader := utils.FileUploader{}
		filename, err := uploader.UploadImgFile(c, "profile_pic")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		pc.UserUsecase.UpdateProfilePicture(c, userID, filename)
		c.JSON(http.StatusCreated, gin.H{"message": "profile picture updated"})

	}
}