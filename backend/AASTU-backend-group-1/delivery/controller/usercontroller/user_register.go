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
	var userData struct {
		FirstName string `form:"firstname"`
		LastName  string `form:"lastname"`
		Username  string `form:"username" `
		Password  string `form:"password"`
		Email     string `form:"email"`
		Bio       string `form:"bio"`
		Avatar    string `form:"avatar"`
		Address   string `form:"address"`
	}

	err := ctx.ShouldBind(&userData)
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
			log.Println(err)
			ctx.JSON(code, gin.H{"error": "Internal server error"})
			return
		}

		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "verification email has been sent successfully",
	})
}
