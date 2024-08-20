package routers

import (
	"blogapp/Delivery/controllers"
	"blogapp/Infrastructure/auth_middleware"
	"blogapp/Repositories"
	usecases "blogapp/UseCases"
)

func AuthRouter() {
	authRouter := Router.Group("")
	{
		userRepo := Repositories.NewUserRepository(BlogCollections.Users, BlogCollections.RefreshTokens)

		// generate new auth repo
		authrepo := Repositories.NewAuthRepository(BlogCollections.Users, BlogCollections.RefreshTokens)
		authusecase := usecases.NewAuthUseCase(authrepo)
		authcontroller := controllers.NewAuthController(authusecase, userRepo)

		// register
		authRouter.POST("/auth/register", authcontroller.Register)
		//login
		authRouter.POST("/auth/login", authcontroller.Login)

		// oauth login with google
		authRouter.GET("/auth/login/google", authcontroller.LoginHandlerGoogle)
		authRouter.GET("/callback", authcontroller.CallbackHandler)

		//logout
		authRouter.GET("/auth/logout", auth_middleware.AuthMiddleware(), authcontroller.Logout)
		// forget password
		authRouter.POST("/auth/forget-password", authcontroller.ForgetPassword)
		authRouter.GET("/auth/forget-password/:reset_token", authcontroller.ForgetPasswordForm)
		authRouter.POST("/auth/forget-password/:reset_token", authcontroller.ResetPassword)
	}
}
