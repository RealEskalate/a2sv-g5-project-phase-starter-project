package usecases

import (
	"context"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type setup_password struct {
	urlService      interfaces.URLService
	jwtService      interfaces.JwtService
	emailService    interfaces.EmailService
	passwordService interfaces.PasswordService
	repo            interfaces.UserRepository
}

func NewSetupPassword(
	urlService interfaces.URLService,
	jwtService interfaces.JwtService,
	repo interfaces.UserRepository,
	emailService interfaces.EmailService,
	passwordService interfaces.PasswordService,
	
) interfaces.PasswordUsecase {

	return &setup_password{
		urlService:      urlService,
		jwtService:      jwtService,
		repo:            repo,
		emailService:    emailService,
		passwordService: passwordService,
	}
}

func (sp *setup_password) GenerateResetURL(ctx context.Context, email string) (string, *models.ErrorResponse) {

	// get user data
	user, uErr := sp.repo.GetUserByEmailOrUsername(ctx, email, email)
	if uErr != nil {
		return "", uErr
	}

	// generate token
	token, tErr := sp.jwtService.CreateAccessToken(*user, 60*60)
	if tErr != nil {
		return "", models.InternalServerError("An error occurred while generating the reset URL")
	}

	// generate reset URL
	resetURL, rErr := sp.urlService.GenerateURL(token)
	if rErr != nil {
		return "", rErr
	}

	return resetURL, nil
}

func (sp *setup_password) SendResetEmail(ctx context.Context, email string, resetURL string) *models.ErrorResponse {
	subject := "Password Reset"
	body := "Click the link below to reset your password\n" + resetURL + "\n\nThis link will expire in 1 hour"

	// validate email
	valid := sp.emailService.IsValidEmail(email)
	if !valid {
		return models.BadRequest("Invalid email address")
	}

	// send email
	err := sp.emailService.SendEmail(email, subject, body)
	if err != nil {
		return models.InternalServerError("An error occurred while sending the reset email")
	}

	return nil
}

func (sp *setup_password) SetPassword(ctx context.Context, shortURlCode string, password string) *models.ErrorResponse {
	// get token
	urls, tErr := sp.urlService.GetURL(shortURlCode)

	if tErr != nil {
		return tErr
	}

	// get user data
	u, uErr := sp.jwtService.ValidateToken(urls.Token)
	if uErr != nil {
		return models.BadRequest("Invalid token")
	}

	user, gErr := sp.repo.GetUserByID(ctx, u.ID)
	if gErr != nil {
		return gErr
	}

	// hash password
	hashedPassword, hErr := sp.passwordService.EncryptPassword(password)
	if hErr != nil {
		return models.InternalServerError("An error occurred while setting the password")
	}

	user.Password = hashedPassword

	// update user
	err := sp.repo.UpdateUser(ctx, user, u.ID)
	if err != nil {
		return err
	}

	// remove token
	err = sp.urlService.RemoveURL(shortURlCode)
	if err != nil {
		return models.InternalServerError("An error occurred while setting the password")
	}

	return nil
}
