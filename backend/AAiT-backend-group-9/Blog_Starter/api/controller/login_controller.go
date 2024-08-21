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
	
	 ctx := c.Request.Context()
	loginResponse,err := lc.LoginUsecase.Login(ctx, &request)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}
	c.JSON(http.StatusOK, loginResponse)

}

func(lc *LoginController) ForgotPassword(c *gin.Context){

	var request domain.ForgotPasswordRequest
	err:= c.BindJSON(&request)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad email request"})
		return
	}
	request.Email = strings.ToLower(request.Email)
	user, err := lc.UserUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{"error": "user not found with given email"})
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
	c.JSON(http.StatusOK, gin.H{"success": "Message sent successfully"})


}

func (lc *LoginController) UpdatePassword(c *gin.Context) {
    var request domain.ChangePasswordRequest
    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request to update password"})
        return
    }

    userResponse, err := lc.UserUsecase.GetUserByEmail(c, request.Email)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    userID := userResponse.UserID.String()

    // Convert gin.Context to standard context.Context
    ctx := c.Request.Context()

    if err := lc.LoginUsecase.UpdatePassword(ctx, request, userID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"success": "Password updated successfully"})
}



