package handlers

import (
	"blogApp/internal/domain"
	"blogApp/internal/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase user.UserUseCaseInterface
}

func NewUserHandler(uc user.UserUseCaseInterface) *UserHandler {
	return &UserHandler{
		UserUsecase: uc,
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	var user *domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, token, err := h.UserUsecase.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"token": token,
	})
}

func (h *UserHandler) Register(c *gin.Context) {
	var user *domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}
	user, err := h.UserUsecase.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	})

}

func (h *UserHandler) GetUser(c *gin.Context){
	user, ok := c.Request.Context().Value("user").(*domain.User)

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	
	c.JSON(http.StatusOK, user)
}


func (h *UserHandler) DeleteUser(c *gin.Context){
	userID, exists := c.Request.Context().Value("id").(string)
	if !exists || userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

		
	if err := h.UserUsecase.DeleteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h *UserHandler) UpdateUser(c * gin.Context){
	var user domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
        return
    }

    err := h.UserUsecase.UpdateUser(&user)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}


