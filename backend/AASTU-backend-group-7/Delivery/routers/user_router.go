package routers

import (
	"blogapp/Delivery/controllers"
	"blogapp/Infrastructure/auth_middleware"
	"blogapp/Repositories"
	usecases "blogapp/UseCases"
)

func UserRouter() {
	userRouter := Router.Group("/users")
	userRouter.Use(auth_middleware.AuthMiddleware())
	userRouter.Use(auth_middleware.IsAdminMiddleware())
	{

		// generate new auth repo
		user_repo := Repositories.NewUserRepository(BlogCollections.Users, BlogCollections.RefreshTokens)

		user_usecase := usecases.NewUserUseCase(user_repo)
		user_controller := controllers.NewUserController(user_usecase)

		// get all users
		userRouter.POST("/", user_controller.CreateUser)
		userRouter.GET("/", user_controller.GetUsers)
		userRouter.GET("/:id", user_controller.GetUser)
		userRouter.PUT("/:id", user_controller.UpdateUser)
		userRouter.DELETE("/:id", user_controller.DeleteUser)
		userRouter.PUT("/:id/promote", user_controller.PromoteUser)
		userRouter.PUT("/:id/demote", user_controller.DemoteUser)
	}
}
