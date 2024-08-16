package main

import (
	"blogs/Delivery/routers"
	infrastructure "blogs/Infrastructure"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	config, err := infrastructure.LoadEnv()
	if err != nil {
		fmt.Print("Error in env.load")
	}
	// fmt.Print(config)
	DB, client, err := infrastructure.ConnectDB(config.DatabaseUrl, config.Dbname)

	if err != nil {
		fmt.Print("Error in connectDB")
	}
	defer infrastructure.CloseDB(client)
	routers.Router(server, config, DB)
	server.Run(fmt.Sprintf(":%d", config.Port))

}
