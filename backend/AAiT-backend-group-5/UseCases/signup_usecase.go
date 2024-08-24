package usecases

import (
	"context"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/mssola/user_agent"
)

type signupUsecase struct {
	repository   interfaces.UserRepository
	emailService interfaces.EmailService
	jwtService   interfaces.JwtService
	urlService   interfaces.URLService
	otpService   interfaces.OTPService
}

func NewSignupUsecase(repository interfaces.UserRepository,
	emailService interfaces.EmailService,
	jwtService interfaces.JwtService,
	urlService interfaces.URLService,
	otpService interfaces.OTPService,
) interfaces.SignupUsecase {
	return &signupUsecase{
		repository:   repository,
		emailService: emailService,
		jwtService:   jwtService,
		urlService:   urlService,
		otpService:   otpService,
	}
}

func (uc *signupUsecase) CreateUser(ctx context.Context, user *models.User, agent string) *models.ErrorResponse {

	userExist, userExistError := uc.repository.GetUserByEmailOrUsername(ctx, user.Username, user.Email)
	if userExist != nil && userExistError == nil {
		return models.BadRequest("User already exists")
	}

	isValid := uc.emailService.IsValidEmail(user.Email)
	if !isValid {
		return models.BadRequest("Invalid Email")
	}

	token, err := uc.jwtService.CreateURLToken(*user, 60*60)
	if err != nil {
		return models.InternalServerError("Error while creating token")
	}

	subject := "Email Verification"
	body, err := uc.getBody(token, agent)

	e := uc.emailService.SendEmail(user.Email, subject, body)
	if e != nil {
		return models.InternalServerError("Error while sending email")
	}

	return nil

}

func (uc *signupUsecase) handleWebSignup(token string) (string, *models.ErrorResponse) {
	url, err := uc.urlService.GenerateURL(token, "confirmRegistration")

	if err != nil {
		return "", models.InternalServerError("Error while creating url" + err.Error())
	}

	body := "Please click the link below to verify your email address\n" + url + "\nThis link will expire in 1 hour"
	return body, nil
}

func (uc *signupUsecase) handleMobileSignup(token string) (string, *models.ErrorResponse) {
	code, err := uc.otpService.GenerateOTP(token)

	if err != nil {
		return "", models.InternalServerError("Error while creating token")
	}

	body := "Please use the code below to verify you email \n" + code + "\nThis code will expire in 1 hour"
	return body, nil
}

func (uc *signupUsecase) getBody(token string, agent string) (string, *models.ErrorResponse) {
	ua := user_agent.New(agent)
	isMobile := ua.Mobile()

	if isMobile {
		b, nErr := uc.handleMobileSignup(token)
		if nErr != nil {
			return "", nErr
		}

		return b, nil

	}
	b, nErr := uc.handleWebSignup(token)
	if nErr != nil {
		return "", nErr
	}

	return b, nil
}
