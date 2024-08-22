package bootstrap

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
	Redis *redis.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	app.Redis = ConnectRedis(app.Env)
	return *app
}

func (app *Application) Close() {
	CloseMongoDatabase(&app.Mongo)
}
