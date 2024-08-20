package controllers

import (
	"blog_api/domain"
	"blog_api/domain/dtos"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	usecase domain.UserUsecaseInterface
}

// Returns the HTTP equivalent codes of the domain error codes
func GetHTTPErrorCode(err domain.CodedError) int {
	switch err.GetCode() {
	case domain.ERR_BAD_REQUEST:
		return http.StatusBadRequest
	case domain.ERR_UNAUTHORIZED:
		return http.StatusUnauthorized
	case domain.ERR_FORBIDDEN:
		return http.StatusForbidden
	case domain.ERR_NOT_FOUND:
		return http.StatusNotFound
	case domain.ERR_CONFLICT:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

// NewAuthController initializes the Auth controller
func NewAuthController(usecase domain.UserUsecaseInterface) *AuthController {
	return &AuthController{usecase: usecase}
}

// Returns the contents of the Authorization header
func (controller *AuthController) GetAuthHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header not found")
	}

	headerSegments := strings.Split(authHeader, " ")
	if len(headerSegments) != 2 || strings.ToLower(headerSegments[0]) != "bearer" {
		return "", fmt.Errorf("authorization header is invalid")
	}

	return headerSegments[1], nil
}

// Returns the domain from the context of the app
func (controller *AuthController) GetDomain(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	host := c.Request.Host
	return fmt.Sprintf("%s://%s", scheme, host)
}

// HandleRegister handles the register user endpoint
func (controller *AuthController) HandleSignup(c *gin.Context) {
	var newUser domain.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": "Invalid input"})
		return
	}

	err := controller.usecase.Signup(c, &newUser, controller.GetDomain(c))
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(201, domain.Response{"message": "User created. Please verify your email."})
}

// HandleLogin handles the login endpoint
func (controller *AuthController) HandleLogin(c *gin.Context) {
	var newUser domain.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": "Invalid input"})
		return
	}

	acK, rfK, err := controller.usecase.Login(c, &newUser)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(201, domain.Response{"accessToken": acK, "refreshToken": rfK})
}

// HandleRenewAccessToken handles the renew access token endpoint
func (controller *AuthController) HandleRenewAccessToken(c *gin.Context) {
	token, gErr := controller.GetAuthHeader(c)
	if gErr != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{"error": gErr.Error()})
		return
	}

	accessToken, err := controller.usecase.RenewAccessToken(c, token)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(200, domain.Response{"accessToken": accessToken})
}

// HandleUpdateUser handles the update user endpoint
func (controller *AuthController) HandleUpdateUser(c *gin.Context) {
	reqUsername := strings.TrimSpace(c.Param("username"))
	if reqUsername == "" {
		c.JSON(http.StatusBadRequest, domain.Response{"error": "Username is required"})
		return
	}

	var updatedUser dtos.UpdateUser
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": "Invalid input"})
		return
	}

	tokenUsername, ok := c.Keys["username"]
	if !ok {
		c.JSON(http.StatusBadRequest, domain.Response{"error": "Username not found in token"})
		return
	}

	resData, err := controller.usecase.UpdateUser(c, reqUsername, tokenUsername.(string), &updatedUser)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(200, domain.Response{"message": "User updated", "data": resData})
}

// HandlePromoteUser handles the promote user endpoint
func (controller *AuthController) HandlePromoteUser(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(400, domain.Response{"error": "Username is required"})
		return
	}

	err := controller.usecase.PromoteUser(c, username)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(200, domain.Response{"message": "User promoted"})
}

// HandleDemoteUser handles the demote user endpoint
func (controller *AuthController) HandleDemoteUser(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(400, domain.Response{"error": "Username is required"})
		return
	}

	err := controller.usecase.DemoteUser(c, username)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(200, domain.Response{"message": "User demoted"})
}

// HandleVerifyEmail handles the verify email endpoint
func (controller *AuthController) HandleVerifyEmail(c *gin.Context) {
	username := c.Param("username")
	token := c.Param("token")
	if username == "" || token == "" {
		c.JSON(400, domain.Response{"error": "Username and token are required"})
		return
	}

	err := controller.usecase.VerifyEmail(c, username, token, controller.GetDomain(c))
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(200, domain.Response{"message": "User verified"})
}

// HandleInitResetPassword handles the init reset password endpoint
func (controller *AuthController) HandleInitResetPassword(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, domain.Response{"error": "Invalid input " + err.Error()})
		return
	}

	uErr := controller.usecase.InitResetPassword(c, user.Username, user.Email, controller.GetDomain(c))
	if uErr != nil {
		c.JSON(GetHTTPErrorCode(uErr), domain.Response{"error": uErr.Error()})
		return
	}

	c.JSON(200, domain.Response{"message": "A reset password token has been sent to your email."})
}

// HandleResetPassword handles the reset password endpoint
func (controller *AuthController) HandleResetPassword(c *gin.Context) {
	var resetData dtos.ResetPassword
	token, err := controller.GetAuthHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{"error": err.Error()})
		return
	}

	if err = c.BindJSON(&resetData); err != nil {
		c.JSON(400, domain.Response{"error": "Invalid input " + err.Error()})
		return
	}

	uErr := controller.usecase.ResetPassword(c, resetData, token)
	if uErr != nil {
		c.JSON(GetHTTPErrorCode(uErr), domain.Response{"error": uErr.Error()})
		return
	}

	c.JSON(200, domain.Response{"message": "Password reset successful"})
}

// HandleLogout handles the logout endpoint
func (controller *AuthController) HandleLogout(c *gin.Context) {
	authHeader, err := controller.GetAuthHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{"error": err.Error()})
		return
	}

	controller.usecase.Logout(c, c.Keys["username"].(string), authHeader)
}

// HandleGoogleAuth handles the Google OAuth login endpoint
func (controller *AuthController) HandleGoogleAuth(c *gin.Context) {
	var response dtos.GoogleResponse
	if err := c.ShouldBindJSON(&response); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	acK, rfK, err := controller.usecase.GoogleOAuthAccess(c, &response)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(201, domain.Response{"accessToken": acK, "refreshToken": rfK})
}
