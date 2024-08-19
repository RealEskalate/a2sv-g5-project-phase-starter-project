package Controller

import (
    "github.com/gin-gonic/gin"
    "AAiT-backend-group-8/Domain"
)

type UserHandler struct {
    useCase Domain.IUserUseCase
}

func NewUserHandler(useCase Domain.IUserUseCase) *UserHandler {
    return &UserHandler{useCase: useCase}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
    var user Domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    err := h.useCase.RegisterUser(&user)
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

    err := h.useCase.VerifyEmail(token)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"message": "Email verified successfully"})
}


