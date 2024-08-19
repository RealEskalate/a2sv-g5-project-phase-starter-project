package controllers

import (
	"blogapp/Domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userController struct {
	userUseCase Domain.UserUseCases
}

func NewUserController(service_reference Domain.UserUseCases) *userController {
	return &userController{
		userUseCase: service_reference,
	}
}

func (uc *userController) GetUsers(c *gin.Context) {

	users, err, statusCode := uc.userUseCase.GetUsers(c)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(statusCode, gin.H{"users": users})
	}
}

func (uc *userController) GetUser(c *gin.Context) {
	cur_user, err := Getclaim(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	user, err, statusCode := uc.userUseCase.GetUsersById(c, objectID, *cur_user)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(statusCode, gin.H{"user": user})
	}
}

func (uc *userController) CreateUser(c *gin.Context) {
	var user Domain.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	v := validator.New()
	if err := v.Struct(user); err != nil {
		fmt.Printf(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid or missing data", "error": err.Error()})
		return
	}

	OmitedUser, err, statusCode := uc.userUseCase.CreateUser(c, &user)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(statusCode, gin.H{"user": OmitedUser})
	}
}

func (uc *userController) UpdateUser(c *gin.Context) {
	logeduser, err := Getclaim(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	var user Domain.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	OmitedUser, err, statusCode := uc.userUseCase.UpdateUsersById(c, objectID, user, *logeduser)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(statusCode, gin.H{"user": OmitedUser})
	}
}

func (uc *userController) DeleteUser(c *gin.Context) {
	user, err := Getclaim(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}

	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	err, statusCode := uc.userUseCase.DeleteUsersById(c, objectID, *user)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(statusCode, gin.H{"message": "User deleted successfully"})
	}
}

func (uc *userController) PromoteUser(c *gin.Context) {
	user, err := Getclaim(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	OmitedUser, err, statusCode := uc.userUseCase.PromoteUser(c, objectID, *user)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(statusCode, gin.H{"message": "User promoted successfully", "user": OmitedUser})
	}
}

func (uc *userController) DemoteUser(c *gin.Context) {
	user, err := Getclaim(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	OmitedUser, err, statusCode := uc.userUseCase.DemoteUser(c, objectID, *user)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(statusCode, gin.H{"message": "User demoted successfully", "user": OmitedUser})
	}
}
