package controllers

import (
	"blog_g2/domain"
	"blog_g2/infrastructure"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Userusecase domain.UserUsecase
}

// Blog-controller constructor
func NewUserController(Usermgr domain.UserUsecase) *UserController {
	return &UserController{
		Userusecase: Usermgr,
	}
}

// RegisterUser is a controller method to register a user
func (uc *UserController) RegisterUser(c *gin.Context) {

	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if user.Email == "" || user.Password == "" || user.UserName == "" {
		c.JSON(400, gin.H{"error": "Please provide all fields"})
		return
	}
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid email address"})
		return
	}

	if err := infrastructure.PasswordValidator(user.Password); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.JoinedAt = time.Now()
	user.IsAdmin = false
	err = uc.Userusecase.RegisterUser(c, &user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User registered successfully", "user": user})
}

// LoginUser is a controller method to login a user
func (uc *UserController) LoginUser(c *gin.Context) {

	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if user.Email == "" || user.Password == "" {
		c.JSON(400, gin.H{"error": "Please provide all fields"})
		return
	}
	token, err := uc.Userusecase.LoginUser(c, user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user logged in", "token": token})

}

// ForgotPassword is a controller method to reset a user's password
func (uc *UserController) ForgotPassword(c *gin.Context) {

	var info domain.RestRequest
	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := uc.Userusecase.ForgotPassword(c, info.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "email succefully sent to the email provided"})
}

func (uc *UserController) ResetPassword(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, "Token is required")
		return
	}

	// Log the token for debugging purposes
	fmt.Printf("Received Token: %s\n", token)

	var info domain.RestRequest
	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := uc.Userusecase.ResetPassword(c, token, info.NewPassword)
	if err != nil {
		fmt.Printf("Error resetting password: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Password has been reset successfully"})
}

// LogoutUser is a controller method to logout a user
func (uc *UserController) LogoutUser(c *gin.Context) {
	userid := c.GetString("userid")
	log.Println(userid)
	err := uc.Userusecase.LogoutUser(c, userid)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User logged out successfully"})

}

// PromoteDemoteUser is a controller method to promote or demote a user
func (uc *UserController) PromoteDemoteUser(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id is empty"})
		return
	}

	isAdminstr := c.Query("isadmin")
	if isAdminstr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "isAdmin is empty"})
		return
	}

	isAdmin, err := strconv.ParseBool(isAdminstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = uc.Userusecase.PromoteDemoteUser(c, ID, isAdmin)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "user admin privilege succesfully updated"})

}
