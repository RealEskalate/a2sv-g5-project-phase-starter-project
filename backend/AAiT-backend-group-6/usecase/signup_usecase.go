package usecase

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/infrastructure"
	"AAiT-backend-group-6/utils"
	"context"
	"errors"
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

	password := infrastructure.HashPassword(user.Password)
	created_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	
	code, err := utils.GenerateRandomCode(6)
	if err != nil {
		return err
    }
	verification_code := infrastructure.HashPassword(code)
    verification_code_expiry := time.Now().Add(24 * time.Hour) // Code expires in 24 hours

	user = &domain.User{
		ID:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Password: password,
		User_type: "USER",
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

	msg := su.emailService.EmailVerificationMsg(user.Email, user.Name, code)

	return su.emailService.SendEmail(user.Email, msg)
	
}


func (su *signupUsecase) VerifyEmail(c context.Context, email string, code string) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()

	user, err := su.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	if user.Is_active{
		return errors.New("this accout is already verified")
	}

	if err := infrastructure.VerifyPassword(code, user.VerificationCode); err != nil{
		return err
	}

	if time.Now().After(user.VerificationCodeExpiry) {
		return errors.New("verification code has expired")
	}

	return nil
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return utils.GenerateAccessToken(user, expiry, secret)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return utils.GenerateRefreshToken(user, expiry, secret)
}