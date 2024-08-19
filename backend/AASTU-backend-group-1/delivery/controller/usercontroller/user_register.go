package usercontroller

import (
	"blogs/config"
	"blogs/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (u *UserController) RegisterUser(ctx *gin.Context) {
	var userData struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Username  string `json:"username" `
		Password  string `json:"password"`
		Email     string `json:"email"`
		Bio       string `json:"bio"`
		Avatar    string `json:"avatar"`
		Address   string `json:"address"`
	}

	err := ctx.ShouldBindJSON(&userData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if userData.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return
	}

	if userData.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	if userData.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
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

	err = u.UserUsecase.RegisterUser(user)
	if err != nil {
		code := config.GetStatusCode(err)

		if code == http.StatusInternalServerError {
			ctx.JSON(code, gin.H{"error": "Internal server error"})
			return
		}

		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user registered successfully",
	})
}
