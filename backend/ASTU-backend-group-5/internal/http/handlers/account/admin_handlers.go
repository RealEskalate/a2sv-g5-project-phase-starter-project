package account

import (
	"blogApp/internal/domain"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) DemoteUser(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not found"})
		return
	}
	err := h.UserUsecase.DemoteFromAdmin(id)
	fmt.Println(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "the user demoted to admin successfully!"})

}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	filter := domain.UserFilter{
		Username:   c.Query("username"),
		Email:      c.Query("email"),
		Role:       c.Query("role"),
		Gender:     c.Query("gender"),
		Profession: c.Query("profession"),
		Verified:   c.Query("verified"),
		OrderBy:    c.DefaultQuery("orderBy", "alphabet"),
	}

	users, err := h.UserUsecase.GetAllUsers(page, pageSize, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	baseUrl := c.Request.URL.Scheme + "://" + c.Request.Host + c.Request.URL.Path
	nextUrl := baseUrl + "?page=" + strconv.Itoa(page+1) + "&pageSize=" + strconv.Itoa(pageSize)
	prevUrl := ""
	if page > 1 {
		prevUrl = baseUrl + "?page=" + strconv.Itoa(page-1) + "&pageSize=" + strconv.Itoa(pageSize)
	}

	c.JSON(http.StatusOK, gin.H{
		"users":   users,
		"nextUrl": nextUrl,
		"prevUrl": prevUrl,
	})
}

func (h *UserHandler) AdminRemoveUser(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	if err := h.UserUsecase.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

}

func (h *UserHandler) PromoteToAdmin(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not found"})
		return
	}
	err := h.UserUsecase.PromoteToAdmin(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "the user promoted to admin successfully!"})

}

func (h *UserHandler) AdminGetUser(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not found"})
		return
	}
	user, err := h.UserUsecase.FindUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

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
	users, err := h.UserUsecase.FilterUsers(filters)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the filtered users
	c.JSON(http.StatusOK, users)
}
