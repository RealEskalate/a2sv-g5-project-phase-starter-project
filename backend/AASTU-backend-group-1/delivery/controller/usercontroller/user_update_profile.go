package usercontroller

import (
	"blogs/config"
	"blogs/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) UpdateProfile(ctx *gin.Context) {
	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   "cannot get claims",
		})
		return
	}

	// Initialize avatarPath
	avatarPath := ""

	// Check if a file was uploaded
	file, err := ctx.FormFile("avatar")
	if err != nil && err != http.ErrMissingFile {
		// Handle the error only if it's not because the file is missing
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	if file != nil {
		// Upload to Cloudinary and get the URL
		avatarPath, err = config.UploadToCloudinary(file)
		if err != nil {
			log.Println("Error uploading file to Cloudinary:", err)
			ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error uploading file to Cloudinary",
				Error:   err.Error(),
			})
			return
		}
		log.Println("Avatar successfully uploaded to Cloudinary:", avatarPath)
	}

	// Handle other form data
	var userData struct {
		FirstName string `form:"firstname"`
		LastName  string `form:"lastname"`
		Bio       string `form:"bio"`
		Username  string `form:"username"`
		Password  string `form:"password"`
		Email     string `form:"email"`
		Address   string `form:"address"`
	}

	err = ctx.ShouldBind(&userData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	// Prepare user data for update
	user := &domain.User{
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Bio:       userData.Bio,
		Username:  userData.Username,
		Password:  userData.Password,
		Email:     userData.Email,
		Address:   userData.Address,
	}

	// If a new avatar was uploaded, update the Avatar field
	if avatarPath != "" {
		user.Avatar = avatarPath
		log.Println("Avatar path set for user:", user.Avatar) // Log the avatar in user object
	}

	err = u.UserUsecase.UpdateProfile(user, claims)
	if err != nil {
		code := config.GetStatusCode(err)
		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Failed to update user profile",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "User profile updated successfully",
	})
}
