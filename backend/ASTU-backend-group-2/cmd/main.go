package main

import (
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/route"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/validators"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if ginValidator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		ginValidator.RegisterValidation("StrongPassword", validators.StrongPassword)
	}
}
func main() {
	// Initialize the app
	app := bootstrap.App()

	// Get the environment variables
	env := app.Env

	// Initialize the OAuth
	auth.NewAuth(env)

	// Connect to the database
	db := app.Mongo.Database(env.DBName)

	// Close the database connection when the main function is done
	defer app.CloseDBConnection()

	// Set the timeout for the context of the request
	timeout := time.Duration(env.ContextTimeout) * time.Second

	// Initialize the gin
	gin := gin.Default()

	// Setup the routes
	route.Setup(env, timeout, db, gin, app.Cloudinary)

	// Run the server
	gin.Run(env.ServerAddress)

}
