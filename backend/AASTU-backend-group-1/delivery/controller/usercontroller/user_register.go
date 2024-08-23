package usercontroller

import (
	"blogs/config"
	"blogs/domain"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (u *UserController) RegisterUser(ctx *gin.Context) {

	// Initialize avatarPath
	avatarPath := ""

	// Check if a file was uploaded
	file, err := ctx.FormFile("avatar")
	if err != nil && err != http.ErrMissingFile {
		// Handle the error only if it's not because the file is missing
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Error retrieving file",
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

	var userData struct {
		FirstName string `form:"firstname"`
		LastName  string `form:"lastname"`
		Bio       string `form:"bio"`
		Username  string `form:"username"`
		Password  string `form:"password"`
		Email     string `form:"email"`
		Address   string `form:"address"`
		Avatar    string `form:"avatar"`
	}

	err = ctx.ShouldBind(&userData)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	log.Println("User data:", userData)

	if userData.Username == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "username is required",
		})
		return
	}

	if userData.Email == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "email is required",
		})
		return
	}

	if userData.Password == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "password is required",
		})
		return
	}

	user := &domain.User{
		FirstName:  userData.FirstName,
		LastName:   userData.LastName,
		Username:   userData.Username,
		Password:   userData.Password,
		Email:      userData.Email,
		Bio:        userData.Bio,
		Avatar:     userData.Avatar,
		Address:    userData.Address,
		Role:       "user",
		JoinedDate: time.Now(),
		IsVerified: false,
	}

	if avatarPath != "" {
		user.Avatar = avatarPath
		log.Println("Avatar path set for user:", user.Avatar) // Log the avatar in user object
	}

	err = u.UserUsecase.RegisterUser(user)
	if err != nil {
		code := config.GetStatusCode(err)
		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Failed to register user",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, domain.APIResponse{
		Status:  http.StatusCreated,
		Message: "User registered successfully",
	})
}
