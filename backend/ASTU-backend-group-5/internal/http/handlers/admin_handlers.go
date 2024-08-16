package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *UserHandler) DemoteUser(c *gin.Context){
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not found"})
		return
	}
	err := h.UserUsecase.DemoteFromAdmin(id)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err})
	}

	c.JSON(http.StatusOK, gin.H{"message" : "the user demoted to admin successfully!"})

}

func (h *UserHandler) GetAllUsers(c *gin.Context){
	users, err := h.UserUsecase.GetAllUsers()
	
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
        return
    }

    c.JSON(http.StatusOK, users)
}

func (h *UserHandler) AdminRemoveUser(c *gin.Context){
	id :=c.Param("id")

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

func (h *UserHandler) PromoteToAdmin(c *gin.Context){
	id :=c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not found"})
		return
	}
	err := h.UserUsecase.PromoteToAdmin(id)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err})
	}

	c.JSON(http.StatusOK, gin.H{"message" : "the user promoted to admin successfully!"})

}

func(h *UserHandler) AdminGetUser(c *gin.Context){
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




