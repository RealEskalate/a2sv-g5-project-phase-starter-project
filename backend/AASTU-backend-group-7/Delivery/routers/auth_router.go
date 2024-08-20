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
		userRepo := Repositories.NewUserRepository(BlogCollections.Users, BlogCollections.RefreshTokens)

		// generate new auth repo
		authrepo := Repositories.NewAuthRepository(BlogCollections.Users, BlogCollections.RefreshTokens, userRepo)
		authusecase := usecases.NewAuthUseCase(authrepo)
		authcontroller := controllers.NewAuthController(authusecase)

		// register
		authRouter.POST("/register", authcontroller.Register)
		//login
		authRouter.POST("/login", authcontroller.Login)

		// oauth login with google
		authRouter.GET("/login/google", authcontroller.LoginHandlerGoogle)
		authRouter.GET("/callback", authcontroller.CallbackHandler)

		//logout
		authRouter.GET("/logout", auth_middleware.AuthMiddleware(), authcontroller.Logout)
		// forget password
		authRouter.POST("/forget-password", authcontroller.ForgetPassword)
		authRouter.GET("/forget-password/:reset_token", authcontroller.ForgetPasswordForm)
		authRouter.POST("/forget-password/:reset_token", authcontroller.ResetPassword)

		// activate account
		authRouter.GET("/activate/:activation_token", authcontroller.ActivateAccount)

	}
}
