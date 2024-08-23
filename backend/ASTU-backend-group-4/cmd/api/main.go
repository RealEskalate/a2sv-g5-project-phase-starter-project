package main

import (
	"log"
	"os"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/auth"
	authMongo "github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/auth/mongodb"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/blog"
	blogMongo "github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/blog/mongodb"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/chat"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/chat/gemini_ai"
	chatMongo "github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/chat/mongodb"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/internal/db"
	controllers "github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/internal/http/gin"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getHost() string {
	hostUrl := os.Getenv("HOST_URL")
	if hostUrl != "" {
		return hostUrl
	}

	return "localhost:8000"
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	mongoClient := db.NewMongoClient()
	mongoDB := mongoClient.Database(os.Getenv("MONGO_DB"))

	authRepository := authMongo.NewAuthStorage(mongoDB.Collection("users"), mongoDB.Collection("tokens"))
	chatRepository := chatMongo.NewChatRepository(mongoDB, "chats")
	blogRepository := blogMongo.NewBlogStorage(mongoDB)
	aiService := gemini_ai.NewAIService(gemini_ai.NewModel())

	authUsecase := auth.NewAuthUserUsecase(authRepository, infrastructure.NewEmail(
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_PORT"),
	))
	chatUsecase := chat.NewUsecase(chatRepository, aiService)
	blogUsecase := blog.NewBlogUseCaseImpl(blogRepository, authRepository)

	authController := controllers.NewUserController(authUsecase)
	blogController := controllers.NewBlogController(blogUsecase)
	chatController := controllers.NewChatHandler(chatUsecase)

	controllers.SetUpRouter(r, blogController, chatController, authController)

	r.Run(getHost())
}
