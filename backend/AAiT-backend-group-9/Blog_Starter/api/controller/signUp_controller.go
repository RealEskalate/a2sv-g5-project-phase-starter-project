package controller

import (
	"Blog_Starter/domain"
	EmailUtil "Blog_Starter/utils"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type SignUpController struct {
	signUpUsecase domain.SignupUsecase
	otpUsecase    domain.OtpUsecase
}

func NewSignUpController(signUpUsecase domain.SignupUsecase, otpUsecase domain.OtpUsecase) *SignUpController {
	return &SignUpController{
		signUpUsecase: signUpUsecase,
		otpUsecase:    otpUsecase,
	}
}

func (s *SignUpController) SignUp(c *gin.Context) {
	var UserSignUp domain.UserSignUp

	if err := c.ShouldBindJSON(&UserSignUp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := s.signUpUsecase.CreateUser(c, &UserSignUp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate a random number between 0 and 9999 (inclusive).
	randNumber := rand.Intn(10000)

	// Format the code as a 4-digit string with leading zeros.
	code := fmt.Sprintf("%04d", randNumber)

	otp := domain.Otp{
		Email:      user.Email,
		Otp:        code,
		Expiration: time.Now().Add(3 * time.Minute),
	}

	oldOtp, err := s.otpUsecase.GetOtpByEmail(c, UserSignUp.Email)
	if err == nil {
		otp.ID = oldOtp.ID
	} else {
		otp.ID = primitive.NewObjectID()
	}

	// Save OTP to database
	err = s.otpUsecase.SaveOtp(c, &otp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	emailContent := `
	<p>Thank you for signing up. To verify your account and complete the signup process, please use the following verification code:</p>
	<h3>` + code + `</h3>
	<p><strong>This verification code is valid for 5 minutes.</strong> Please enter it on the verification page to proceed.</p>
	<p>If you did not sign up for the BlogApp, please ignore this email.</p>`

	// Create the email subject
	emailSubject := "Verify Your Email"

	// Generate the email body using the template function
	emailBody := EmailUtil.GenerateEmailTemplate("Account Verification", emailContent)

	// Create the email template
	emailTemplate := domain.EmailTemplate{
		Subject: emailSubject,
		Body:    emailBody,
	}

	err = EmailUtil.SendTestEmail(UserSignUp.Email, emailTemplate.Subject, emailTemplate.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "succesfuly sent otp"})

}

func (s *SignUpController) VerifyEmail(c *gin.Context) {
	var VerifyEmailRequest domain.VerifyEmailRequest

	if err := c.ShouldBindJSON(&VerifyEmailRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	VerifyEmailRequest.Email = strings.ToLower(VerifyEmailRequest.Email)

	otp, err := s.otpUsecase.GetOtpByEmail(c, VerifyEmailRequest.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No OTP requested with the given email"})
		return
	}

	if otp.Otp != VerifyEmailRequest.OTP {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP"})
		return
	}else if time.Now().After(otp.Expiration) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OTP expired"})
		return
	}

	user, err := s.signUpUsecase.VerifyEmail(c, &VerifyEmailRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified", "user": user})
}

func (s *SignUpController) ResendOTP(c *gin.Context) {
	var ResendOTPRequest domain.ResendOTPRequest

	if err := c.ShouldBindJSON(&ResendOTPRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ResendOTPRequest.Email = strings.ToLower(ResendOTPRequest.Email)

	err := s.signUpUsecase.ResendOTP(c, &ResendOTPRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	otp, err := s.otpUsecase.GetOtpByEmail(c, ResendOTPRequest.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No OTP requested with the given email"})
		return
	}


	// Generate a random number between 0 and 9999 (inclusive).
	randNumber := rand.Intn(10000)

	// Format the code as a 4-digit string with leading zeros.
	code := fmt.Sprintf("%04d", randNumber)

	otp.Otp = code
	otp.Expiration = time.Now().Add(1 * time.Minute)


	// Save OTP to database
	err = s.otpUsecase.SaveOtp(c, &otp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	emailContent := `
	<p>Thank you for signing up. To verify your account and complete the signup process, please use the following verification code:</p>
	<h3>` + code + `</h3>
	<p><strong>This verification code is valid for 5 minutes.</strong> Please enter it on the verification page to proceed.</p>
	<p>If you did not sign up for the BlogApp, please ignore this email.</p>`

	// Create the email subject
	emailSubject := "Verify Your Email"

	// Generate the email body using the template function
	emailBody := EmailUtil.GenerateEmailTemplate("Account Verification", emailContent)

	// Create the email template
	emailTemplate := domain.EmailTemplate{
		Subject: emailSubject,
		Body:    emailBody,
	}

	err = EmailUtil.SendTestEmail(ResendOTPRequest.Email, emailTemplate.Subject, emailTemplate.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "succesfuly sent otp"})
}