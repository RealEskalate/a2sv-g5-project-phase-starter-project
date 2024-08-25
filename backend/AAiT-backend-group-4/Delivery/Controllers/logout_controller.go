package controllers

import (
	domain "aait-backend-group4/Domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// LogoutController handles the logout functionality for users.
type LogoutController struct {
	LogoutUsecase domain.LogoutUsecase
}

// NewLogoutController creates a new instance of LogoutController with the given LogoutUsecase.
func NewLogoutController(logoutUsecase domain.LogoutUsecase) *LogoutController {
	return &LogoutController{
		LogoutUsecase: logoutUsecase,
	}
}

// Logout handles user logout requests.
// It expects an Authorization token in the request header.
// If the token is missing, it responds with an Unauthorized status.
// If the token is provided, it calls the LogoutUsecase to process the logout.
// The result of the logout operation is then returned in the response.
func (c *LogoutController) Logout(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization") // Retrieve the Authorization token from the header
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization token required"})
		return
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization header"})
		return
	}

	tokens := strings.Split(authParts[1], ":")
	if len(tokens) != 2 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Both access and refresh tokens are required"})
		return
	}

	accessToken := tokens[0]

	// Call the LogoutUsecase to handle the logout process
	response, err := c.LogoutUsecase.Logout(ctx, accessToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Respond with the result of the logout operation
	ctx.JSON(http.StatusOK, response)
}
