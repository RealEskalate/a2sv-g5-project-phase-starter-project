package routers

import (
	"blogapp/Delivery/controllers"
	"blogapp/Infrastructure/auth_middleware"
	"blogapp/Repositories"
	usecases "blogapp/UseCases"
)

func ProfileRouter() {
	profileRouter := Router.Group("/me")
	profileRouter.Use(auth_middleware.AuthMiddleware())
	{

		// generate new auth repo
		profile_repo := Repositories.NewProfileRepository(BlogCollections.Users, BlogCollections.RefreshTokens)

		profile_usecase := usecases.NewProfileUseCase(profile_repo)
		profile_controller := controllers.NewProfileController(profile_usecase)

		// get all users
		profileRouter.GET("/", profile_controller.GetProfile)
		profileRouter.PUT("/", profile_controller.UpdateProfile)
		profileRouter.DELETE("/", profile_controller.DeleteProfile)

	}
}
