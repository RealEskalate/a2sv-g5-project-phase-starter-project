package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"backend-starter-project/utils"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/exp/rand"
)

type authService struct {
	userService          interfaces.UserService
	passwordResetService interfaces.PasswordResetService
	passwordService      interfaces.PasswordService
	tokenService         interfaces.TokenService
	otpService           interfaces.OTPService
}

func NewAuthService(userService interfaces.UserService, passwordResetService interfaces.PasswordResetService,
	passService interfaces.PasswordService, tokenService interfaces.TokenService, otpService interfaces.OTPService) interfaces.AuthenticationService {
	return &authService{
		userService:          userService,
		tokenService:         tokenService,
		passwordResetService: passwordResetService,
		otpService:           otpService,
		passwordService:      passService,
	}
}

func (service *authService) RegisterUser(user *entities.User) (*entities.User, error) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if the user already exists
	_, err := service.userService.FindUserByEmail(user.Email)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	// Hash the password
	hashedPassword, err := utils.NewPasswordService().HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	// Set the hashed password
	user.Password = hashedPassword

	// Create the user
	createdUser, err := service.userService.CreateUser(user)
	if err != nil {
		return nil, err
	}

	// Helper function to delete the user
	deleteUser := func() error {
		return service.userService.DeleteUser(user.ID.Hex())
	}
	
	rand.Seed(uint64(time.Now().UnixNano()))

	// Generate a random number between 10000 and 99999 (inclusive).
	randNum := rand.Intn(99999-10000+1) + 10000
	
	// Format the code as a 5-digit string with leading zeros.
	code := fmt.Sprintf("%05d", randNum)
	
	// Create a new OTP object with the email and code.
	otp := entities.OTP{
		Email:      user.Email,
		Code:       code,
		Expiration: time.Now().Add(5 * time.Minute),
	}
	
	// Retrieve existing OTP for the email
	oldOtp, err := service.otpService.GetOtpByEmail(user.Email)
	if err == nil {
		// Invalidate the old OTP if it exists
		if time.Now().Before(oldOtp.Expiration) {
			err = service.otpService.InvalidateOtp(&oldOtp) // Call to invalidate old OTP
			if err != nil {
				deleteErr := deleteUser() // Attempt to delete the user
				if deleteErr != nil {
					return nil, fmt.Errorf("failed to invalidate OTP and delete user: %v", deleteErr)
				}
				return nil, err
			}
		}
		otp.ID = oldOtp.ID
		} else {
			otp.ID = primitive.NewObjectID()
		}
		
		// Save the new OTP
		err = service.otpService.SaveOtp(&otp)
		if err != nil {
			deleteErr := deleteUser() // Attempt to delete the user
			if deleteErr != nil {
				return nil, fmt.Errorf("failed to save OTP and delete user: %v", deleteErr)
			}
			return nil, err
	}

	emailContent := `
		<p>Thank you for signing up with Blog. To verify your account and complete the signup process, please use the following verification code:</p>
		<h3>` + code + `</h3>
		<p><strong>This verification code is valid for 5 minutes.</strong> Please enter it on the verification page to proceed.</p>
		<p>If you did not sign up for a Blog account, please ignore this email.</p>`

	// Create the email subject
	emailSubject := "Verify Your Email"

	smtpConfig := entities.SMTPConfig{
		Server:   "smtp.gmail.com:587",
		Username: "haloitisme0912@gmail.com",
		Password: "btnb soyo xqpm ooxw",
	}

	// Generate the email body using the template function
	emailBody := utils.NewEmailService(smtpConfig.Server, smtpConfig.Password, smtpConfig.Username).GenerateEmailTemplate("Blog Account Verification", emailContent)

	// Create the email template
	emailTemplate := entities.EmailTemplate{
		Subject: emailSubject,
		Body:    emailBody,
	}

	// Send the email
	err = utils.NewEmailService(smtpConfig.Server, smtpConfig.Password, smtpConfig.Username).SendEmail(user.Email, emailTemplate.Subject, emailTemplate.Body)
	if err != nil {
		deleteErr := deleteUser() // Attempt to delete the user
		if deleteErr != nil {
			return nil, fmt.Errorf("failed to send email and delete user: %v", deleteErr)
		}

		return nil, err
	}

	return createdUser, nil
}

