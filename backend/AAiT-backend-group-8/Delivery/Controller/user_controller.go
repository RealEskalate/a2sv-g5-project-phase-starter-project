package Controller

import (
	"AAiT-backend-group-8/Domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase Domain.IUserUseCase
}

func NewUserHandler(userUseCase Domain.IUserUseCase) *UserHandler {
	return &UserHandler{UserUsecase: userUseCase}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user Domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.UserUsecase.RegisterUser(&user)
	if err != nil {
		if err.Error() == "email already exists" {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(200, gin.H{"message": "Registration successful. Check your email for verification link."})
}

func (h *UserHandler) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(400, gin.H{"error": "Invalid token"})
		return
	}

	err := h.UserUsecase.VerifyEmail(token)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Email verified successfully"})
}

func (h *UserHandler) Login(c *gin.Context) {
    // Corrected struct with exported fields
    type EmailPass struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    var ep EmailPass

    // Bind the JSON payload to the struct
    bindErr := c.BindJSON(&ep)
    if bindErr != nil {
        c.IndentedJSON(400, gin.H{"message": "invalid request payload"})
        return
    }

    // Debugging line to print the payload (c.Request.Body is already consumed by BindJSON)
    fmt.Printf("Received payload: %v\n", ep)

    // Perform login using the extracted email and password
    token, refresher, err := h.UserUsecase.Login(ep.Email, ep.Password)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Respond with token and refresher
    c.IndentedJSON(http.StatusOK, gin.H{"token": token, "refresher": refresher})
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	var cred Domain.Credential

	bind_err := c.BindJSON(&cred)
	if bind_err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid request payload"})
		return
	}

	token, err := h.UserUsecase.RefreshToken(cred.Email, cred.Refresher)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid refresh token "})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"token": token})

}
