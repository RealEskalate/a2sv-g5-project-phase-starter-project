package controllers

import (
	"astu-backend-g1/config"
	"astu-backend-g1/domain"
	"astu-backend-g1/infrastructure"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

// Register godoc
// @Summary      Register a new user
// @Description  Creates a new user account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body domain.User true "User info"
// @Success      201 {object} map[string]string "message"
// @Failure      406 {object} map[string]string "error"
// @Failure      500 {object} map[string]string "error"
// @Router       /users/register [post]
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

// AccountVerification godoc
// @Summary      Verify account
// @Description  Verifies a user account using email and token
// @Tags         users
// @Produce      json
// @Param        email query string true "Email"
// @Param        token query string true "Verification token"
// @Success      200 {object} map[string]string "message"
// @Failure      406 {object} map[string]string "error"
// @Router       /users/accountVerification [get]
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

// ForgetPassword godoc
// @Summary      Forget password
// @Description  Sends a password reset token to the user's email
// @Tags         users
// @Produce      json
// @Param        email query string true "Email"
// @Success      200 {object} map[string]string "message"
// @Failure      406 {object} map[string]string "error"
// @Router       /users/forgetPassword [get]
func (c *UserController) ForgetPassword(ctx *gin.Context) {
	email := ctx.Query("email")
	_, err := c.userUsecase.ForgetPassword(email)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Password reset token sent to your email"})
}

// ResetPassword godoc
// @Summary      Reset password
// @Description  Resets the user's password using the token
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        email query string true "Email"
// @Param        token query string true "Reset token"
// @Param        newPassword body object true "New password info"
// @Success      200 {object} map[string]string "message"
// @Failure      406 {object} map[string]string "error"
// @Router       /users/resetPassword [post]
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

// ForgetPasswordUser godoc
// @Summary      Forget password (alternative)
// @Description  Sends a password reset token to the user's email (alternative method)
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        email body string true "Email"
// @Success      200 {object} map[string]string "message"
// @Failure      406 {object} map[string]string "error"
// @Router       /users/forgetPassword [post]
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

// LoginUser godoc
// @Summary      Login user
// @Description  Authenticates a user and returns a JWT access token
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body domain.User true "Login info"
// @Success      200 {object} map[string]string "access_token"
// @Failure      406 {object} map[string]string "error"
// @Router       /users/login [post]
func (c *UserController) LoginUser(ctx *gin.Context) {
	user := &domain.User{}
	if err := ctx.ShouldBind(user); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	access_token, err := c.userUsecase.LoginUser(user.Username, user.Password,user.Email)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"access_token": access_token})
}

// GetUsers godoc
// @Summary      Get users
// @Description  Retrieves users by optional username or email query
// @Tags         users
// @Produce      json
// @Param        username query string false "Username"
// @Param        email query string false "Email"
// @Success      200 {array} domain.User "List of users"
// @Failure      404 {object} map[string]string "error"
// @Router       /users [get]
// logout user
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

// GetUserByID godoc
// @Summary      Get user by ID
// @Description  Retrieves a user by their ID
// @Tags         users
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200 {object} domain.User "User data"
// @Failure      404 {object} map[string]string "error"
// @Router       /users/{id} [get]
func (c *UserController) GetUserByID(ctx *gin.Context) {
	userID := ctx.Param("id")
	user, err := c.userUsecase.GetByID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Deletes a user by their ID
// @Tags         users
// @Produce      json
// @Param        id path string true "User ID"
// @Success      204 "No content"
// @Failure      406 {object} map[string]string "error"
// @Router       /users/{id} [delete]
func (c *UserController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	err := c.userUsecase.Delete(userId)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"message":"user deleted from database"})
}

// UpdateUser godoc
// @Summary      Update user
// @Description  Updates user information by their ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Param        user body domain.User true "Updated user info"
// @Success      200 {object} domain.User "Updated user data"
// @Failure      406 {object} map[string]string "error"
// @Failure      500 {object} map[string]string "error"
// @Router       /users/{id} [put]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	updateData := domain.User{}
	if err := ctx.ShouldBind(&updateData); err != nil {
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

// RefreshAccessToken godoc
// @Summary      Refresh access token
// @Description  Refreshes the access token using the refresh token
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        uid path string true "User ID"
// @Param        pwd body object true "Password"
// @Success      200 {object} map[string]string "refreshed access token"
// @Failure      401 {object} map[string]string "error"
// @Failure      403 {object} map[string]string "error"
// @Failure      500 {object} map[string]string "error"
// @Router       /users/{uid}/refresh [post]
func (c *UserController) RefreshAccessToken(ctx *gin.Context) {
	configJwt, err := config.LoadConfig()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	jwtSecret := []byte(configJwt.Jwt.JwtKey)
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

			_, err := c.userUsecase.Update(TheUser.ID, domain.User{RefreshToken: ""})
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
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
			_, err = c.userUsecase.Update(TheUser.ID, domain.User{RefreshToken: refresh})
			if err != nil {
				ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err})
				return
			}

		} else {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
		}
	}
	ctx.IndentedJSON(http.StatusForbidden, gin.H{"error": "couldn't refrsh the token"})
}
func (c *UserController) Promote(ctx *gin.Context) {
	username := ctx.Param("username")
	_, err := c.userUsecase.PromteUser(username)
	log.Println(err)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}
func (c *UserController) Demote(ctx *gin.Context) {
	username := ctx.Param("username")
	_, err := c.userUsecase.DemoteUser(username)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User demoted to user"})
}
func (c *UserController) PromoteByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	_, err := c.userUsecase.PromteUserByEmail(email)
	log.Println(err)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}
func (c *UserController) DemoteByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	_, err := c.userUsecase.DemoteUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User demoted to user"})
}
