package infrastructure

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "time"
)

func InitMongoDB(uri string) *mongo.Client {
    ctx := context.TODO()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err = mongo.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    return client
}
