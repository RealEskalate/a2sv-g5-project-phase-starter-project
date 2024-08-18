package initdb

import (
	"blog_api/domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
Accepts a connection string and a database name and returns a pointer to a mongo.Client
once it has successfully connected to the DB and called the SetupIndicies function without
any errors.
*/
func ConnectDB(connectionString string, databaseName string) (*mongo.Client, error) {
	if connectionString == "" {
		return nil, fmt.Errorf("error: DB connection string not found. Make sure the environment variables are set correctly")
	}

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err.Error())
	}

	// ping DB client to verify connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err.Error())
	}

	db := client.Database(databaseName)
	err = SetupIndicies(db)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err.Error())
	}

	return client, nil

}

/* Accepts a pointer to a mongo.Database and creates unique indices for the email and username fields */
func SetupIndicies(db *mongo.Database) error {
	_, err := db.Collection(domain.CollectionUsers).Indexes().CreateOne(context.TODO(), mongo.IndexModel{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetUnique(true)})
	if err != nil {
		return fmt.Errorf("\n\n Error " + err.Error())
	}

	_, err = db.Collection(domain.CollectionUsers).Indexes().CreateOne(context.TODO(), mongo.IndexModel{Keys: bson.D{{Key: "username", Value: 1}}, Options: options.Index().SetUnique(true)})
	if err != nil {
		return fmt.Errorf("\n\n Error " + err.Error())
	}

	return nil
}

/* Accepts a pointer to a mongo.Client and disconnects the client from the DB */
func DisconnectDB(client *mongo.Client) {
	client.Disconnect(context.Background())
}
