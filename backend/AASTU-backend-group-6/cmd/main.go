package main

import (
	"blogs/Delivery/routers"
	infrastructure "blogs/Infrastructure"
	"fmt"

	xss "github.com/araujo88/gin-gonic-xss-middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	 var xssMdlwr xss.XssMw
    server.Use(xssMdlwr.RemoveXss())

	


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
