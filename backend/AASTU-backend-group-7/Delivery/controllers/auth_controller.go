package controllers

import (
	"blogapp/Domain"

	"github.com/gin-gonic/gin"
)

type authController struct {
	AuthUseCase Domain.AuthUseCase
}
func NewAuthController(usecase Domain.AuthUseCase) (*authController) {

	return &authController{
		AuthUseCase: usecase,
	}
}

// login
func (ac *authController) Login(c *gin.Context) {
	// return error
	 
}

// register
func (ac *authController) Register(c *gin.Context) {
	// return error
 
}