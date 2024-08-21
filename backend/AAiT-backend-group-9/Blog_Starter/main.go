package main

import (
	"log"
	"time"

	route "Blog_Starter/api/router"
	"Blog_Starter/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	env := config.NewEnv()

	// Initialize MongoDB client
	db := config.NewMongoDatabase(env)
	defer config.CloseMongoDBConnection(db)

	// Set up the Gin router

	router := gin.Default()
	router.Static("/upload", "./upload")

	// Global CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	router.Use(cors.New(config))

	// Set up routes
	timeout := time.Duration(env.ContextTimeout) * time.Second
	route.Setup(env, timeout, db, router)

	// Run the server
	if err := router.Run("localhost:" + env.DBPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
