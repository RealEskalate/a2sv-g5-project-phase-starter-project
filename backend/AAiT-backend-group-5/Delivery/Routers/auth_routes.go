package routers

import (
	config "github.com/aait.backend.g5.main/backend/Config"
	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	repository "github.com/aait.backend.g5.main/backend/Repository"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthenticationRouter(env *config.Env, database mongo.Database, group *gin.RouterGroup) {
	emailConfig := config.NewEmailServer(*env)

	url_repository := repository.NewURLRepository(&database)
	user_repository := repository.NewUserRepository(&database)
	session_repository := repository.NewSessionRepository(&database)

	jwt_service := infrastructure.NewJwtService(env)
	password_service := infrastructure.NewPasswordService()
	email_service := infrastructure.NewEmailService(emailConfig, *env)
	url_service := infrastructure.NewURLService(env, url_repository)

	// instantiate login controller
	LoginController := &controllers.LoginController{
		LoginUsecase: usecases.NewLoginUsecase(jwt_service, password_service, user_repository, session_repository, *env),
		Env:          env,
	}

	// instantiate signup controller
	SignupController := &controllers.SignupController{
		SignupUsecase:   usecases.NewSignupUsecase(user_repository, email_service, jwt_service, url_service),
		PasswordUsecase: usecases.NewSetupPassword(url_service, jwt_service, user_repository, email_service, password_service),
	}

	group.POST("/signup", SignupController.Signup)
	group.POST("/confirmRegistration/:id", SignupController.ConfirmRegistration)

	group.POST("/login", LoginController.Login)
}
