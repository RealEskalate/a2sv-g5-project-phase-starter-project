package controllers

import (
	"errors"
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

	// attempt to bind the json payload
	err := ctx.ShouldBind(&userCreateRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.BadRequest("invalid request"))
		return
	}

	// populate fields for new user
	newUser := &models.User{
		Username: userCreateRequest.Username,
		Name:     userCreateRequest.Name,
		Email:    userCreateRequest.Email,
	}

	// create user
	e := signupController.SignupUsecase.CreateUser(ctx, newUser)
	if e != nil {
		ctx.JSON(e.Code, e.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "check your email"})
}

func (signupController *SignupController) ForgotPasswordConfirm(ctx *gin.Context) {
	var setUpPasswordRequest *dtos.SetUpPasswordRequest

	// attempt to bind the payload carrying the new password
	err := ctx.ShouldBind(&setUpPasswordRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("invalid request"))
		return
	}

	// get short code from the URL
	shortURLCode := ctx.Param("id")

	e := signupController.PasswordUsecase.UpdateUserPassword(ctx, setUpPasswordRequest.Password, shortURLCode)
	if e != nil {
		ctx.JSON(e.Code, e.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "registeration successful, proceed to login"})
}
