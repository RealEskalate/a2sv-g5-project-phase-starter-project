package main

import (
	"blogs/bootstrap"
	"blogs/delivery/router"
)

func main() {
	var uri = "mongodb://localhost:27017"

	client, err := bootstrap.ConnectToMongoDB(uri)

	if err != nil {
		panic(err)
	}

	defer bootstrap.DisconnectFromMongoDB(client)

	router := router.SetupRouter(client)
	router.Run(":8080")
}
