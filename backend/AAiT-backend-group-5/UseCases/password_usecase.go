package usecases

import (
	"context"
	"fmt"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/mssola/user_agent"
)

type setup_password struct {
	urlService      interfaces.URLService
	jwtService      interfaces.JwtService
	emailService    interfaces.EmailService
	passwordService interfaces.PasswordService
	repo            interfaces.UserRepository
	otpService      interfaces.OTPService
}

func NewSetupPassword(
	urlService interfaces.URLService,
	jwtService interfaces.JwtService,
	repo interfaces.UserRepository,
	emailService interfaces.EmailService,
	passwordService interfaces.PasswordService,
	otpService interfaces.OTPService,

) interfaces.PasswordUsecase {

	return &setup_password{
		urlService:      urlService,
		jwtService:      jwtService,
		repo:            repo,
		emailService:    emailService,
		passwordService: passwordService,
		otpService:      otpService,
	}
}

func (sp *setup_password) GenerateResetURL(ctx context.Context, email string, agent string) (string, *models.ErrorResponse) {

	// get user data
	user, uErr := sp.repo.GetUserByEmailOrUsername(ctx, email, email)
	if uErr != nil {
		return "", uErr
	}

	fmt.Println(user)

	// generate token
	token, tErr := sp.jwtService.CreateAccessToken(*user, 60*60)
	if tErr != nil {
		return "", models.InternalServerError("An error occurred while generating the reset URL")
	}

	// generate reset URL
	resetURL, rErr := sp.getBody(token, agent)
	if rErr != nil {
		return "", rErr
	}

	return resetURL, nil
}

func (sp *setup_password) SendResetEmail(ctx context.Context, email string, resetURL string) *models.ErrorResponse {
	subject := "Password Reset"
	body := "you can use the below to reset your password \n" + resetURL + "\nThis link will expire in 1 hour"

	valid := sp.emailService.IsValidEmail(email)
	if !valid {
		return models.BadRequest("Invalid email address")
	}

	err := sp.emailService.SendEmail(email, subject, body)
	if err != nil {
		return models.InternalServerError("An error occurred while sending the reset email")
	}

	return nil
}

func (sp *setup_password) SetNewUserPassword(ctx context.Context, shortURlCode string, password string) *models.ErrorResponse {
	// get token
	urls, tErr := sp.urlService.GetURL(shortURlCode)
	if tErr != nil {
		return tErr
	}

	u, uErr := sp.jwtService.ValidateURLToken(urls.Token)
	if uErr != nil {
		return models.BadRequest("Invalid token")
	}

	if err := sp.passwordService.ValidatePasswordStrength(password); err != nil {
		return err
	}

	hashedPassword, hErr := sp.passwordService.EncryptPassword(password)
	if hErr != nil {
		return models.InternalServerError("An error occurred while setting the password")
	}

	newUser := models.User{
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Email,
		Password: hashedPassword,
		Role:     models.RoleUser,
	}

	// check if user already exists
	user, _ := sp.repo.GetUserByEmailOrUsername(ctx, newUser.Username, newUser.Email)
	if user != nil {
		return models.BadRequest("user already registered")
	}

	// create user
	err := sp.repo.CreateUser(ctx, &newUser)
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

func (sp *setup_password) SetUpdateUserPassword(ctx context.Context, shortURlCode string, password string) *models.ErrorResponse {
	// get token
	urls, tErr := sp.urlService.GetURL(shortURlCode)
	if tErr != nil {
		return tErr
	}

	// get user data
	u, uErr := sp.jwtService.ValidateURLToken(urls.Token)
	if uErr != nil {
		return models.BadRequest("Invalid token")
	}

	// check if password is too short
	if err := sp.passwordService.ValidatePasswordStrength(password); err != nil {
		return err
	}

	// hash password
	hashedPassword, hErr := sp.passwordService.EncryptPassword(password)
	if hErr != nil {
		return models.InternalServerError("An error occurred while setting the password")
	}

	// populate fields for the new user
	updatedUser := models.User{
		Password: hashedPassword,
	}

	// update user
	fmt.Println(u.ID)
	fmt.Println(u.Email)
	fmt.Println(u.Role)

	err := sp.repo.UpdateUser(ctx, &updatedUser, u.ID)
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

func (uc *setup_password) handleWebForgetPassword(token string) (string, *models.ErrorResponse) {
	resetURL, err := uc.urlService.GenerateURL(token, "resetPassword")

	if err != nil {
		return "", models.InternalServerError("Error while creating url" + err.Error())
	}

	return resetURL, nil
}

func (uc *setup_password) handleMobileForgetPassword(token string) (string, *models.ErrorResponse) {
	code, err := uc.otpService.GenerateOTP(token)

	if err != nil {
		return "", models.InternalServerError("Error while creating token")
	}

	return code, nil
}

func (uc *setup_password) getBody(token string, agent string) (string, *models.ErrorResponse) {
	ua := user_agent.New(agent)
	isMobile := ua.Mobile()

	if isMobile {
		url, nErr := uc.handleMobileForgetPassword(token)
		if nErr != nil {
			return "", nErr
		}

		return url, nil

	}
	url, nErr := uc.handleWebForgetPassword(token)
	if nErr != nil {
		return "", nErr
	}

	return url, nil
}
