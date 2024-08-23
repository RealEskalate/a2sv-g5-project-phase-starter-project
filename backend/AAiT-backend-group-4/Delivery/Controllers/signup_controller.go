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

// SignUp handles the signup request and processes the data sent by the client.
// It expects a JSON payload containing the signup request data.
// The 'request' form field must be present in the request.
// If the JSON data is valid, it unmarshals it into the domain.SignupRequest struct.
// The 'profileImagePath' value is extracted from the context and assigned to the 'Image_Path' field of the request.
// The signup request is then passed to the SignupUsecase for further processing.
// If any errors occur during the signup process, an appropriate error response is sent back to the client.
// If the signup is successful, the response containing the OTP (one-time password) is sent back to the client.
func (sc *SignupController) SignUp(c *gin.Context) {

	var request domain.SignupRequest

	imagePath, ok := c.Get("profileImagePath")

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
