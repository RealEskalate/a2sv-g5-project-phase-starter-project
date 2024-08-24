package routers

import (
	config "github.com/aait.backend.g5.main/backend/Config"
	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	repository "github.com/aait.backend.g5.main/backend/Repository"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"

	"github.com/gin-gonic/gin"
)

func NewForgotPasswordRouter(env *config.Env, database interfaces.Database, group *gin.RouterGroup) {
	emailConfig := config.NewEmailServer(*env)

	url_repo := repository.NewURLRepository(database)
	user_repo := repository.NewUserRepository(database)

	jwt_service := infrastructure.NewJwtService(env)
	password_service := infrastructure.NewPasswordService()
	email_service := infrastructure.NewEmailService(emailConfig, *env)
	url_service := infrastructure.NewURLService(env, url_repo)

	otp_service := infrastructure.NewOTPService(url_repo)

	// instantiate ForgotPassword controller
	ForgotPasswordController := &controllers.ForgotPasswordController{
		PasswordUsecase: usecases.NewSetupPassword(url_service, jwt_service, user_repo, email_service, password_service, otp_service),
	}

	group.POST("/forgotPassword", ForgotPasswordController.ForgotPasswordRequest)
	group.POST("/resetPassword/:id", ForgotPasswordController.SetNewPassword)
}
