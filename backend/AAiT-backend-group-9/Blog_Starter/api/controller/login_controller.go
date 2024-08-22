package controller

import (
	"Blog_Starter/domain"
	"Blog_Starter/utils"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	UserUsecase  domain.UserUsecase
	OtpUsecase domain.OtpUsecase
}


func NewLoginController(LoginUsecase domain.LoginUsecase, OtpUsecase domain.OtpUsecase, UserUsecase  domain.UserUsecase) *LoginController {
	return &LoginController{
		LoginUsecase: LoginUsecase,
		UserUsecase:  UserUsecase,
		OtpUsecase: OtpUsecase,
	}
}

func (lc *LoginController) Login(c *gin.Context){
	var request domain.UserLogin
	
	err:= c.BindJSON(&request)
	if err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid request"})
		return

	}


	request.Email = strings.ToLower(request.Email)
	
	loginResponse,err := lc.LoginUsecase.Login(c, &request)
	if err!=nil{
		if err.Error() == "mongo: no documents in result"{
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}
	c.JSON(http.StatusOK, loginResponse)

}



func(lc *LoginController) ForgotPassword(c *gin.Context){

	var request domain.ForgotPasswordRequest
	err:= c.BindJSON(&request)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	request.Email = strings.ToLower(request.Email)
	err = domain.ValidateEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := lc.UserUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{"error": "user not found with given email"})
		return
	}
	
	if !user.IsActivated{
		c.JSON(http.StatusBadRequest, gin.H{"error": "user is not activated, Verify your email"})
		return
	}


	// Generate a random number between 0 and 9999 (inclusive).
	randNumber := rand.Intn(10000)

	// Format the code as a 4-digit string with leading zeros.
	code := fmt.Sprintf("%04d", randNumber)

	otp := domain.Otp{
		Email:      user.Email,
		Otp:        code,
		Expiration: time.Now().Add(5 * time.Minute),
	}

	oldOtp, err := lc.OtpUsecase.GetOtpByEmail(c, request.Email)
	if err == nil {
		otp.ID = oldOtp.ID
	} else {
		otp.ID = primitive.NewObjectID()
	}

	// Save OTP to database
	err = lc.OtpUsecase.SaveOtp(c, &otp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "couldnt save otp"})
		return
	}
	// TODO send email to user 
	emailContent := `
	<p>reset your password. please insert the following code in the required field to reset your password:</p>
	<h3>` + code + `</h3>
	<p><strong>This verification code is valid for 5 minutes.</strong> Please enter it on the reset Password page to proceed.</p>
	<p>If you did not sign up for the BlogApp, please ignore this email.</p>`
	// Create the email subject
	emailSubject := "Reset your password "

	// Generate the email body using the template function
	emailBody := utils.GenerateEmailTemplate("Reset Password ", emailContent)
	// Create the email template
	emailTemplate := domain.EmailTemplate{
		Subject: emailSubject,
		Body:    emailBody,
	}
	err = utils.SendTestEmail(request.Email, emailTemplate.Subject, emailTemplate.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "password reset otp sent successfully"})


}

func (lc *LoginController) UpdatePassword(c *gin.Context) {
    var request domain.ChangePasswordRequest
    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request to update password"})
        return
    }
	request.Email = strings.ToLower(request.Email)
	err := request.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    userResponse, err := lc.UserUsecase.GetUserByEmail(c, request.Email)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
	otp, err := lc.OtpUsecase.GetOtpByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No OTP requested with the given email"})
		return
	}

	if request.OTP != otp.Otp {
		// Otp from request doesn't match stored Otp or Otp has already been used
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Otp"})
		return
	} else if time.Now().After(otp.Expiration) {
		// Otp is correct but has expired
		c.JSON(http.StatusForbidden, gin.H{"error": "Otp Expired"})
		return
	}

    userID := userResponse.UserID.Hex()

    // Convert gin.Context to standard context.Context

    if err := lc.LoginUsecase.UpdatePassword(c, request, userID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"success": "Password updated successfully"})
}




func (lc *LoginController) LogOut(c *gin.Context) {
	
	user,err := utils.CheckUser(c)
	if err!=nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
		return
	}
	userID := user.UserID
	err = lc.LoginUsecase.LogOut(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}


