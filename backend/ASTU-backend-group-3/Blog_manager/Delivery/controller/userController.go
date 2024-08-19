package controller

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/Usecases"
	"ASTU-backend-group-3/Blog_manager/infrastructure"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// UserController handles user-related endpoints
type UserController struct {
	UserUsecase Usecases.UserUsecase
}

// NewUserController creates a new instance of UserController
func NewUserController(userUsecase Usecases.UserUsecase) *UserController {
	return &UserController{UserUsecase: userUsecase}
}

// Register handles user registration
func (uc *UserController) Register(c *gin.Context) {
	var input Domain.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := uc.UserUsecase.Register(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// UpdateUser handles updating user information
func (uc *UserController) UpdateUser(c *gin.Context) {
	usernameParam := c.Param("username")

	// Extract the username from the token (set by AuthMiddleware)
	usernameFromToken, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
		return
	}

	// Check if the user is updating their own profile
	if usernameParam != usernameFromToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own profile"})
		return
	}

	var input Domain.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := uc.UserUsecase.UpdateUser(usernameParam, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser handles deleting a user
func (uc *UserController) DeleteUser(c *gin.Context) {
	username := c.Param("username")

	err := uc.UserUsecase.DeleteUser(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// Login handles user login
func (uc *UserController) Login(c *gin.Context) {
	var input Domain.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	accessToken, err := uc.UserUsecase.Login(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

func (uc *UserController) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "refresh token not found"})
		return
	}
	var jwtKey = []byte("BlogManagerSecretKey")

	token, err := jwt.ParseWithClaims(refreshToken, &infrastructure.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify the token's signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.Abort()
		return
	}

	// Get the username from the token
	username, err := infrastructure.GetUsernameFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to get username from token"})
		c.Abort()
		return
	}

	// Set token claims in context
	claims, ok := token.Claims.(*infrastructure.Claims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
		return
	}
	claims.Username = username

	accessToken, err := infrastructure.GenerateJWT(claims.Username, claims.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

func (uc *UserController) ForgotPassword(c *gin.Context) {

	var input Domain.ForgetPasswordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := uc.UserUsecase.ForgotPassword(input.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
func (uc *UserController) ResetPassword(c *gin.Context) {
	reset_token := c.Param("token")

	new_token, err := uc.UserUsecase.Reset(reset_token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": new_token})

}

func (uc *UserController) ChangePassword(c *gin.Context) {
	var input Domain.ChangePasswordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := uc.UserUsecase.UpdatePassword(input.Username, input.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

func (uc *UserController) Logout(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	var token string
	parts := strings.Split(tokenString, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		token = parts[1]
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		return
	}

	err := uc.UserUsecase.Logout(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}

func (uc *UserController) PromoteToAdmin(c *gin.Context) {
	username := c.Param("username")

	err := uc.UserUsecase.PromoteTOAdmin(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin successfully"})
}

func (uc *UserController) Verify(c *gin.Context) {
	token := c.Param("token")
	err := uc.UserUsecase.Verify(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}
