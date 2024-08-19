package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/Usecases"
	
)

func Chat(c *gin.Context) {
	// take the message from the body

	var input *Domain.Chat
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	res , err := Usecases.Chat(input.Prompt)

	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	c.JSON(http.StatusOK, gin.H{"message": res})
	}

