package route

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"backend-starter-project/mongo"
	"context"

	"github.com/gin-gonic/gin"
)

func NewProfileRouter(db mongo.Database, group *gin.RouterGroup) {
	profilecollection := db.Collection("profiles")

	profile_repo := repository.NewProfileRepository(context.TODO(), db, profilecollection)
	profile_service := service.NewProfileService(profile_repo)
	profile_controller := controller.NewProfileController(profile_service)

	group.GET("/profile/:userId", profile_controller.GetUserProfile)
	group.POST("/profile", profile_controller.CreateUserProfile)
	group.PUT("/profile/:userId", profile_controller.UpdateUserProfile)
	group.DELETE("/profile/:userId", profile_controller.DeleteUserProfile)

}
