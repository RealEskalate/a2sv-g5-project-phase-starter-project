package main

import (
	"blog_api/delivery/env"
	"blog_api/delivery/router"
	initdb "blog_api/infrastructure/db"
	"fmt"
	"log"
)

func main() {
	err := env.LoadEnvironmentVariables()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	client, err := initdb.ConnectDB(env.ENV.DB_ADDRESS, env.ENV.DB_NAME)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	database := client.Database(env.ENV.DB_NAME)

	err = initdb.CreateRootUser(database, env.ENV.ROOT_USERNAME, env.ENV.ROOT_PASSWORD)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	defer initdb.DisconnectDB(client)
    fmt.Println(env.ENV.ROUTE_PREFIX)

	router.SetupRouter(env.ENV.PORT, env.ENV.ROUTE_PREFIX, database)
}
