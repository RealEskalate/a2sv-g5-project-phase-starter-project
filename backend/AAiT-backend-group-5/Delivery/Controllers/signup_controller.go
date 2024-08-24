package controllers

import (
	"net/http"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	SignupUsecase   interfaces.SignupUsecase
	PasswordUsecase interfaces.PasswordUsecase
}

func NewSignupController(signupUsecase interfaces.SignupUsecase, passwordUsecase interfaces.PasswordUsecase) *SignupController {
	return &SignupController{
		SignupUsecase:   signupUsecase,
		PasswordUsecase: passwordUsecase,
	}
}

func (signupController *SignupController) Signup(ctx *gin.Context) {
	var userCreateRequest dtos.CreateAccountRequest

	err := ctx.ShouldBind(&userCreateRequest)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	if err := userCreateRequest.Validate(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "One or more fields are missing"})
		return
	}

	newUser := &models.User{
		Username: userCreateRequest.Username,
		Name:     userCreateRequest.Name,
		Email:    userCreateRequest.Email,
	}

	agent := ctx.Request.UserAgent()

	e := signupController.SignupUsecase.CreateUser(ctx, newUser, agent)
	if e != nil {
		ctx.IndentedJSON(e.Code, gin.H{"error": e.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "check your email"})
}

func (signupController *SignupController) ConfirmRegistration(ctx *gin.Context) {
	var setUpPasswordRequest dtos.SetUpPasswordRequest

	err := ctx.ShouldBind(&setUpPasswordRequest)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	if err := setUpPasswordRequest.Validate(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "One or more fields are missing"})
		return
	}

	// get short code from the URL
	shortURLCode := ctx.Param("id")

	e := signupController.PasswordUsecase.SetNewUserPassword(ctx, shortURLCode, setUpPasswordRequest.Password)
	if e != nil {
		ctx.IndentedJSON(e.Code, gin.H{"error": e.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "registration successful, proceed to login"})
}
