package main

import (
	"backend-starter-project/bootstrap"
)

func main()  {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	// Do something with the database
	_ = db
}