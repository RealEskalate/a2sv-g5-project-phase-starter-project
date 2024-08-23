package routers

import (
	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	repository "github.com/aait.backend.g5.main/backend/Repository"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserProfileRouter(database mongo.Database, group *gin.RouterGroup) {

	user_repo := repository.NewUserRepository(&database)
	password_service := infrastructure.NewPasswordService()
	cld, _ := cloudinary.NewFromParams("dyninpclo", "889999668498927", "GIjLypw6vC7pkXuSuVBh94D8rXA")

	// instantiate userProfile_controller
	UserProfileController := &controllers.UserProfileController{
		UserProfileUC: usecases.NewUserProfileUpdateUsecase(user_repo, password_service, cld),
	}

	group.PUT("/profile", UserProfileController.ProfileUpdate)
}
