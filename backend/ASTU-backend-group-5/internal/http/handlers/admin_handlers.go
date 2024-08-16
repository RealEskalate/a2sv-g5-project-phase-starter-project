package handlers

import (
	"blogApp/internal/domain"
	"blogApp/internal/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *UserHandler) DemoteUser(c *gin.Context){
	id, ok := c.Request.Context().Value("id").(string)

	if !ok{
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

func (h *UserHandler) PromoteToAdmin(c *gin.Context){
	id, ok := c.Request.Context().Value("id").(string)

	if !ok{
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not found"})
		return
	}
	err := h.UserUsecase.PromoteToAdmin(id)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err})
	}

	c.JSON(http.StatusOK, gin.H{"message" : "the user promoted to admin successfully!"})

}



