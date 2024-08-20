package route

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewProfileRouter(db *mongo.Database, group *gin.RouterGroup) {
	profile_repo := repository.NewProfileRepository(context.TODO(), db)
	profile_service := service.NewProfileService(profile_repo)
	profile_controller := controller.NewProfileController(profile_service)

	group.POST("profile", profile_controller.CreateUserProfile)
	group.GET("profile/:userId", profile_controller.GetUserProfile)
	group.PUT("profile", profile_controller.UpdateUserProfile)
	group.DELETE("profile/:userId", profile_controller.DeleteUserProfile)

}
