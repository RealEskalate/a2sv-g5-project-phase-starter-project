package routers

import (
	"blogapp/Delivery/controllers"
	"blogapp/Infrastructure/auth_middleware"
	"blogapp/Repositories"
	usecases "blogapp/UseCases"
)

func AuthRouter() {
	authRouter := Router.Group("/auth")
	{

		// generate new auth repo
		authrepo := Repositories.NewAuthRepository(BlogCollections.Users, BlogCollections.RefreshTokens)
		authusecase := usecases.NewAuthUseCase(authrepo)
		authcontroller := controllers.NewAuthController(authusecase)

		// register
		authRouter.POST("/register", authcontroller.Register)
		//login
		authRouter.POST("/login", authcontroller.Login)
		//logout
		authRouter.POST("/logout",auth_middleware.AuthMiddleware(), authcontroller.Logout)

	}
}
