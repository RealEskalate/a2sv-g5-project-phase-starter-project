package controllers

import (
	domain "blogs/Domain"

	"github.com/gin-gonic/gin"
)
type SignupController struct {
	SignupUsecase domain.SignupUseCase
	
}


func (s *SignupController) Signup(c *gin.Context) {	
	var user domain.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	response := s.SignupUsecase.Create(c , user)
	HandleResponse(c , response)
}
