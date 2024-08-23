package user_controller

import (
	"blog-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) SignUp(c *gin.Context) {
	var req domain.SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := uc.userUsecase.SignUp(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
