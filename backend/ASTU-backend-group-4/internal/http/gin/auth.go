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
	ctx.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refToken})
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

func (uc *UserController) UpdateProfile(ctx *gin.Context) {
	var user auth.User
	userid := ctx.Value("userID").(string)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	user.ID = userid

	err := uc.authuserusecase.UpdateProfile(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (uc *UserController) ActivateUser(ctx *gin.Context) {
	token := ctx.Param("token")
	userID := ctx.Param("userID")

	err := uc.authuserusecase.Activate(ctx, userID, token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "activated"})
}

func (uc *UserController) Logout(ctx *gin.Context) {
	userid := ctx.Value("userID")
	uc.authuserusecase.Logout(ctx, userid.(string))
	ctx.JSON(http.StatusOK, gin.H{"message": "loged out successfully"})
}
