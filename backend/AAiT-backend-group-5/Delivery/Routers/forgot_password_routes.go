package routers

import (
	"time"

	config "github.com/aait.backend.g5.main/backend/Config"
	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	repository "github.com/aait.backend.g5.main/backend/Repository"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewForgotPasswordRouter(env *config.Env, timeout time.Duration, database mongo.Database, group *gin.RouterGroup) {
	emailConfig := config.NewEmailServer(*env)

	url_repository := repository.NewURLRepository(&database)
	user_repository := repository.NewUserRepository(&database)

	jwt_service := infrastructure.NewJwtService(env)
	password_service := infrastructure.NewPasswordService()
	email_service := infrastructure.NewEmailService(emailConfig, *env)
	url_service := infrastructure.NewURLService(env, url_repository)

	ForgotPasswordController := &controllers.ForgotPasswordController{
		PasswordUsecase: usecases.NewSetupPassword(url_service, jwt_service, user_repository, email_service, password_service),
	}

	group.POST("/forgotPassword", ForgotPasswordController.ForgotPasswordRequest)
	group.POST("/resetPassword", ForgotPasswordController.SetNewPassword)
}
