package main

import config "github.com/aait.backend.g5.main/backend/Config"



func main() {

	// Initialize MongoDB connection
	client := config.ConnectDB()
	db := config.GetDatabase(client)

		
}
