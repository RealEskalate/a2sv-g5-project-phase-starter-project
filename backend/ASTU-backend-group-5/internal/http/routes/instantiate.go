package routes

import (
	"blogApp/internal/http/handlers/account"
	BlogHandler "blogApp/internal/http/handlers/blog"
	"blogApp/internal/repository/mongodb"
	"blogApp/internal/usecase"
	"blogApp/internal/usecase/blog"
	"blogApp/internal/usecase/user"

	localMongo "blogApp/pkg/mongo"
)

func InstantaiteUserHandler() *account.UserHandler {
	usersCollection := localMongo.UserCollection
	userRepo := &mongodb.UserRepositoryMongo{Collection: usersCollection}
	userUsecase := user.NewUserUsecase(userRepo)
	userHandler := account.NewUserHandler(userUsecase)
	return userHandler
}

func InstantaiteTokenUsecase() *usecase.TokenUsecase {
	tokensCollection := localMongo.TokenCollection
	tokenRepo := mongodb.NewMongoTokenRepository(tokensCollection)
	tokenUsecase := usecase.NewTokenUsecase(tokenRepo)
	return tokenUsecase
}

func InstantaiteBlogHandler() *BlogHandler.BlogHandler {
	blogsCollection := localMongo.BlogsCollection
	commentsCollection := localMongo.CommentsCollection
	likesCollection := localMongo.LikesCollection
	viewsCollection := localMongo.ViewsCollection
	tagsCollection := localMongo.TagsCollection

	blogRepo := mongodb.NewMongoBlogRepository(blogsCollection, commentsCollection, likesCollection, viewsCollection, tagsCollection)

	userCollection := localMongo.UserCollection
	userRepo := mongodb.NewUserRepositoryMongo(userCollection)

	blogUsecase := blog.NewBlogUseCase(blogRepo, userRepo)
	blogHandler := BlogHandler.NewBlogHandler(blogUsecase)
	return blogHandler
}
