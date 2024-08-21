package controllers

import (
	"astu-backend-g1/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (c *UserController) Register(ctx *gin.Context) {
	user := domain.User{}
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	_, err := c.userUsecase.Create(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Activate your Account in the your email link"})
}

func (c *UserController) AccountVerification(ctx *gin.Context) {
	email := ctx.Query("email")
	token := ctx.Query("token")
	err := c.userUsecase.AccountVerification(email, token)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Account Activated"})
}

func (c *UserController) ForgetPassword(ctx *gin.Context) {
	email := ctx.Query("email")
	_, err := c.userUsecase.ForgetPassword(email)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Password reset token sent to your email"})
}

func (c *UserController) ResetPassword(ctx *gin.Context) {
	email := ctx.Query("email")
	token := ctx.Query("token")
	newPassword := struct {
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}{}
	if err := ctx.BindJSON(&newPassword); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	if newPassword.Password != newPassword.ConfirmPassword {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "the password and confirm password should be the same"})
		return
	}
	_, err := c.userUsecase.ResetPassword(email, token, newPassword.Password)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Password Reset Successful"})
}

func (c *UserController) ForgetPasswordUser(ctx *gin.Context) {
	email := ""
	if err := ctx.BindJSON(&email); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	_, err := c.userUsecase.ForgetPassword(email)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Password reset token sent to your email"})
}

func (c *UserController) LoginUser(ctx *gin.Context) {
	user := &domain.User{}
	if err := ctx.ShouldBind(user); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	access_token, err := c.userUsecase.LoginUser(user.Username, user.Password)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"access_token": access_token})
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	username := ctx.Query("username")
	email := ctx.Query("email")
	if username == "" && email == "" {
		users, err := c.userUsecase.Get()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusOK, users)
		return
	} else if username != "" {
		user, err := c.userUsecase.GetByUsername(username)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusOK, user)
		return
	} else if email != "" {
		user, err := c.userUsecase.GetByEmail(email)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusOK, user)
		return
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	userID := ctx.Param("id")
	user, err := c.userUsecase.GetByID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	err := c.userUsecase.Delete(userId)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	ctx.String(http.StatusNoContent, "")
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	updateData := domain.User{}
	if err := ctx.ShouldBind(updateData); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := c.userUsecase.Update(userId, updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, updatedUser)
}
