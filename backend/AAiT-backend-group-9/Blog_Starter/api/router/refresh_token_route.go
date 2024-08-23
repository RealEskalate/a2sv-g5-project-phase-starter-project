package router

import (
	"time"

	"Blog_Starter/api/controller"
	"Blog_Starter/config"
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"Blog_Starter/usecase"
	"Blog_Starter/utils/infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewRefreshTokenRouter sets up the refreshtoken route.

func NewRefreshTokenRouter(env *config.Env, timeout time.Duration, db *mongo.Client, group *gin.RouterGroup) {
	// Initialize the database
	database := db.Database(env.DBName)

	// Initialize repositories
	ur := repository.NewUserRepository(database, domain.CollectionUser)
	tm := &infrastructure.NewTokenManager{} 
	refreshTokenUsecase := usecase.NewRefreshTokenUsecase(ur, tm, timeout, env)

	refreshTokenController := controller.NewRefreshTokenController(refreshTokenUsecase)

	group.POST("/refreshtoken", refreshTokenController.RefreshToken)
}