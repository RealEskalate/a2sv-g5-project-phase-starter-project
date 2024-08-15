package main

import (
	"blogapp/Delivery/routers"
	"log"
)

func main() {
	// Setuprouter()
	r := routers.Setuprouter()
	if r != nil {
		r.Run()
	} else {
		log.Fatal("Failed to start server")
	}
}
