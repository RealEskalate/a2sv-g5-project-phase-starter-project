package router

import (
	"time"

	"Blog_Starter/api/controller"
	"Blog_Starter/config"
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"Blog_Starter/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewSignupRouter sets up the signup routes.
func NewUserRouter(env *config.Env, timeout time.Duration, db *mongo.Client, group *gin.RouterGroup) {
    database := db.Database(env.DBName) // Replace with your actual database name
    ur := repository.NewUserRepository(database, domain.CollectionUser)
    uc := controller.NewUserController(
        usecase.NewUserUsecase(ur, timeout),
    )

    group.GET("/users", uc.GetAllUsers)
    group.POST("/promote", uc.PromoteUser)
    group.POST("/demote", uc.DemoteUser)
    group.DELETE("/delete", uc.DeleteUser)
    group.PATCH("/updateuser", uc.UpdateUser)
    group.PATCH("/updateprofilepicture", uc.UpdateProfilePicture)
	group.DELETE("/deleteprofilepicture", uc.DeleteProfilePicture)
	
}