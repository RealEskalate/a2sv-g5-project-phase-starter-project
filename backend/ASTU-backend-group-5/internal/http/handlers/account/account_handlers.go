package account

import (
	"blogApp/internal/domain"
	"blogApp/internal/usecase/user"
	"blogApp/pkg/checker"


	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	if err := checker.IsValidEmail(user.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := checker.IsValidPassword(user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

func (h *UserHandler) GetUser(c *gin.Context) {

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	userClaims, ok := claims.(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	dbUser, err := h.UserUsecase.FindUserById(userClaims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Me := domain.GetUserDTO{
		ID:      dbUser.ID,
		Email:   dbUser.Email,
		Role:    dbUser.Role,
		Profile: dbUser.Profile,
	}
	c.JSON(http.StatusOK, Me)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	userClaims, ok := claims.(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if err := h.UserUsecase.DeleteUser(userClaims.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}



func (h *UserHandler) UpdateUser(c *gin.Context) {
    var updateUserDTO domain.UpdateUserDTO
    if err := c.BindJSON(&updateUserDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    claims, exists := c.Get("claims")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }
    userClaims, ok := claims.(*domain.Claims)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    objectID, err := primitive.ObjectIDFromHex(userClaims.UserID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
        return
    }

    dbUser := domain.User{
        ID:       objectID,
        UserName: updateUserDTO.UserName,
        Email:    updateUserDTO.Email,
        Profile:  updateUserDTO.Profile,
    }

    err = h.UserUsecase.UpdateUser(&dbUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}


func (h *UserHandler) GetAnyUser(c *gin.Context) {
	userId := c.Param("userId")
	user, err := h.UserUsecase.FindUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	GetUser := domain.GetUserDTO{
		UserName: user.UserName,
		Email:    user.Email,
		Role:     user.Role,
		Profile:  user.Profile,
	}
	c.JSON(http.StatusOK, GetUser)
}
