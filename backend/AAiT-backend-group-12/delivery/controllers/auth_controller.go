package controllers

import (
	"blog_api/domain"
	"blog_api/domain/dtos"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	usecase domain.UserUsecaseInterface
}

func GetHTTPErrorCode(err domain.CodedError) int {
	switch err.GetCode() {
	case domain.ERR_BAD_REQUEST:
		return 400
	case domain.ERR_UNAUTHORIZED:
		return 401
	case domain.ERR_FORBIDDEN:
		return 403
	case domain.ERR_NOT_FOUND:
		return 404
	case domain.ERR_CONFLICT:
		return 409
	default:
		return 500
	}
}

func NewAuthController(usecase domain.UserUsecaseInterface) *AuthController {
	return &AuthController{usecase: usecase}
}

func (controller *AuthController) HandleSignup(c *gin.Context) {
	var newUser domain.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(400, domain.Response{"error": "Invalid input"})
		return
	}

	err := controller.usecase.Signup(c, &newUser)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(201, domain.Response{"message": "User created"})
}

func (controller *AuthController) HandleLogin(c *gin.Context) {
	var newUser domain.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(400, domain.Response{"error": "Invalid input"})
		return
	}

	acK, rfK, err := controller.usecase.Login(c, &newUser)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(201, domain.Response{"accessToken": acK, "refreshToken": rfK})
}

func (controller *AuthController) HandleRenewAccessToken(c *gin.Context) {
	// obtain token from the request header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, domain.Response{"error": "Authorization header not found"})
		return
	}

	headerSegments := strings.Split(authHeader, " ")
	if len(headerSegments) != 2 || strings.ToLower(headerSegments[0]) != "bearer" {
		c.JSON(401, domain.Response{"error": "Authorization header is invalid"})
		return
	}

	accessToken, err := controller.usecase.RenewAccessToken(c, headerSegments[1])
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(200, domain.Response{"accessToken": accessToken})
}

func (controller *AuthController) HandleUpdateUser(c *gin.Context) {
	reqUsername := strings.TrimSpace(c.Param("username"))
	if reqUsername == "" {
		c.JSON(400, domain.Response{"error": "Username is required"})
		return
	}

	var updatedUser dtos.UpdateUser
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(400, domain.Response{"error": "Invalid input"})
		return
	}

	tokenUsername, ok := c.Keys["username"]
	if !ok {
		c.JSON(400, domain.Response{"error": "Username not found in token"})
		return
	}

	resData, err := controller.usecase.UpdateUser(c, reqUsername, tokenUsername.(string), &updatedUser)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(200, domain.Response{"message": "User updated", "data": resData})
}

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
