package config

import (
	"context"
	"log"
	"time"

	"blog/database"
	"blog/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func NewMongoDatabase(env *Env) database.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongodbURI := env.MONGO_URI

	client, err := database.NewClient(mongodbURI)
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

func CloseMongoDBConnection(client database.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
func CreateRootUser(client *database.Client, env *Env) error {
	// Get the user collection
	collection := (*client).Database(env.DBName).Collection(domain.CollectionUser)

	// Check if the root user exists
	var user domain.User
	filter := bson.M{"username": env.RootUsername}
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err == nil {
		return nil
	}

	// Create the root user
	rootUser := domain.User{
		ID:       primitive.NewObjectID(),
		Username: env.RootUsername,
		Password: env.RootPassword,
		Role:     "root",
	}

	// Hash the password
	bytes, err := bcrypt.GenerateFromPassword([]byte(rootUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	rootUser.Password = string(bytes)

	// Insert the root user
	_, err = collection.InsertOne(context.Background(), rootUser)
	if err != nil {
		return err
	}

	log.Println("Root user created!")
	return nil
}
