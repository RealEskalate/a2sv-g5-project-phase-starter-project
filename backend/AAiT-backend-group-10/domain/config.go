package domain

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// setup connection
var Collections *mongo.Collection
var Database *mongo.Database
var UserCollection *mongo.Collection

func ConnectDB() *mongo.Database {
	var Database *mongo.Database

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var db_url = os.Getenv("DB_URI")
	defer func() {

		fmt.Println(db_url)
	}()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(db_url).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {

		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	Database = client.Database("Blog_manager")

	return Database
}

// Task struct
