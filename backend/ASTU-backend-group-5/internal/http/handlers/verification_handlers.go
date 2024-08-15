package handlers

import (
	"blogApp/internal/usecase/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	RequestEmailVerification(user user.User) error
	RequestPasswordResetUsecase(userEmail string) error
	ResetPassword(token string, password string, email string) error
	VerifyEmail(token string, email string) error
}

type UserHandler struct {
	Usecase UserUsecase
}

func NewUserHandler(uc UserUsecase) *UserHandler {
	return &UserHandler{
		Usecase: uc,
	}
}

func (h *UserHandler) RequestVerifyEmail(c *gin.Context) {
	var request struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user := user.User{
		Email: request.Email,
	}

	err := h.Usecase.RequestEmailVerification(user)
	if err != nil {
		log.Printf("Error sending verification email: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Verification email sent successfully"})
}

func (h *UserHandler) VerifyEmail(c *gin.Context) {
	email := c.Query("email")
	token := c.Query("token")
	if email == "" || token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing email or token"})
		return
	}

	err := h.Usecase.VerifyEmail(token, email)
	if err != nil {
		log.Printf("Error verifying email: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}

func (h *UserHandler) ResetPasswordRequest(c *gin.Context) {
	var request struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	err := h.Usecase.RequestPasswordResetUsecase(request.Email)
	if err != nil {
		log.Printf("Error sending password reset email: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send password reset email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password reset email sent successfully"})
}

func (h *UserHandler) ResetPassword(c *gin.Context) {
	var request struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		Token    string `json:"token" binding:"required"`
	}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err = h.Usecase.ResetPassword(request.Token, request.Password, request.Email)
	if err != nil {
		log.Printf("Error resetting password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})

}