func (service *authService) Login(emailOrUsername, password string) (*entities.RefreshToken, string, error) {
	user, _ := service.userService.FindUserByEmail(emailOrUsername)
	if user == nil {
		return nil, "", errors.New("User not found")
	}
	err := service.passwordService.ComparePassword(user.Password, password)
	if err != nil {
		return nil, "", errors.New("Invalid password")
	}
	token, err := service.tokenService.GenerateAccessToken(user)
	if err != nil {
		return nil, "", err
	}
	refresh_tok, err := service.tokenService.GenerateRefreshToken(user)
	if err != nil {
		return nil, "", err
	}
	return refresh_tok, token, nil
}

func (service *authService) Logout(userId string) error {

	//delete the token from database
	err := service.tokenService.DeleteRefreshTokenByUserId(userId)
	if err != nil {
		return err
	}
	return nil

}

func (service *authService) RefreshAccessToken(token string) (string, error) {
	refreshed, err := service.tokenService.RefreshAccessToken(token)
	if err != nil {
		return "", err
	}
	return refreshed, nil
}

func (service *authService) VerifyEmail(email string, code string) error {
	// Retrieve the OTP by email
	otp, err := service.otpService.GetOtpByEmail(email)
	if err != nil {
		return err
	}

	// Check if the OTP is expired
	if time.Now().After(otp.Expiration) {
		return errors.New("verification code has expired")
	}

	// Check if the code matches
	if otp.Code != code {
		return errors.New("verification code is incorrect")
	}

	// Mark the user as verified
	err = service.userService.MarkUserAsVerified(email)
	if err != nil {
		return err
	}

	// Invalidate the OTP
	err = service.otpService.InvalidateOtp(&otp)
	if err != nil {
		return err
	}

	return nil
}

func (service *authService) ResendOtp(request entities.ResendOTPRequest) error {
	request.Email = strings.ToLower(request.Email)

	user, err := service.userService.FindUserByEmail(request.Email)
	if err != nil {
		return err
	}

	if user.IsVerified {
		return fmt.Errorf("failed to resend otp. User account already activated")
	}

	otp, err := service.otpService.GetOtpByEmail(request.Email)
	if err != nil {
		return fmt.Errorf("no OTP requested with the given email")
	}

	rand.Seed(uint64(time.Now().UnixNano()))

	// Generate a random number between 10000 and 99999 (inclusive).
	randNum := rand.Intn(99999-10000+1) + 10000
	
	// Format the code as a 5-digit string with leading zeros.
	code := fmt.Sprintf("%05d", randNum)

	// Modify OTP to database
	otp.Code = code
	otp.Expiration = time.Now().Add(5 * time.Minute)

	err = service.otpService.SaveOtp(&otp)
	if err != nil {
		return fmt.Errorf("failed to save OTP: %v", err.Error())
	}

	emailContent := `
		<p>Thank you for signing up with Blog. To verify your account and complete the signup process, please use the following verification code:</p>
		<h3>` + code + `</h3>
		<p><strong>This verification code is valid for 5 minutes.</strong> Please enter it on the verification page to proceed.</p>
		<p>If you did not sign up for a Blog account, please ignore this email.</p>`

	// Create the email subject
	emailSubject := "Verify Your Email"

	smtpConfig := entities.SMTPConfig{
		Server:   "smtp.gmail.com:587",
		Username: "haloitisme0912@gmail.com",
		Password: "btnb soyo xqpm ooxw",
	}
 
	// Generate the email body using the template function
	emailBody := utils.NewEmailService(smtpConfig.Server, smtpConfig.Password, smtpConfig.Username).GenerateEmailTemplate("Blog Account Verification", emailContent)
 
	// Create the email template
	emailTemplate := entities.EmailTemplate{
		Subject: emailSubject,
		Body:    emailBody,
	}
 
	// Send the email
	err = utils.NewEmailService(smtpConfig.Server, smtpConfig.Password, smtpConfig.Username).SendEmail(user.Email, emailTemplate.Subject, emailTemplate.Body)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err.Error())
	}

	return nil
}