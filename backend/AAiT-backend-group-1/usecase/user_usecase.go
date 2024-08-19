package usecase

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/gin-gonic/gin"
)

type UserUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRespository domain.UserRepository) UserUseCase {
	return UserUseCase{
		userRepo: userRespository,
	}
}

var (
	timeoutStr = os.Getenv("CONTEXT_TIMEOUT")
	timeout    int64
	err        error
)

timeout, err = strconv.ParseInt(timeoutStr, 10, 64)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
func (userUC *UserUseCase) Register(cxt *gin.Context, user *domain.User) domain.Error {
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()
	errValidity := user.Validate()
	if errValidity != nil {
		return &domain.CustomError{Message: errValidity.Error(), Code: http.StatusBadRequest}
	}
	if _, errRepo := userUC.userRepo.Create(context, user); errRepo != nil {
		return errRepo
	}
	return nil
}

func (userUC *UserUseCase) Login(context context.Context, username, password string) (string, domain.Error) {
  
}
func (userUC *UserUseCase) ForgotPassword(context context.Context, email string) domain.Error {

}
func (userUC *UserUseCase) Logout(context context.Context, token string) domain.Error {

}
func (userUC *UserUseCase) PromoteUser(context context.Context, userID string) domain.Error {

}
func (userUC *UserUseCase) DemoteUser(context context.Context, userID string) domain.Error {

}
func (userUC *UserUseCase) UpdateProfile(context context.Context, userID string, user *User) domain.Error {
}
