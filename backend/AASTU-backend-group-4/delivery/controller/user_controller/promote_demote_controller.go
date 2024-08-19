package user_controller

import (
	"blog-api/domain/user"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) PromoteDemoteController(c *gin.Context) {
	var request user.PromoteDemoteRequest
	var user_ user.User

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Email != "" {
		if user_, err := uc.usecase.GetByEmail(c, request.Email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email doesn't exist.", "request": user_})
			return
		}
	} else if request.Username != "" {
		if user_, err := uc.usecase.GetByUsername(c, request.Username); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username doesn't exist.", "request": user_})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You must specify the username or email of the User."})
		return
	}

	if request.Action == "promote" {
		user_.IsAdmin = true
	} else if request.Action == "demote" {
		user_.IsAdmin = false
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Action must be specified (i.e promote or demote)"})
		return
	}

	// if err := uc.UpdateUser(c, user_.ID, user_); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Status updated successfully", "User": user_})
}
