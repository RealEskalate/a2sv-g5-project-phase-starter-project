package usecases

import (
	"context"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type signupUsecase struct {
	repository   interfaces.UserRepository
	emailService interfaces.EmailService
	jwtService   interfaces.JwtService
	urlService   interfaces.URLService
}

func NewSignupUsecase(repository interfaces.UserRepository,
	emailService interfaces.EmailService,
	jwtService interfaces.JwtService,
	urlService interfaces.URLService) interfaces.SignupUsecase {
	return &signupUsecase{
		repository:   repository,
		emailService: emailService,
		jwtService:   jwtService,
		urlService:   urlService,
	}
}

func (uc *signupUsecase) CreateUser(ctx context.Context, user *models.User) *models.ErrorResponse {

	// check user doesn't exist
	userExist, userExistError := uc.repository.GetUserByEmailOrUsername(ctx, user.Username, user.Email)

	if userExist != nil && userExistError == nil {
		return models.BadRequest("User already exists")
	}

	// check if email is valid
	isValid := uc.emailService.IsValidEmail(user.Email)

	if !isValid {
		return models.BadRequest("Invalid Email")
	}

	// send the email for varification of the email address
	token, err := uc.jwtService.CreateURLToken(*user, 60*60)

	if err != nil {
		return models.InternalServerError("Error while creating token")
	}

	url, uErr := uc.urlService.GenerateURL(token, "confirmRegistration")

	if uErr != nil {
		return models.InternalServerError("Error while creating url" + uErr.Error())
	}

	// send the email
	subject := "Email Verification"
	body := "Please click the link below to verify your email address\n" + url + "This link will expire in 1 hour"
	e := uc.emailService.SendEmail(user.Email, subject, body)

	if e != nil {
		return models.InternalServerError("Error while sending email")
	}

	return nil

}
