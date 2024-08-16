package main

import (
	Config "blogapp/Config"
	"blogapp/Delivery/routers"
	"context"
	"log"
)

func main() {
	// Connect to the database
	client := Config.ConnectDB()

	// Defer the closing of the database
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// Setuprouter()
	r := routers.Setuprouter(client)
	if r != nil {
		r.Run()
	} else {
		log.Fatal("Failed to start server")
	}
}
