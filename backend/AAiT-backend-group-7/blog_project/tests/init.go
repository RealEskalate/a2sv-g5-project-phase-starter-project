package tests

// import (
// 	"blog_project/delivery/controllers"
// 	"blog_project/domain"
// 	"blog_project/infrastructure"
// 	"blog_project/repositories"
// 	"blog_project/usecases"
// 	"context"
// 	"os"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var (
// 	UserRepo       domain.IUserRepository
// 	BlogRepo       domain.IBlogRepository
// 	TokenRepo      domain.ITokenRepository
// 	UserUsecase    domain.IUserUsecase
// 	BlogUsecase    domain.IBlogUsecase
// 	BlogController domain.IBlogController
// 	UserController domain.IUserController
// )

// func init() {
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	db := client.Database("test_blog_project")
// 	userCollection := db.Collection("users")
// 	blogCollection := db.Collection("blogs")
// 	tokenCollection := db.Collection("tokens")
// 	UserRepo = repositories.NewUserRepository(userCollection)
// 	BlogRepo = repositories.NewBlogRepository(blogCollection)
// 	TokenRepo = repositories.NewTokenRepository(tokenCollection)
// 	UserUsecase = usecases.NewUserUsecase(UserRepo, TokenRepo)
// 	aiService := infrastructure.NewGeminiService(os.Getenv("GEMINI_API_KEY"))
// 	BlogUsecase = usecases.NewBlogUsecase(aiService, BlogRepo, UserUsecase)
// 	BlogController = controllers.NewBlogController(BlogUsecase)
// 	UserController = controllers.NewUserController(UserUsecase)

// }
