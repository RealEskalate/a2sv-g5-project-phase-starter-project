package usecase

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/infrastructure"
	"AAiT-backend-group-6/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
	emailService infrastructure.EmailService
}

func NewSignupUsecase(userRepository domain.UserRepository, timeout time.Duration, emailSrv infrastructure.EmailService) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
		emailService: emailSrv,
	}
}

func (su *signupUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()

	validationErr := infrastructure.ValidateUser(user)
	if validationErr != nil{
		return validationErr
	}

	password := infrastructure.HashPassword(user.Password)
	created_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	
	code, err := utils.GenerateRandomCode(6)
	if err != nil {
		return err
    }
	verification_code := code
    verification_code_expiry := time.Now().Add(24 * time.Hour) // Code expires in 24 hours

	user = &domain.User{
		ID:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Password: password,
		Created_at: created_at,
		Updated_at: updated_at,
		Is_active: false,
		VerificationCode: verification_code,
		VerificationCodeExpiry: verification_code_expiry,
	}

	err = su.userRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return su.emailService.SendEmail(user.Email, user.Name, user.VerificationCode)
	
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

func (su *signupUsecase) GetUserByUsername(c context.Context, username string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByUsername(ctx, username)
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return utils.GenerateAccessToken(user, expiry, secret)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return utils.GenerateRefreshToken(user, expiry, secret)
}