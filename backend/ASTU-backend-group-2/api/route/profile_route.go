package route

import (
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/controller"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/middleware"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/repository"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/usecase"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewProfileRouter is a function that defines all the routes for the profile
func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cloudinary *cloudinary.Cloudinary) {
	ur := repository.NewUserRepository(*db, entities.CollectionUser)
	pc := controller.ProfileController{
		UserUsecase: usecase.NewUserUsecase(ur, timeout),
		Env:         env,
	}

	group.GET("/profiles", middleware.AdminMiddleware(), pc.GetProfiles())
	group.GET("/profiles/:id", pc.GetProfile())
	group.PUT("/profiles/:id", pc.UpdateProfile())
	group.PATCH("/profiles/:id", pc.UpdateProfile())
	group.DELETE("/profiles/:id", pc.DeleteProfile())
	group.PUT("/profile-pic/", middleware.FileUploadMiddleware(), pc.UploadProfilePicture(cloudinary))
	group.POST("/profiles/", pc.ChangePassword())

	// promote/demote user to admin
	group.POST("/profiles/:id/promote", middleware.AdminMiddleware(), pc.PromoteUser())
	group.POST("/profiles/:id/demote", middleware.AdminMiddleware(), pc.DemoteUser())
}
