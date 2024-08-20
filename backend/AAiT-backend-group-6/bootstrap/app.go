package bootstrap

import "AAiT-backend-group-6/mongo"

type Application struct {
	Env   *Env
	Mongo mongo.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	print("me")
	app.Mongo = NewMongoDatabase(app.Env)
	print("here")
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
