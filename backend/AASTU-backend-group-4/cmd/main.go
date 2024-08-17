package main

import (
	"blog-api/infrastructure/bootstrap"
	"fmt"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	_ = app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()
	fmt.Println("Work correctly")
}
