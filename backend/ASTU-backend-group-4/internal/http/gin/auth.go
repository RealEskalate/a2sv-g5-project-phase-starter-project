package gin

import (
	"net/http"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/auth"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	authuserusecase auth.AuthUserUsecase
}

func NewUserController(authuserusecase auth.AuthUserUsecase) UserController {
	return UserController{
		authuserusecase: authuserusecase,
	}
}

func (uc *UserController) Login(ctx *gin.Context) {
	var userInfo auth.LoginForm
	if err := ctx.ShouldBindJSON(&userInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refToken, accessToken, err := uc.authuserusecase.Login(ctx, userInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"access token": accessToken})
	ctx.JSON(http.StatusOK, gin.H{"refresh token": refToken})
}

func (uc *UserController) RegisterUser(ctx *gin.Context) {
	var user auth.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.authuserusecase.RegisterUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully registered"})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user auth.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})

}
