package usecase

import (
	"blog_g2/domain"
	"context"

	"time"
)

type UserUsecase struct {
	UserRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(Userrepo domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &UserUsecase{
		UserRepo:       Userrepo,
		contextTimeout: timeout,
	}

}

func (uuse *UserUsecase) UpdateUserDetails(c context.Context, user *domain.User) *domain.AppError {
	_, cancel := context.WithTimeout(c, uuse.contextTimeout)
	defer cancel()
	return uuse.UserRepo.UpdateUserDetails(user)
}
func (uuse *UserUsecase) RegisterUser(c context.Context, user *domain.User) *domain.AppError {
	return uuse.UserRepo.RegisterUser(user)

}

func (uuse *UserUsecase) LoginUser(c context.Context, user domain.User) (string, *domain.AppError) {
	return uuse.UserRepo.LoginUser(user)
}
func (uuse *UserUsecase) VerifyUserEmail(c context.Context, token string) *domain.AppError {
	_, cancel := context.WithTimeout(c, uuse.contextTimeout)
	defer cancel()
	return uuse.UserRepo.VerifyUserEmail(token)
}

func (uuse *UserUsecase) ForgotPassword(c context.Context, email string) *domain.AppError {
	_, cancel := context.WithTimeout(c, uuse.contextTimeout)
	defer cancel()
	return uuse.UserRepo.ForgotPassword(email)
}

func (uuse *UserUsecase) ResetPassword(c context.Context, token string, newPassword string) *domain.AppError {
	_, cancel := context.WithTimeout(c, uuse.contextTimeout)
	defer cancel()
	return uuse.UserRepo.ResetPassword(token, newPassword)
}

func (uuse *UserUsecase) LogoutUser(c context.Context, uid string) *domain.AppError {
	return uuse.UserRepo.LogoutUser(uid)
}

func (uuse *UserUsecase) PromoteDemoteUser(c context.Context, userid string, isAdmin bool) *domain.AppError {
	_, cancel := context.WithTimeout(c, uuse.contextTimeout)
	defer cancel()
	return uuse.UserRepo.PromoteDemoteUser(userid, isAdmin)
}
