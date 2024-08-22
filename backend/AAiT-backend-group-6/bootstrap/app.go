package bootstrap

import (
	"AAiT-backend-group-6/mongo"

	"AAiT-backend-group-6/redis"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
	Redis redis.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	// app.Redis = redis.NewClient(app.Env.REDISADDR)
	return *app
}

func (app *Application) Close() {
	CloseMongoDBConnection(app.Mongo)
	// app.Redis.Close()
}
