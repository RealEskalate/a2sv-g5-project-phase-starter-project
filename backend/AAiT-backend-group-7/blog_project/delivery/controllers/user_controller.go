package controllers

import (
	"blog_project/domain"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	UserUsecase domain.IUserUsecase
}

func NewUserController(userUsecase domain.IUserUsecase) domain.IUserController {
	return &userController{UserUsecase: userUsecase}
}

func (uc *userController) GetAllUsers(c *gin.Context) {
	users, err := uc.UserUsecase.GetAllUsers(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Prepend the URL to the ProfilePic path
	for i := range users {
		if users[i].ProfilePic != "" {
			users[i].ProfilePic = fmt.Sprintf("http://localhost:8080/%s", users[i].ProfilePic)
		}
	}

	c.JSON(200, users)
}

func (uc *userController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.UserUsecase.GetUserByID(c, idInt)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Prepend the URL to the ProfilePic path
	if user.ProfilePic != "" {
		user.ProfilePic = fmt.Sprintf("http://localhost:8080/%s", user.ProfilePic)
	}

	c.JSON(200, user)
}

func (uc *userController) CreateUser(c *gin.Context) {
	var user domain.User

	// Parse form data
	err := c.Request.ParseMultipartForm(10 << 20) // Limit upload size to 10MB
	if err != nil {
		c.JSON(400, gin.H{"error": "Could not parse form data"})
		return
	}

	// Retrieve JSON fields
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.Email = c.PostForm("email")
	user.Role = c.PostForm("role")
	user.Bio = c.PostForm("bio")
	user.Phone = c.PostForm("phone")

	// Handle profile picture upload (optional)
	file, handler, err := c.Request.FormFile("profile_pic")
	if err == nil { // Only proceed if the file was uploaded
		defer file.Close()

		// Define the path where the file will be stored
		filePath := fmt.Sprintf("./uploads/%s", handler.Filename)

		// Save the file to disk
		out, err := os.Create(filePath)
		if err != nil {
			c.JSON(500, gin.H{"error": "Unable to create the file for writing: " + err.Error()})
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			c.JSON(500, gin.H{"error": "Unable to save the file"})
			return
		}

		// Set the profile picture path
		user.ProfilePic = "/uploads/" + handler.Filename
	}

	newUser, err := uc.UserUsecase.CreateUser(c, user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, newUser)
}

func (uc *userController) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user domain.User

	// Parse form data
	err = c.Request.ParseMultipartForm(10 << 20) // Limit upload size to 10MB
	if err != nil {
		c.JSON(400, gin.H{"error": "Could not parse form data"})
		return
	}

	// Retrieve JSON fields
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.Email = c.PostForm("email")
	user.Role = c.PostForm("role")
	user.Bio = c.PostForm("bio")
	user.Phone = c.PostForm("phone")

	// Handle profile picture upload (optional)
	file, handler, err := c.Request.FormFile("profile_pic")
	if err == nil { // Only proceed if the file was uploaded
		defer file.Close()

		// Define the path where the file will be stored
		filePath := fmt.Sprintf("./uploads/%s", handler.Filename)

		// Save the file to disk
		out, err := os.Create(filePath)
		if err != nil {
			c.JSON(500, gin.H{"error": "Unable to create the file for writing: " + err.Error()})
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			c.JSON(500, gin.H{"error": "Unable to save the file"})
			return
		}

		// Set the profile picture path
		user.ProfilePic = "/uploads/" + handler.Filename
	}

	newUser, err := uc.UserUsecase.UpdateUser(c, idInt, user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, newUser)
}

func (uc *userController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = uc.UserUsecase.DeleteUser(c, idInt)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

func (uc *userController) AddBlog(c *gin.Context) {
	userID := c.Param("userID")

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var Blog domain.Blog
	err = c.BindJSON(&Blog)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newUser, err := uc.UserUsecase.AddBlog(c, userIDInt, Blog)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, newUser)
}

func (uc *userController) Login(c *gin.Context) {
	var user domain.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := uc.UserUsecase.Login(c, user.Username, user.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"access token": accessToken, "refresh token": refreshToken})
}

func (uc *userController) Logout(c *gin.Context) {
	var token struct {
		Token string `json:"token"`
	}
	err := c.ShouldBindJSON(&token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = uc.UserUsecase.Logout(c, token.Token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Logged out successfully"})

}

func (uc *userController) ForgetPassword(c *gin.Context) {
	email := c.Param("email")
	err := uc.UserUsecase.ForgetPassword(c, email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Password reset link sent to email"})
}

func (uc *userController) ResetPassword(c *gin.Context) {

	username := c.Param("username")
	password := c.Param("password")
	err := uc.UserUsecase.ResetPassword(c, username, password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Password reset successfully"})
}

func (uc *userController) PromoteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.UserUsecase.PromoteUser(c, idInt)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User promoted successfully", "user": user})
}

func (uc *userController) DemoteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.UserUsecase.DemoteUser(c, idInt)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User demoted successfully", "user": user})
}

func (uc *userController) RefreshToken(c *gin.Context) {
	var refreshToken struct {
		RefreshToken string `json:"refresh_token"`
	}
	err := c.ShouldBindJSON(&refreshToken)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newToken, err := uc.UserUsecase.RefreshToken(c, refreshToken.RefreshToken)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Token refreshed successfully", "new_access_token": newToken})
}
