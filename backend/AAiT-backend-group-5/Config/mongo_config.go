package config

import (
	"context"
	"log"
	"time"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"

)


func ConnectDB(env *Env) interfaces.Client {
	mongoURI := env.MONGO_URI
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in .env file")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := interfaces.NewClient(mongoURI)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetDatabase(client interfaces.Client, env *Env) interfaces.Database {
	dbName := env.DB_NAME
	if dbName == "" {
		log.Fatal("DB_NAME is not set in .env file")
	}
	return client.Database(dbName)
}
