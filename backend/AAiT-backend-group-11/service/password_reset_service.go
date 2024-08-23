package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"backend-starter-project/utils"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type passwordResetService struct {
	passwordTokenRepository	interfaces.PasswordTokenRepository	
	userRepository interfaces.UserRepository
	passwordResetTokenSecret string
}

func NewPasswordResetService(passwordResetTokenSecret string, userRepository interfaces.UserRepository, passwordTokenRepository	interfaces.PasswordTokenRepository) interfaces.PasswordResetService{
	return &passwordResetService{
		userRepository: userRepository,
		passwordResetTokenSecret: passwordResetTokenSecret,
		passwordTokenRepository: passwordTokenRepository,
	}
}

func (pass_service *passwordResetService) GeneratePasswordResetToken(user *entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": user.ID.Hex(),
			"email":  user.Email,
			"exp":    time.Now().Add(time.Hour * 5).Unix(), //TODO: reduce the expiry time
		})
	
	passwordToken := []byte(pass_service.passwordResetTokenSecret)

	// Sign the token with the secret
	passwordResetToken, err := token.SignedString(passwordToken)
	if err != nil {
		return "", err
	}

	// Store the token in the database
	passwordResetTokenEntity := entities.PasswordResetToken{
		Token:     passwordResetToken,
		ExpiresAt: time.Now().Add(2 * time.Minute),
		UserID:    user.ID,
	}

	_, err = pass_service.passwordTokenRepository.CreatePasswordResetToken(&passwordResetTokenEntity)

	if err != nil {
		return "", err
	}

	return passwordResetToken, nil
}

func (pass_service *passwordResetService) RequestPasswordReset(email string) error {

	// Find the user with the given email
    user, err := pass_service.userRepository.FindUserByEmail(email)
	if err != nil{
		return err
	}

    // Generate a secure token
	token, err := pass_service.GeneratePasswordResetToken(user)
	if err != nil {
		return err
	}

	smtpConfig := entities.SMTPConfig{
		Server:   "smtp.gmail.com:587",
		Username: "haloitisme0912@gmail.com",
		Password: "btnb soyo xqpm ooxw",
	}

    // send an email to the user with the password reset link
	emailSubject := "Blog Password Reset Request"

	resetLink := fmt.Sprintf("http://localhost:8080/reset-password?token=%s", token)

	emailContent := fmt.Sprintf(`
		<p>We received a request to reset your password for your Blog account. You can reset your password by clicking the link below:</p>
		<p><a href="%s">Reset Password</a></p>
		<p>If you did not request a password reset, please ignore this email or contact support if you have questions.</p>
		<p><strong>This password reset link is valid for 5 minutes.</strong></p>`, resetLink)
	
	emailTemplate := entities.EmailTemplate{
		Subject: emailSubject,
		Body:    emailContent,
	}

	err = utils.NewEmailService(smtpConfig.Server, smtpConfig.Password, smtpConfig.Username).SendEmail(email, emailTemplate.Subject, emailTemplate.Body)

	if err != nil {
		return err
	}

    return nil
}


func (pass_service *passwordResetService) ResetPassword(token,newPass string)error{
	// Find the token in the database
	passwordResetToken, err := pass_service.passwordTokenRepository.FindPasswordReset(token)
	if err != nil {
		return err
	}

	// Check if the token has expired
	if time.Now().After(passwordResetToken.ExpiresAt) {
		pass_service.passwordTokenRepository.DeletePasswordResetTokenByUserId(passwordResetToken.UserID.Hex())
		return fmt.Errorf("password reset token has expired")
	}

	// Find the user with the given token
	user, err := pass_service.userRepository.FindUserById(passwordResetToken.UserID.Hex())
	if err != nil {
		return err
	}

	// Hash the new password
	hashedPassword, err := utils.NewPasswordService().HashPassword(newPass)
	if err != nil {
		return err
	}

	// Update the user's password
	user.Password = hashedPassword
	_, err = pass_service.userRepository.UpdateUser(user)
	if err != nil {
		return err
	}

	// Delete the password reset token
	err = pass_service.passwordTokenRepository.DeletePasswordResetTokenByUserId(user.ID.Hex())
	if err != nil {
		return err
	}

	return nil
}