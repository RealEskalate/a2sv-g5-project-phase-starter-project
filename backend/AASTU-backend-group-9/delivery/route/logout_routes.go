package route

import (
    "blog/config"
    "blog/database"
    "blog/delivery/controller"
    "blog/repository"
    "blog/usecase"
    "time"

    "github.com/gin-gonic/gin"
)

func NewLogoutRouter(env *config.Env, timeout time.Duration, db database.Database, r *gin.RouterGroup) {
    tokenRepo := repository.NewMongoTokenRepository(db.Collection("tokens"))
    logoutUsecase := usecase.NewLogoutUsecase(tokenRepo, timeout)
    logoutController := &controller.LogoutController{
        LogoutUsecase: logoutUsecase,
    }

    r.POST("/logout", logoutController.Logout)
}
