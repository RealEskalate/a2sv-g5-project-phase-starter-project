package bootstrap

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Env        *Env
	Mongo      *mongo.Client
	Cloudinary *cloudinary.Cloudinary
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	app.Cloudinary = NewCloudinaryService(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
