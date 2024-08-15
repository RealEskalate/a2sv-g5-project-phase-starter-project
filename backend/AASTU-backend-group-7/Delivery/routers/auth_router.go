package routers

import (
	"blogapp/Repositories"
	usecases "blogapp/UseCases"
)

func AuthRouter() error {
	authRouter := Router.Group("/auth")
	{

		// generate new auth repo
		authrepo := Repositories.NewAuthRepository(BlogCollections.Users)
		usecase := usecases.NewAuthUseCase(authrepo)

	}
	return nil
}
