package usecase

import (
	"errors"

	"github.com/RealEskalate/blogpost/config"
	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/infrastructure/emailservices"
	passwordservice "github.com/RealEskalate/blogpost/infrastructure/password_service"
	tokenservice "github.com/RealEskalate/blogpost/infrastructure/token_service"
	"github.com/RealEskalate/blogpost/repository"
)

type EmailVUsecase struct{
	UserUseCase
	repository.EmailVRepo
}

func NewEmailVUsecase(user_usecase *UserUseCase , email_repo *repository.EmailVRepo)*EmailVUsecase {
	return &EmailVUsecase{
		UserUseCase: *user_usecase,
		EmailVRepo:* email_repo,
	}
}

func (uc *EmailVUsecase) SendVerifyEmail(id string , vuser domain.VerifyEmail) error {
	user,err := uc.UserRepo.GetUserDocumentByID(id)
	if err != nil {
		return err
	}

	if user.IsVerified {
		return errors.New("user already verified")
	}

	var tokenizer tokenservice.VerifyToken
	token,err := tokenizer.GenrateToken(id , vuser.Email)
	if err != nil {
		return err
	}
	subject,body := config.ConfigBody(token)

	err = emailservices.SendVerificationEmail(vuser.Email, subject , body)
	if err != nil {
		return err
	}
	
	return nil
}


func (uc *EmailVUsecase) VerifyUser(token string) error {
	id,err := emailservices.IsValidVerificationToken(token)
	if err != nil {
		return err
	}
	return uc.EmailVRepo.VerifyUser(id)
}

func (uc *EmailVUsecase) SendForgretPasswordEmail(id string , vuser domain.VerifyEmail) error {
	var tokenizer tokenservice.VerifyToken
	token,err := tokenizer.GenrateToken(id , vuser.Email)
	if err != nil {
		return err
	}
	subject,body := config.ConfigFogetBody(token , id)

	err = emailservices.SendVerificationEmail(vuser.Email, subject , body)
	if err != nil {
		return err
	}
	
	return nil	
}

func (uc *EmailVUsecase) ValidateForgetPassword(id string , token string) error {
	return passwordservice.IsValidForgetToken(token , id)
}
