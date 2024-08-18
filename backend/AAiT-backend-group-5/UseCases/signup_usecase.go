package usecases

import (
	"context"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type signupUsecase struct {
	repository   interfaces.UserRepository
	emailService interfaces.EmailService
}

func NewSignupUsecase(repository interfaces.UserRepository, emailService interfaces.EmailService) interfaces.SignupUsecase {
	return &signupUsecase{
		repository:   repository,
		emailService: emailService,
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

	// update the role
	user.Role = models.RoleUser

	// createUser
	if err := uc.repository.CreateUser(ctx, *user); err != nil {
		return err
	}

	// send the email for varification of the email address

	return nil

}

func (uc *signupUsecase) GetUserByID(ctx context.Context, id string) (*models.User, *models.ErrorResponse) {
	return uc.repository.GetUserByID(ctx, id)
}
