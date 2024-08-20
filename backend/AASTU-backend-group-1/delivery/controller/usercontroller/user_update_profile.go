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
		log.Println("Error getting claims")
		ctx.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	// Initialize avatarPath
	avatarPath := ""

	// Check if a file was uploaded
	file, err := ctx.FormFile("avatar")
	if err != nil && err != http.ErrMissingFile {
		// Handle the error only if it's not because the file is missing
		log.Println("Error retrieving file:", err)
		ctx.JSON(http.StatusBadRequest, "Error uploading file")
		return
	}

	if file != nil {
		// Upload to Cloudinary and get the URL
		avatarPath, err = config.UploadToCloudinary(file)
		if err != nil {
			log.Println("Error uploading file to Cloudinary:", err)
			ctx.JSON(http.StatusInternalServerError, "Internal server error")
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
		log.Println("Avatar path set for user:", user.Avatar) // Log the avatar in user object
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
