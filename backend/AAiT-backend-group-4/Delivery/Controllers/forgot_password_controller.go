package controllers

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordController struct {
	OtpService            domain.OtpInfrastructure
	ForgotPasswordUsecase domain.ForgotPasswordUsecase
	Env                   *bootstrap.Env
}

// ForgotPassword is a method of the ForgotPasswordController struct that handles the forgot password functionality.
// It receives a gin.Context object and a request map containing the email as input parameters.
// The method binds the JSON request to the request map and returns an error if the binding fails.
// It then calls the ForgotPassword method of the ForgotPasswordUsecase, passing the gin.Context, email, and EmailApiKey as arguments.
// If an error occurs during the process, it returns the error as a JSON response.
// Otherwise, it returns the response as a JSON response with HTTP status OK.
func (fc *ForgotPasswordController) ForgotPassword(c *gin.Context) {
	var request map[string]string

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := fc.ForgotPasswordUsecase.ForgotPassword(c, request["email"], fc.Env.EmailApiKey)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ServePage serves the password reset page.
//
// It takes a gin.Context object as a parameter and returns no values.
// The function renders the "password_reset.html" template and sends it as a response with a status code of http.StatusOK.
func (fc *ForgotPasswordController) ServePage(c *gin.Context) {
	c.HTML(http.StatusOK, "password_reset.html", nil)
}

// VerifyForgotPassowrd is a method of the ForgotPasswordController struct that handles the verification of a forgot password request.
// It takes a gin.Context object as a parameter and returns the response as JSON.
// The method first checks if the email exists in the context, and if not, it returns an error response.
// Then, it binds the request body to the passwordChangeForm variable and returns an error response if the binding fails.
// Next, it asserts the email value from the context to a string and returns an error response if the assertion fails.
// Finally, it calls the VerifyChangePassword method of the ForgotPasswordUsecase, passing the email, passwordChangeForm, and context as parameters.
// If an error occurs during the verification process, it returns an error response, otherwise, it returns the response as JSON.
func (fc *ForgotPasswordController) VerifyForgotPassowrd(c *gin.Context) {
	var passwordChangeForm domain.ForgotPasswordRequest

	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not found in context"})
		return
	}

	if err := c.ShouldBind(&passwordChangeForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emailStr, ok := email.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Email type assertion failed"})
		return
	}

	response, err := fc.ForgotPasswordUsecase.VerifyChangePassword(c, emailStr, passwordChangeForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

}
