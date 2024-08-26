package gin

import (
	"net/http"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/auth"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	authuserusecase auth.AuthServices
}

func NewUserController(authServices auth.AuthServices) *UserController {
	return &UserController{
		authuserusecase: authServices,
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
	userid := ctx.Value("user_id").(string)

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

	user, err := uc.authuserusecase.Activate(ctx, userID, token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message activated": user})
}

func (uc *UserController) Logout(ctx *gin.Context) {
	userid := ctx.Value("user_id")
	uc.authuserusecase.Logout(ctx, userid.(string))
	ctx.JSON(http.StatusOK, gin.H{"message": "loged out successfully"})
}

func (uc *UserController) PromoteUser(ctx *gin.Context) {
	userID := ctx.Param("userID")
	err := uc.authuserusecase.PromoteUser(ctx, userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": " un able to promote user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "promored "})

}

func (uc *UserController) DemoteUser(ctx *gin.Context) {
	userID := ctx.Param("userID")
	err := uc.authuserusecase.DemoteUser(ctx, userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": " unable to demote user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "demoted "})
}

func (uc *UserController) ForgetPassword(ctx *gin.Context) {
	var email auth.Email
	if err := ctx.ShouldBindJSON(&email); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err := uc.authuserusecase.ForgetPassword(ctx, email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "reset email is sent"})
}

func (uc *UserController) ResetPassword(ctx *gin.Context) {
	uc.authuserusecase.ResetPassword(ctx)
}
