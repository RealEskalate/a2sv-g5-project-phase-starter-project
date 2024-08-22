package controllers

import (
	infrastructure "astu-backend-g1/Infrastructure"
	"astu-backend-g1/config"
	"astu-backend-g1/domain"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userUsecase    domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{
		userUsecase:    userUsecase,
	}

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

//logout user

func (c *UserController) LogoutUser(ctx *gin.Context) {
	email := ctx.MustGet("claims").(*domain.Claims).Email
	err := c.userUsecase.Logout(email)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
	
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
func (c *UserController) RefreshAccessToken(ctx *gin.Context) {
	configJwt, err := config.LoadConfig()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var jwtSecret = []byte(configJwt.Jwt.JwtKey)
	type Pass struct {
		Password string `json:"password"`
	}
	var NUID Pass
	err = ctx.ShouldBindJSON(&NUID)
	TheUser, err := c.userUsecase.GetByID(ctx.Param("uid"))
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	fmt.Println("this is the user:", TheUser)
	refreshToken, err := jwt.ParseWithClaims(TheUser.RefreshToken, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	fmt.Println("this is the refresh token", refreshToken)
	if err == nil {
		refreshClaims, ok := refreshToken.Claims.(*domain.Claims)
		fmt.Println("this is the refresh claims", refreshClaims)

		if refreshClaims.ExpiresAt < time.Now().Unix() {

			_,err :=c.userUsecase.Update(TheUser.ID,domain.User{RefreshToken: ""})
			if err != nil{
				ctx.JSON(http.StatusInternalServerError,gin.H{"error":err})
			}
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "refresh token is expired"})
			return
		}
		if ok && refreshToken.Valid {
			if err != nil {
				ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err})
				return
			}
			// User login logic
			fmt.Println("this is the user", TheUser)
			fmt.Println("this is the NUID", NUID)
			if bcrypt.CompareHashAndPassword([]byte(TheUser.Password), []byte(NUID.Password)) != nil {
				ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "password is incorrect"})
				return
			}

			newToken, refresh, err := infrastructure.GenerateToken(&TheUser, NUID.Password)
			if err != nil {
				ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
			ctx.IndentedJSON(200, gin.H{"refreshed access token": newToken})
			TheUser.RefreshToken = refresh
			_,err = c.userUsecase.Update(TheUser.ID,domain.User{RefreshToken: refresh})
			if err != nil{
				ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err})
				return
			}
			
		} else {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
		}
	}
	ctx.IndentedJSON(http.StatusForbidden, gin.H{"error": "couldn't refrsh the token"})
}
