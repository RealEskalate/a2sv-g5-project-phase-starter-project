package routers

import (
	"blogapp/Delivery/controllers"
	"blogapp/Infrastructure/auth_middleware"
	"blogapp/Repositories"
	usecases "blogapp/UseCases"
)

// refreshRouter
func RefreshRouter() {
	refreshRouter := Router.Group("/refresh")
	{
		// generate new auth repo
		refreshrepo := Repositories.NewRefreshRepository(BlogCollections.RefreshTokens)
		refreshusecase := usecases.NewRefreshUseCase(refreshrepo)
		refreshcontroller := controllers.NewRefreshController(refreshusecase)

		refreshRouter.POST("/refresh", auth_middleware.AuthMiddleware(), refreshcontroller.Refresh)
	}
}
