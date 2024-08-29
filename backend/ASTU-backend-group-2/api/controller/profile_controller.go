package controller

import (
	"context"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/middleware"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	custom_error "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/errors"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/assetutil"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// interface for blog controllers
type profileController interface {
	GetProfile() gin.HandlerFunc
	GetCurrentProfile() gin.HandlerFunc
	UpdateCurrentProfile() gin.HandlerFunc
	DeleteCurrentProfile() gin.HandlerFunc
	UpdateProfile() gin.HandlerFunc
	DeleteProfile() gin.HandlerFunc
	PromoteUser() gin.HandlerFunc
	DemoteUser() gin.HandlerFunc
	ChangePassword() gin.HandlerFunc
}

// ProfileController is a struct to hold the usecase and env
type ProfileController struct {
	UserUsecase entities.UserUsecase
	Env         *bootstrap.Env
}

func (pc *ProfileController) GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		claimUserID := c.MustGet("x-user-id").(string)
		role := c.MustGet("x-user-role").(string)
		user, err := pc.UserUsecase.GetUserById(c, userID)
		if err != nil {
			c.Error(err)
			return
		}
		if userID != claimUserID && role != "admin" {
			c.JSON(http.StatusUnauthorized, custom_error.ErrorMessage{Message: "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

func (pc *ProfileController) GetCurrentProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("x-user-id").(string)

		user, err := pc.UserUsecase.GetUserById(c, userID)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

func (pc *ProfileController) UpdateCurrentProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("x-user-id").(string)

		var user entities.UserUpdate
		if err := c.ShouldBindJSON(&user); err != nil {
			if err == io.EOF {
				c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.EreInvalidRequestBody))
				return
			}
			if user.FirstName == "" && user.LastName == "" && user.Bio == "" {
				c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.EreInvalidRequestBody))
				return
			}
			middleware.CustomErrorResponse(c, err)
			return
		}

		newUser, err := pc.UserUsecase.UpdateUser(c, userID, &user)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": newUser})
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

		var userFilter entities.UserFilter

		userFilter = entities.UserFilter{
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
			c.Error(err)
			return
		}

		res := entities.PaginatedResponse{
			Data:     users,
			MetaData: pagination,
		}

		c.JSON(http.StatusOK, res)
	}
}

func (pc *ProfileController) ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updatedPassword entities.UpdatePassword
		if err := c.ShouldBindJSON(&updatedPassword); err != nil {
			if err == io.EOF {
				c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.EreInvalidRequestBody))
				return
			}
			middleware.CustomErrorResponse(c, err)
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
			c.Error(err)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(updatedPassword.OldPassword))
		if err != nil {
			c.Error(err)
			return
		}

		err = pc.UserUsecase.UpdateUserPassword(c, userIDStr, &updatedPassword)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
	}
}
func (pc *ProfileController) UpdateProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		claimUserID := c.MustGet("x-user-id").(string)
		role := c.MustGet("x-user-role").(string)

		var user entities.UserUpdate
		if err := c.ShouldBindJSON(&user); err != nil {
			if err == io.EOF {
				c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.EreInvalidRequestBody))
				return
			}
			if user.FirstName == "" && user.LastName == "" && user.Bio == "" {
				c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.EreInvalidRequestBody))
				return
			}
			middleware.CustomErrorResponse(c, err)
			return
		}
		if userID != claimUserID && role != "admin" {
			c.JSON(http.StatusUnauthorized, custom_error.ErrorMessage{Message: "unauthorized"})
			return
		}
		updatedUser, err := pc.UserUsecase.UpdateUser(c, userID, &user)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": updatedUser})
	}
}

func (pc *ProfileController) DeleteProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		claimUserID := c.MustGet("x-user-id").(string)
		role := c.MustGet("x-user-role").(string)

		if userID != claimUserID && role != "admin" {
			c.JSON(http.StatusUnauthorized, custom_error.ErrorMessage{Message: "unauthorized"})
			return
		}

		err := pc.UserUsecase.DeleteUser(c, userID)

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}

func (pc *ProfileController) DeleteCurrentProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("x-user-id").(string)

		err := pc.UserUsecase.DeleteUser(c, userID)

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}

func (pc *ProfileController) PromoteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		err := pc.UserUsecase.PromoteUserToAdmin(c, userID)
		if err != nil {
			c.Error(err)
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
			c.Error(err)
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
		log.Println(imageUrl, userID)
		if err != nil {
			c.Error(err)
			return
		}
		pc.UserUsecase.UpdateProfilePicture(c, userID, imageUrl)
		c.JSON(http.StatusOK, gin.H{"message": "Admin demoted to user successfully"})
	}
}
