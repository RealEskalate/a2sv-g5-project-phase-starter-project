package controller

import (
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	"github.com/gin-gonic/gin"
)

// interface for blog controllers
type profileController interface {
	GetProfile() gin.HandlerFunc
	UpdateProfile() gin.HandlerFunc
	DeleteProfile() gin.HandlerFunc
	PromoteUser() gin.HandlerFunc
	DemoteUser() gin.HandlerFunc
}

// ProfileController is a struct to hold the usecase and env
type ProfileController struct {
	UserUsecase domain.UserUsecase
	Env         *bootstrap.Env
}

func (pc *ProfileController) GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		user, err := pc.UserUsecase.GetUserProfile(c, userID)
		if err != nil {
			c.JSON(500, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(200, user)
	}
}

func (pc *ProfileController) UpdateProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		var updatedProfile domain.User
		err := c.ShouldBind(&updatedProfile)
		if err != nil {
			c.JSON(400, domain.ErrorResponse{Message: err.Error()})
			return
		}

		err = pc.UserUsecase.UpdateUserProfile(c, userID, &updatedProfile)
		if err != nil {
			c.JSON(500, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Profile updated successfully"})
	}
}

func (pc *ProfileController) DeleteProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		err := pc.UserUsecase.DeleteUserProfile(c, userID)
		if err != nil {
			c.JSON(500, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Profile deleted successfully"})
	}
}

func (pc *ProfileController) PromoteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		err := pc.UserUsecase.PromoteUserToAdmin(c, userID)
		if err != nil {
			c.JSON(500, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "User promoted to admin successfully"})
	}
}

func (pc *ProfileController) DemoteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		err := pc.UserUsecase.DemoteAdminToUser(c, userID)
		if err != nil {
			c.JSON(500, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Admin demoted to user successfully"})
	}
}
