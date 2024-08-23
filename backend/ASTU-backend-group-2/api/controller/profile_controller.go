package controller

import (
	"context"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/assetutil"
	"github.com/cloudinary/cloudinary-go/v2"
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

		c.JSON(http.StatusOK, gin.H{"user": user.ToUserOut()})
	}
}

func (pc *ProfileController) GetProfiles() gin.HandlerFunc {
	return func(c *gin.Context) {

		var page int64 = 1
		var limit int64 = 10

		in_page, err := strconv.ParseInt(c.Query("page"), 10, 64)
		if err == nil {
			page = in_page
		}

		in_limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
		if err == nil {
			limit = in_limit
		}

		dateFrom, _ := time.Parse(time.RFC3339, c.Query("date_from"))
		dateTo, _ := time.Parse(time.RFC3339, c.Query("date_to"))

		var userFilter domain.UserFilter

		userFilter = domain.UserFilter{
			Email:     c.Query("email"),
			FirstName: c.Query("first_name"),
			LastName:  c.Query("last_name"),
			Role:      c.Query("role"),
			IsOwner:   c.Query("is_owner"),
			Active:    c.Query("active"),
			Bio:       c.Query("bio"),
			DateFrom:  dateFrom,
			DateTo:    dateTo,
			Limit:     limit,
			Pages:     page,
		}

		users, pagination, err := pc.UserUsecase.GetUsers(context.Background(), userFilter)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		res := domain.PaginatedResponse{
			Data:     users,
			MetaData: pagination,
		}

		c.JSON(http.StatusOK, res)
	}
}

func (pc *ProfileController) ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var UpdatedPassword domain.UpdatePassword
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

		user, err := pc.UserUsecase.GetUserById(c, userIDStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UpdatedPassword.OldPassword))
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Old password is not correct"})
			return
		}

		err = pc.UserUsecase.UpdateUserPassword(c, userIDStr, &UpdatedPassword)
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
		var user domain.UserUpdate
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
			return
		}

		updatedUser, err := pc.UserUsecase.UpdateUser(c, userID, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": updatedUser.ToUserOut()})
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
func (pc *ProfileController) UploadProfilePicture(cloudinary *cloudinary.Cloudinary) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("x-user-id").(string)

		filename, ok := c.Get("filePath")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "filename not found"})
			return
		}

		file, ok := c.Get("file")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file not found"})
			return
		}

		imageUrl, err := assetutil.UploadToCloudinary(file.(multipart.File), filename.(string), cloudinary)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Admin demoted to user successfully"})
	}
}
