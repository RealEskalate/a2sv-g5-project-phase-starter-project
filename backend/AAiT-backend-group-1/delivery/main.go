package main

import (
	"log"
	"os"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/delivery/controllers"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/usecases"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	// userRepo := repository.NewUserRepository()
	// userUC := usecases.NewUserUseCase(userRepo)
	// userController := delivery.NewUserController(userUC)
	// userController.Route(r)

	// blogRepo := repository.NewBlogRepository()
	// blogUC := usecases.NewBlogUseCase(blogRepo)
	// blogController := delivery.NewBlogController(blogUC)
	// blogController.Route(r)

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	geminiApiKey := os.Getenv("GEMINI_API_KEY")

	blogAssistantUC := usecases.NewBlogAssistantUsecase(geminiApiKey)
	blogAssistantController := controllers.NewBlogAssistantController(blogAssistantUC)
	
	r.POST("/generate-blog", blogAssistantController.GenerateBlog)
	r.POST("/enhance-blog", blogAssistantController.EnhanceBlog)
	r.GET("/suggest-blog", blogAssistantController.SuggestBlog)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}