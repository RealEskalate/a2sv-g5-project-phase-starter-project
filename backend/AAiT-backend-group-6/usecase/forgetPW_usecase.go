package usecase

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/infrastructure"
	"AAiT-backend-group-6/utils"
	"context"
	"errors"
	"time"
)

type forgetPWUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
	emailService   infrastructure.EmailService
}

func NewForgetPWUsecase(userRepository domain.UserRepository, timeout time.Duration, emailService infrastructure.EmailService) domain.ForgetPWUsecase {
	return &forgetPWUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
		emailService:   emailService,
	}
}


func (fpu *forgetPWUsecase) ForgetPW(c context.Context, email string, server_address string) error {
	ctx, cancel := context.WithTimeout(c, fpu.contextTimeout)
	defer cancel()
	user, err := fpu.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	pw_recovery_code := utils.GenerateRecoveryToken()
	user.PWRecoveryToken = infrastructure.HashPassword(pw_recovery_code)
	user.PWRecoveryTokenExpiry = time.Now().Add(1 * time.Hour)

	err = fpu.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	pw_recovery_link := utils.GenerateRecoveryLink(server_address, user.Username, pw_recovery_code)
	msg := fpu.emailService.PWRecoveryMsg(user.Email, user.Name, pw_recovery_link)
	return fpu.emailService.SendEmail(user.Email, msg)
}

func (fpu *forgetPWUsecase) ResetPW(c context.Context, resquest domain.ResetPWRequest) error {
	ctx, cancel := context.WithTimeout(c, fpu.contextTimeout)
	defer cancel()

	user, err := fpu.userRepository.GetUserByEmail(ctx, resquest.Email)
	if err != nil {
		return err
	}

	password := infrastructure.HashPassword(user.Password)
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	recovery_link_expiry := time.Now()

	user = &domain.User{
		ID:                   user.ID,
		Password:             password,
		Updated_at:           updated_at,
		PWRecoveryTokenExpiry: recovery_link_expiry,
	}

	err = fpu.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}


func (fpu *forgetPWUsecase) VerifyForgetPWRequest(c context.Context, username string, recovery_token string) error {
	ctx, cancel := context.WithTimeout(c, fpu.contextTimeout)
	defer cancel()
	user, err := fpu.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	if err := infrastructure.VerifyPassword(recovery_token, user.PWRecoveryToken); err != nil{
		return errors.New("Invalid Recovery link")
	}

	if time.Now().After(user.PWRecoveryTokenExpiry) {
		return errors.New("Recovery token has expired")
	}

	return nil
}

func (fpu *forgetPWUsecase) GenerateRecoveryLink(server_address string, username string, recoveryToken string) (recoveryLink string) {
	return utils.GenerateRecoveryLink(server_address, username, recoveryToken)
}