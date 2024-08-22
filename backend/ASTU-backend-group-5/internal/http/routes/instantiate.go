package routes

import (
	"blogApp/internal/http/handlers/account"
	BlogHandler "blogApp/internal/http/handlers/blog"
	"blogApp/internal/repository/mongodb"
	"blogApp/internal/usecase"
	"blogApp/internal/usecase/blog"
	"blogApp/internal/usecase/user"

	"go.mongodb.org/mongo-driver/mongo"
)

func InstantaiteUserHandler(collection *mongo.Collection) *account.UserHandler {
	userRepo := &mongodb.UserRepositoryMongo{Collection: collection}
	userUsecase := user.NewUserUsecase(userRepo)
	userHandler := account.NewUserHandler(userUsecase)
	return userHandler
}

func InstantaiteTokenUsecase(collection *mongo.Collection) *usecase.TokenUsecase {
	tokenRepo := mongodb.NewMongoTokenRepository(collection)
	tokenUsecase := usecase.NewTokenUsecase(tokenRepo)
	return tokenUsecase
}

func InstantaiteBlogHandler(blogsCollection, commentsCollection, likesCollection, viewsCollection, tagsCollection *mongo.Collection) *BlogHandler.BlogHandler {
	blogRepo := mongodb.NewMongoBlogRepository(blogsCollection, commentsCollection, likesCollection, viewsCollection, tagsCollection)
	blogUsecase := blog.NewBlogUseCase(blogRepo)
	blogHandler := BlogHandler.NewBlogHandler(blogUsecase)
	return blogHandler
}
