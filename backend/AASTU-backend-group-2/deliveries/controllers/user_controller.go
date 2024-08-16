package controllers

import (
	"blog_g2/domain"

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

}

// LoginUser is a controller method to login a user
func (uc *UserController) LoginUser(c *gin.Context) {

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
