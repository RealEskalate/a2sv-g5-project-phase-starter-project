package user_controller

import (
	"blog-api/domain/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UserController) SignupController(c *gin.Context) {
	var request user.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if err = infrastructure.ValidateEmail(request.Password); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// if err = infrastructure.ValidatePassword(request.Password); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	//
	// if _, err = uc.usecase.GetByEmail(c, request.Email); err != nil{
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already taken"})
	// }
	//
	// if _, err = uc.usecase.GetByUsername(c, request.Username); err != nil{
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Username is already taken"})
	// }

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid password!"})
		return
	}

	request.Password = string(encryptedPassword)

	user := &user.User{
		ID:        primitive.NewObjectID(),
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Username:  request.Username,
		Email:     request.Email,
		Password:  request.Password,
		Role:      "User",
		Active:    false,
	}

	err = uc.usecase.SignupUsecase(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "user signed up successfully"})
}
