package controller

import (
	"Blog_Starter/domain"
	utils "Blog_Starter/utils"
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
		c.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	err := UserSignUp.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	user, err := s.signUpUsecase.CreateUser(c, &UserSignUp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{
			Success: false,
			Message: err.Error(),
		})
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

	oldOtp, err := s.otpUsecase.GetOtpByEmail(c, UserSignUp.Email)
	if err == nil {
		otp.ID = oldOtp.ID
	} else {
		otp.ID = primitive.NewObjectID()
	}

	// Save OTP to database
	err = s.otpUsecase.SaveOtp(c, &otp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{
			Success: false,
			Message: err.Error(),
		})
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
	emailBody := utils.GenerateEmailTemplate("Account Verification", emailContent)

	// Create the email template
	emailTemplate := domain.EmailTemplate{
		Subject: emailSubject,
		Body:    emailBody,
	}

	err = utils.SendTestEmail(UserSignUp.Email, emailTemplate.Subject, emailTemplate.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, domain.Response{
		Success: true,
		Message: "Successfully sent OTP",
	})
}

func (s *SignUpController) VerifyEmail(c *gin.Context) {
	var VerifyEmailRequest domain.VerifyEmailRequest

	if err := c.ShouldBindJSON(&VerifyEmailRequest); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}
	VerifyEmailRequest.Email = strings.ToLower(VerifyEmailRequest.Email)

	otp, err := s.otpUsecase.GetOtpByEmail(c, VerifyEmailRequest.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "No OTP requested with the given email",
		})
		return
	}

	if otp.Otp != VerifyEmailRequest.OTP {
		c.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid OTP",
		})
		return
	} else if time.Now().After(otp.Expiration) {
		c.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "OTP expired",
		})
		return
	}

	user, err := s.signUpUsecase.VerifyEmail(c, &VerifyEmailRequest)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.JSON(http.StatusNotFound, domain.Response{
				Success: false,
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, domain.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "Email verified",
		Data:    user,
	})
}

func (s *SignUpController) ResendOTP(c *gin.Context) {
	var ResendOTPRequest domain.ResendOTPRequest

	if err := c.ShouldBindJSON(&ResendOTPRequest); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	ResendOTPRequest.Email = strings.ToLower(ResendOTPRequest.Email)

	err := s.signUpUsecase.ResendOTP(c, &ResendOTPRequest)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.JSON(http.StatusNotFound, domain.Response{
				Success: false,
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, domain.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	otp, err := s.otpUsecase.GetOtpByEmail(c, ResendOTPRequest.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "No OTP requested with the given email",
		})
		return
	}

	// Generate a random number between 0 and 9999 (inclusive).
	randNumber := rand.Intn(10000)

	// Format the code as a 4-digit string with leading zeros.
	code := fmt.Sprintf("%04d", randNumber)

	otp.Otp = code
	otp.Expiration = time.Now().Add(5 * time.Minute)

	// Save OTP to database
	err = s.otpUsecase.SaveOtp(c, &otp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{
			Success: false,
			Message: err.Error(),
		})
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
	emailBody := utils.GenerateEmailTemplate("Account Verification", emailContent)

	// Create the email template
	emailTemplate := domain.EmailTemplate{
		Subject: emailSubject,
		Body:    emailBody,
	}

	err = utils.SendTestEmail(ResendOTPRequest.Email, emailTemplate.Subject, emailTemplate.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "Successfully sent OTP",
	})
}

func (s *SignUpController) FederatedSignup(c *gin.Context) {
	var request domain.FederatedSignupRequest

	// Bind the request body to the FederatedSignupRequest struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// Validate and handle the federated authentication token
	if request.Provider == "google" {
		// Handle federated signup using the usecase
		user, err := s.signUpUsecase.HandleFederatedSignup(c, request.Token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		// Set user object
		c.Set("user", user)

		// Create tokens using the usecase
		tokens, err := s.signUpUsecase.CreateTokens(c, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		signupResponse := domain.LoginResponse{
			UserID:       user.UserID.Hex(),
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		}

		c.JSON(http.StatusOK, domain.Response{
			Success: true,
			Message: "Signup successful",
			Data:    signupResponse,
		})
	} else {
		c.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid federated provider",
		})
	}
}
