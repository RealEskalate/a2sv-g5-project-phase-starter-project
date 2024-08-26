package user_controller

import (
	"blog-api/domain"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) PromoteDemote(c *gin.Context) {
	var request domain.PromoteDemoteRequest
	var user_ domain.User

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Email != "" {
		if user_, err := uc.userUsecase.GetByEmail(c, request.Email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email doesn't exist.", "request": user_})
			return
		}
	} else if request.Username != "" {
		if user_, err := uc.userUsecase.GetByUsername(c, request.Username); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username doesn't exist.", "request": user_})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You must specify the username or email of the User."})
		return
	}
	var action string
	if request.Action == "promote" {
		action = "promote"
	} else if request.Action == "demote" {
		action = "demote"
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Action must be specified (i.e promote or demote)"})
		return
	}

	if err := uc.userUsecase.PromoteDemote(c, user_.ID, action); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated successfully", "User": user_})
}
