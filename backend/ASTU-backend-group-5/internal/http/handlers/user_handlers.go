package handlers

import (
	"blogApp/internal/domain"
	"blogApp/internal/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
    loginUseCase *user.LoginUseCase
}



func NewLoginHandler(loginUseCase *user.LoginUseCase) *LoginHandler {
    return &LoginHandler{loginUseCase: loginUseCase}
}

func (h *LoginHandler) Login(c *gin.Context) {
    var user *domain.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    user, token, err := h.loginUseCase.Login(user.Email, user.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "id":       user.ID,
        "email": 	user.Email,
        "role":     user.Role,
        "token":    token,
    })
}