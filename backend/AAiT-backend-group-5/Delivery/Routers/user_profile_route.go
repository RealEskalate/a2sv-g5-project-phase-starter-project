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

func NewUserProfileRouter(database interfaces.Database, env config.Env, group *gin.RouterGroup) {

	user_repo := repository.NewUserRepository(database)
	password_service := infrastructure.NewPasswordService()
	session_repository := repository.NewSessionRepository(database)
	cld, _ := config.NewCloudinaryConfig(env)

	cloudinary_service := infrastructure.NewCloudinaryService(cld, &env)

	UserProfileController := &controllers.UserProfileController{
		UserProfileUC: usecases.NewUserProfileUpdateUsecase(user_repo, password_service, cloudinary_service, session_repository),
	}

	group.PUT("/profile", UserProfileController.ProfileUpdate)
	group.GET("/profile", UserProfileController.ProfileGet)
	group.DELETE("/profile", UserProfileController.ProfileDelete)
}
