package bootstrap

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/option"
)

type Application struct {
	Env           *Env
	Mongo         *mongo.Client
	GenAi         *genai.GenerativeModel
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)

	client, err := genai.NewClient(context.TODO(), option.WithAPIKey(app.Env.GeminiApiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	model.ResponseMIMEType = "application/json"
	app.GenAi = model

	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
