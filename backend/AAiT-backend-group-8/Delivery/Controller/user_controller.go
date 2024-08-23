package controller

import (
	"AAiT-backend-group-8/Domain"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var struct_validator = validator.New()

func (controller *Controller) RegisterUser(c *gin.Context) {
	var user Domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	struct_validator.RegisterValidation("password", passwordValidation)
	struct_err := struct_validator.Struct(user)

	if struct_err != nil {
		c.JSON(400, gin.H{"error": struct_err.Error()})
		return
	}

	err := controller.UserUseCase.RegisterUser(&user)
	if err != nil {
		if err.Error() == "email already exists" {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(200, gin.H{"message": "Registration successful. Check your email for verification link."})
}

func (controller *Controller) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(400, gin.H{"error": "Invalid token"})
		return
	}
	err := controller.UserUseCase.VerifyEmail(token)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Email verified successfully"})
}

func (controller *Controller) Login(c *gin.Context) {
	// Corrected struct with exported fields
	type EmailPass struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var ep EmailPass

	if err := c.ShouldBindJSON(&ep); err != nil {
		c.JSON(400, gin.H{"message": "invalid request payload"})
		return
	}

	token, refresher, err := controller.UserUseCase.Login(ep.Email, ep.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "refresher": refresher})
}

func (controller *Controller) RefreshToken(c *gin.Context) {
	var cred Domain.Credential

	bindErr := c.BindJSON(&cred)
	if bindErr != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid request payload"})
		return
	}

	token, err := controller.UserUseCase.RefreshToken(cred.Email, cred.Refresher)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"token": token})

}

func (controller *Controller) ForgotPassword(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(400, gin.H{"error": "Invalid email"})
		return
	}

	err := controller.UserUseCase.GenerateResetPasswordToken(email)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(404, gin.H{"error": "User not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(200, gin.H{"message": "Password reset email sent"})
}

func (controller *Controller) StoreToken(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(400, gin.H{"error": "Invalid token"})
		return
	}

	err := controller.UserUseCase.StoreToken(token)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Token stored successfully. You can now reset your password."})
}

func (controller *Controller) ResetPassword(c *gin.Context) {
	var payload struct {
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	err := controller.UserUseCase.ResetPassword(payload.Token, payload.NewPassword)
	if err != nil {
		if err.Error() == "invalid or expired token" || err.Error() == "invalid token payload" || err.Error() == "invalid or mismatched token" {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(200, gin.H{"message": "Password reset successful"})
}

func (controller *Controller) PromoteUser(c *gin.Context) {
	email := c.Param("email")

	err := controller.UserUseCase.PromoteUser(email)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User promoted successfully"})
}

func (controller *Controller) DemoteUser(c *gin.Context) {
	email := c.Param("email")

	err := controller.UserUseCase.DemoteUser(email)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User demoted successfully"})
}

func (controller *Controller) DeleteUser(c *gin.Context) {
	email := c.Param("email")

	err := controller.UserUseCase.DeleteUser(email)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (controller *Controller) Logout(c *gin.Context) {
	var cred Domain.Credential

	err := c.ShouldBindJSON(cred)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	err = controller.UserUseCase.Logout(cred.Email, cred.Refresher)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "email not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "refresher deleted successfully"})
}
