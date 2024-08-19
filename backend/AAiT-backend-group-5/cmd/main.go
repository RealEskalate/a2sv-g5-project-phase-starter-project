package main

import config "github.com/aait.backend.g5.main/backend/Config"



func main() {
	env := config.NewEnv()
	client := config.ConnectDB(env)
	db := config.GetDatabase(client, env)

		
}
