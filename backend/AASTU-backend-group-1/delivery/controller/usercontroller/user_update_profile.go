package usercontroller

import (
	"blogs/config"
	"blogs/domain"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (u *UserController) UpdateProfile(ctx *gin.Context) {
	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		log.Println("Error getting claims")
		ctx.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	// Initialize avatarPath
	avatarPath := ""

	// Check if MultipartForm is not nil before accessing it
	if ctx.Request.MultipartForm != nil {
		// Check if the "avatar" field exists in the form data
		if _, headerExists := ctx.Request.MultipartForm.File["avatar"]; headerExists {
			// Attempt to get the file from the request
			fileHeader, err := ctx.FormFile("avatar")

			// Check if a file was uploaded
			if err == nil && fileHeader != nil {
				// A file was uploaded, process it
				filename := filepath.Base(fileHeader.Filename)
				avatarPath = filepath.Join("uploads", filename)

				if err := ctx.SaveUploadedFile(fileHeader, avatarPath); err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the file"})
					return
				}
			} else if err != http.ErrMissingFile && err != nil {
				// If the error is not due to a missing file, handle it
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error uploading file"})
				return
			}
		}
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

	err := ctx.ShouldBind(&userData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
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
	}

	err = u.UserUsecase.UpdateProfile(user, claims)
	if err != nil {
		code := config.GetStatusCode(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(code, gin.H{"error": "Internal server error"})
			return
		}

		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User profile updated successfully",
	})
}
