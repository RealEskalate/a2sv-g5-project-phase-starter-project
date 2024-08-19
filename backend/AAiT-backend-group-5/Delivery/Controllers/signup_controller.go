package controllers

import (
	"net/http"

	config "github.com/aait.backend.g5.main/backend/Config"
	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	SignupUsecase interfaces.SignupUsecase
	Env           *config.Env
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
