package bootstrap

import (
	"context"
	"log"

	"backend-starter-project/mongo"

	"github.com/go-redis/redis/v8"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Application struct {
	Env           *Env
	Mongo         *mongo.Client
	GenAi         *genai.GenerativeModel
	Redis 		  *redis.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)

	app.Redis = NewRedisClient(app.Env)
	app.GenAi = NewAiModel(app.Env)

	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}


func (app *Application) CloseModelClient(){
	//TODO: close the client for the ai model
}

func (app *Application) CloseRedisConnection() {
	CloseRedisConnection(app.Redis)
}