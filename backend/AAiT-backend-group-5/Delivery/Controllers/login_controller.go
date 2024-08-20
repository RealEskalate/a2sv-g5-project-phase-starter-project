package controllers

import (
	"net/http"

	config "github.com/aait.backend.g5.main/backend/Config"
	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUsecase interfaces.LoginUsecase
	Env          *config.Env
}

func NewLoginController(loginUsecase interfaces.LoginUsecase, env *config.Env) *LoginController {
	return &LoginController{
		LoginUsecase: loginUsecase,
		Env:          env,
	}
}

func (loginController *LoginController) Login(ctx *gin.Context) {
	var loginRequest dtos.LoginRequest

	// attempt to bind the json payload
	err := ctx.ShouldBind(&loginRequest)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// envode Login_Usecase
	loginResponse, e := loginController.LoginUsecase.LoginUser(ctx, loginRequest)
	if e != nil {
		ctx.IndentedJSON(e.Code, gin.H{"error": e.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, loginResponse)
}
