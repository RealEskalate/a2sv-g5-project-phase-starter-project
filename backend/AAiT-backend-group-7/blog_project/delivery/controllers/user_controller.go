package controllers

import (
	"blog_project/domain"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase domain.IUserUsecase
}

func NewUserController(userUsecase domain.IUserUsecase) domain.IUserController {
	return &UserController{userUsecase: userUsecase}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.userUsecase.GetAllUsers(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userUsecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(400, gin.H{"error": "Could not parse form data"})
		return
	}

	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.Email = c.PostForm("email")
	user.Bio = c.PostForm("bio")
	user.Phone = c.PostForm("phone")

	file, handler, err := c.Request.FormFile("profile_pic")
	if err == nil {
		defer file.Close()
		filePath := fmt.Sprintf("./uploads/%s", handler.Filename)
		out, err := os.Create(filePath)
		if err != nil {
			c.JSON(500, gin.H{"error": "Unable to create file: " + err.Error()})
			return
		}
		defer out.Close()
		if _, err = io.Copy(out, file); err != nil {
			c.JSON(500, gin.H{"error": "Unable to save the file"})
			return
		}
		user.ProfilePic = "/uploads/" + handler.Filename
	}

	newUser, err := uc.userUsecase.CreateUser(c, user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, newUser)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user domain.User
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(400, gin.H{"error": "Could not parse form data"})
		return
	}

	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.Email = c.PostForm("email")
	user.Bio = c.PostForm("bio")
	user.Phone = c.PostForm("phone")

	file, handler, err := c.Request.FormFile("profile_pic")
	if err == nil {
		defer file.Close()
		filePath := fmt.Sprintf("./uploads/%s", handler.Filename)
		out, err := os.Create(filePath)
		if err != nil {
			c.JSON(500, gin.H{"error": "Unable to create file: " + err.Error()})
			return
		}
		defer out.Close()
		if _, err = io.Copy(out, file); err != nil {
			c.JSON(500, gin.H{"error": "Unable to save the file"})
			return
		}
		user.ProfilePic = "/uploads/" + handler.Filename
	}

	newUser, err := uc.userUsecase.UpdateUser(c, id, user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, newUser)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := uc.userUsecase.DeleteUser(c, id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

func (uc *UserController) Login(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := uc.userUsecase.Login(c, user.Username, user.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err = uc.userUsecase.GetUserByUsername(c, user.Username)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"access_token": accessToken, "refresh_token": refreshToken, "user": user})
}

func (uc *UserController) Logout(c *gin.Context) {
	var token struct {
		Token string `json:"token"`
	}
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := uc.userUsecase.Logout(c, token.Token); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Logged out successfully"})
}

func (uc *UserController) ForgetPassword(c *gin.Context) {
	email := c.Query("email")
	if err := uc.userUsecase.ForgetPassword(c, email); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Password reset link sent to email"})
}

func (uc *UserController) ResetPassword(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
	if err := uc.userUsecase.ResetPassword(c, username, password); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Password reset successfully"})
}

func (uc *UserController) PromoteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userUsecase.PromoteUser(c, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User promoted successfully", "user": user})
}

func (uc *UserController) DemoteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userUsecase.DemoteUser(c, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User demoted successfully", "user": user})
}

func (uc *UserController) RefreshToken(c *gin.Context) {
	var refreshToken struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&refreshToken); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newToken, err := uc.userUsecase.RefreshToken(c, refreshToken.RefreshToken)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Token refreshed successfully", "new_access_token": newToken})
}
