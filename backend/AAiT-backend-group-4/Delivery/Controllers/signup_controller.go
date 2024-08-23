package controllers

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupRequestWrapper struct {
	Request domain.SignupRequest `json:"request"`
}

type SignupController struct {
	SingupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) SignUp(c *gin.Context) {

	var request domain.SignupRequest

	imagePath, ok := c.Get("profileImagePath")
	log.Printf("imagePath: %v", imagePath)

	// Extract and log the 'request' form field
	jsonData := c.Request.FormValue("request")
	if jsonData == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'request' form field"})
		return
	}

	if ok {
		request.Image_Path = imagePath.(string)
	}

	// Unmarshal JSON data into the struct
	err := json.Unmarshal([]byte(jsonData), &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	request.Image_Path = imagePath.(string)
	log.Println("Request:", request)

	otpResponse, err := sc.SingupUsecase.Signup(c, &request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, otpResponse)
}
