package routers

import (
	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	repository "github.com/aait.backend.g5.main/backend/Repository"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserProfileRouter(database mongo.Database, group *gin.RouterGroup) {

	user_repo := repository.NewUserRepository(&database)
	password_service := infrastructure.NewPasswordService()

	// instantiate userProfile_controller
	UserProfileController := &controllers.UserProfileController{
		UserProfileUC: usecases.NewUserProfileUpdateUsecase(user_repo, password_service),
	}

	group.PUT("/profile", UserProfileController.ProfileUpdate)
	group.GET("/profile", UserProfileController.ProfileGet)
	group.DELETE("/profile", UserProfileController.ProfileDelete)
}
