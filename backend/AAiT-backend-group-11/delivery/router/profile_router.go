package route

import (
	"backend-starter-project/bootstrap"
	"backend-starter-project/delivery/controller"
	"backend-starter-project/infrastructure"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewProfileRouter(env *bootstrap.Env,db *mongo.Database, group *gin.RouterGroup) {
	cloudName := env.CloudName
	apiKey := env.ApiKey
	apiSecret := env.ApiSecret

	imageService,err:= infrastructure.NewImageService(cloudName,apiKey,apiSecret)
	if err != nil {
		panic(err)
	}


	profile_repo := repository.NewProfileRepository(context.TODO(), db)
	profile_service := service.NewProfileService(profile_repo,imageService)
	profile_controller := controller.NewProfileController(profile_service)

	group.GET("/profile/:userId", profile_controller.GetUserProfile)
	group.POST("/profile", profile_controller.CreateUserProfile)
	group.PUT("/profile/:userId", profile_controller.UpdateUserProfile)
	group.DELETE("/profile/:userId", profile_controller.DeleteUserProfile)
	group.POST("/profile/profilePicture", profile_controller.UpdateOrCreateProfilePicture)
	group.PUT("/profile/profilePicture", profile_controller.UpdateOrCreateProfilePicture)
	group.GET("/profile/profilePicture/:userId", profile_controller.GetProfilePicture)
	group.DELETE("/profile/profilePicture/:userId", profile_controller.DeleteProfilePicture)
	group.GET("/profile", profile_controller.GetAllProfiles)
}
