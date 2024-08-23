package main

import (
	config "github.com/aait.backend.g5.main/backend/Config"
	routers "github.com/aait.backend.g5.main/backend/Delivery/Routers"
	"github.com/gin-gonic/gin"
)

func main() {
	env := config.NewEnv()
	client := config.ConnectDB(env)
	db := config.GetDatabase(client, env)

	gin := gin.Default()

	routers.Setup(env, db, gin)
	gin.Run(env.SERVER_ADDRESS)
	// gin.Run("0.0.0.0:8081")

}
