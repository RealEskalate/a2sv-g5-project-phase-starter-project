package controllers

import (
	domain "blogs/Domain"

	"github.com/gin-gonic/gin"
)

type NewUserController struct {
	UserUsecase domain.UserUseCase
}


func (s *NewUserController) UpdateUser(c *gin.Context) {
	var user domain.UserUpdateRequest
	userid := c.Param("id")
	authenticatedUserID := c.GetString("user_id")
	role := c.GetString("role")

	if userid != authenticatedUserID && role != "admin" {
		c.JSON(401, gin.H{"error": "Unauthorized to update user Information"})
		return
	}

	user.ID = authenticatedUserID
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := s.UserUsecase.UpdateUser(c, user)
	HandleResponse(c, response)

}

func (s *NewUserController) PromoteUser(c *gin.Context) {

	role := c.GetString("user_id")

	var promotion domain.UserPromotionRequest

	err := c.ShouldBindJSON(&promotion)

	if err != nil { 
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userid := c.Param("id")

	rolecheck , err := s.UserUsecase.FindUserByID(c, role)

	if rolecheck.Role != "admin" {
		c.JSON(401, gin.H{"error": "Unauthorized to promote user"})
		return
	}

	response := s.UserUsecase.PromoteandDemoteUser(c, userid , promotion)

	HandleResponse(c, response)

}


