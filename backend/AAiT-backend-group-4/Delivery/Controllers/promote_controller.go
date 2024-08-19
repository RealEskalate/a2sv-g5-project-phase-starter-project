package controllers

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PromoteController struct {
	UserUsecase domain.UserUsecase
	Env         *bootstrap.Env
}

// Promote is a method of the PromoteController struct that handles the promotion of a user.
// It takes a gin.Context object as a parameter and retrieves the user ID from the request parameters.
// If the user ID is empty, it returns a JSON response with a bad request error.
// Otherwise, it calls the Promote method of the UserUsecase to promote the user.
// If an error occurs during the promotion process, it returns a JSON response with the error message.
// Finally, it constructs a JSON response with a success message and the promoted user, and returns it.
func (pc *PromoteController) Promote(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := pc.UserUsecase.Promote(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response = make(map[string]interface{})
	response["message"] = "User promoted successfully"
	response["user"] = user

	c.JSON(http.StatusOK, response)
}
