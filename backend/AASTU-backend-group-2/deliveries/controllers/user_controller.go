package controllers

import (
	"blog_g2/domain"
	"blog_g2/infrastructure"
	"net/mail"
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
	err = uc.Userusecase.RegisterUser(c, user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User registered successfully"})
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

}

// LogoutUser is a controller method to logout a user
func (uc *UserController) LogoutUser(c *gin.Context) {

}

// PromoteDemoteUser is a controller method to promote or demote a user
func (uc *UserController) PromoteDemoteUser(c *gin.Context) {

}
