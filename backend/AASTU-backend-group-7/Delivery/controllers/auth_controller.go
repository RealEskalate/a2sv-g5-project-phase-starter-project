package controllers

import (
	"blogapp/Domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type authController struct {
	AuthUseCase Domain.AuthUseCase
}

func NewAuthController(usecase Domain.AuthUseCase) *authController {

	return &authController{
		AuthUseCase: usecase,
	}
}

// login
func (ac *authController) Login(c *gin.Context) {
	var newUser Domain.User
	v := validator.New()
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data", "error": err.Error()})
		return
	}
	if err := v.Struct(newUser); err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid or missing data", "error": err.Error()})
		return
	}
	token, err, statusCode := ac.AuthUseCase.Login(c, &newUser)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		//success
		c.IndentedJSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
	}

}

// register
func (ac *authController) Register(c *gin.Context) {
	// return error
	var newUser Domain.User
	v := validator.New()
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data", "error": err.Error()})
		return
	}

	if err := v.Struct(newUser); err != nil {
		fmt.Printf(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid or missing data", "error": err.Error()})
		return
	}

	createdUser, err, statusCode := ac.AuthUseCase.Register(c, &newUser)

	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		//success
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": createdUser})
	}

}
