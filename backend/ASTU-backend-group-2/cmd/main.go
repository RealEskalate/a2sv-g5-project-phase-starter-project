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
	app := bootstrap.App()

	env := app.Env

	// Initialize the OAuth
	auth.NewAuth(env)

	db := app.Mongo.Database(env.DBName)

	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)

}
