package main

import (
	"AAiT-backend-group-6/bootstrap"
	"AAiT-backend-group-6/delivery/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)

	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Application panicked: %v", r)
		}
		app.CloseDBConnection()
	}()

	log.Println("Starting the application...")

	// timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	log.Println("Setting up routers...")
	routers.SetupRouter(db, gin)

	log.Println("Running the server...")
	err := gin.Run(env.ServerAddress)
	if err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
