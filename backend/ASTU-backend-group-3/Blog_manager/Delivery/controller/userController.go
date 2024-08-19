package controller

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/Usecases"
	"ASTU-backend-group-3/Blog_manager/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
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
	username := c.Param("username")

	var input Domain.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := uc.UserUsecase.UpdateUser(username, &input)
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

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

