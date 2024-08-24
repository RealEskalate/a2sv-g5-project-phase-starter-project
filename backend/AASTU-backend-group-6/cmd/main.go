package main

import (
	"blogs/Delivery/routers"
	infrastructure "blogs/Infrastructure"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	confi := cors.DefaultConfig()
	confi.AllowAllOrigins = true

	server.Use(cors.New(confi))

	config, err := infrastructure.LoadEnv()
	if err != nil {
		fmt.Print("Error in env.load")
	}
	// fmt.Print(config)
	DB, client, err := infrastructure.ConnectDB(config.DatabaseUrl, config.DbName)
	fmt.Println("-----", config.GeminiAPIKey)
	if err != nil {
		fmt.Print("Error in connectDB")
	}
	defer infrastructure.CloseDB(client)
	routers.Router(server.Group("api/v0"), config, DB)
	server.Run(fmt.Sprintf(":%d", config.Port))

}
