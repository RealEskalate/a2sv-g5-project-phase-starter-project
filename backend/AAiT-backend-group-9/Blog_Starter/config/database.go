package config

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoDatabase(env *Env) *mongo.Client {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    mongodbURI := env.DBUri

    clientOptions := options.Client().ApplyURI(mongodbURI)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, readpref.Primary())
    if err != nil {
        log.Fatal(err)
    }

    return client
}

func CloseMongoDBConnection(client *mongo.Client) {
    if client == nil {
        return
    }

    err := client.Disconnect(context.TODO())
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connection to MongoDB closed.")
}