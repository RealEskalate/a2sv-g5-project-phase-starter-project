package main

import (
	"astu-backend-g1/config"
	"astu-backend-g1/delivery/controllers"
	_ "astu-backend-g1/delivery/docs"
	"astu-backend-g1/delivery/routers"
	"astu-backend-g1/gemini"
	"astu-backend-g1/infrastructure"
	"astu-backend-g1/repository"
	usecase "astu-backend-g1/usecases"
	"context"
	"log"

	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title TODO APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @securityDefinitions.apiKey JWT
// @in header
// @name token
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
// @schemes http
func main() {
	config_mongo, err := config.LoadConfig()
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	clientOptions := options.Client().ApplyURI(config_mongo.Database.Uri)
	if err != nil {
		log.Fatal(err)
	}
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	blogCollections := client.Database("Starter_Blog_Api").Collection("Blogs")
	userCollections := client.Database("Starter_Blog_Api").Collection("Users")
	commentCollections := client.Database("Starter_Blog_Api").Collection("Comments")
	replyCollections := client.Database("Starter_Blog_Api").Collection("Replies")
	// _ = client.Database("BlogAPI").Collection("Tokens")
	blogRepo := repository.NewBlogRepository(mongoifc.WrapClient(client), mongoifc.WrapCollection(blogCollections), mongoifc.WrapCollection(commentCollections), mongoifc.WrapCollection(replyCollections))
	blogUsecase := usecase.NewBlogUsecase(blogRepo)
	blogController := controllers.NewBlogController(*blogUsecase)
	authController := infrastructure.NewAuthController(blogRepo)
	userRepo := repository.NewUserRepository(mongoifc.WrapCollection(userCollections))
	userUsecase, err := usecase.NewUserUsecase(userRepo)
	if err != nil {
		panic(err)
	}
	prompts, err := infrastructure.LoadPrompt("./prompts.json")
	if err != nil {
		log.Fatal(err)
	}
	model, err := gemini.NewGeminiModel(config_mongo.Gemini.ApiKey, config_mongo.Gemini.Model, prompts)
	if err != nil {
		log.Fatal(err)
	}
	aiController := controllers.NewAIController(model, *blogUsecase)
	UserController := controllers.NewUserController(userUsecase)
	Router := routers.NewMainRouter(*UserController, *blogController, authController, *aiController)
	Router.GinBlogRouter()
}
