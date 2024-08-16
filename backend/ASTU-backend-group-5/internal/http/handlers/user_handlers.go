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

	user, ok := c.Get("user").(*domain.User)

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	dbUser, err := h.UserUsecase.GetUser(user.Id)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, dbUser)
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

package handler

import (
    "net/http"
    "your_project/usecase"
    "github.com/gin-gonic/gin"
)

func (h *UserHandler) FilterUsers(c *gin.Context) {
    // Define valid filters
    validFilters := []string{"role", "email", "username", "firstName", "lastName"}

    // Create a map to hold the filters
    filters := make(map[string]interface{})

    // Loop through query parameters and add valid ones to the filters map
    for _, key := range validFilters {
        if value := c.Query(key); value != "" {
            filters[key] = value
        }
    }

    // Call the usecase to filter users based on the provided filters
    users, err := h.usecase.FilterUsers(filters)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Return the filtered users
    c.JSON(http.StatusOK, users)
}
