package controllers

import (
	"net/http"

	"aait.backend.g10/usecases"
	"aait.backend.g10/usecases/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userUseCase usecases.IUserUseCase
}

func NewUserController(u usecases.IUserUseCase) *UserController {
	return &UserController{
		userUseCase: u,
	}
}

func (u *UserController) PromoteUser(c *gin.Context) {
	promote := struct {
		ID        uuid.UUID `json:"id" binding:"required"`
		IsPromote *bool      `json:"is_promote" binding:"required"`
	}{}
	if err := c.BindJSON(&promote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cerr := u.userUseCase.PromoteUser(promote.ID, *promote.IsPromote)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}

	if *promote.IsPromote {
		c.JSON(http.StatusOK, gin.H{"message": "User promoted successfully"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User demoted successfully"})	
	}
}

func (u *UserController) UpdateProfile(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var user dto.UserUpdate
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requester_id, err := uuid.Parse(c.MustGet("id").(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	user.ID = id
	cerr := u.userUseCase.UpdateUser(requester_id, &user)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully"})
}

func (u *UserController) GetUserByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	user, cerr := u.userUseCase.GetUserByID(id)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) UploadProfilePic(c *gin.Context) {
	// Get the user's ID from context (or from session, JWT token, etc.)
	userID, _ := c.Get("id")
	// Assuming the user ID is passed as a URL parameter

	// Retrieve the file from the form data
	file, header, err := c.Request.FormFile("profile_pic")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file" + "c.request"})
		return
	}
	defer file.Close()
	filePath, errs := uc.userUseCase.UploadProfilePic(userID.(uuid.UUID), &file, header)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile picture uploaded successfully", "profile_pic_url": filePath})
}
