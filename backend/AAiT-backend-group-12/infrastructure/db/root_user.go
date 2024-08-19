package initdb

import (
	"blog_api/domain"
	"blog_api/infrastructure/cryptography"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateRootUser(db *mongo.Database, rootUsername string, rootPassword string) error {
	rootUser := domain.User{
		Username:   rootUsername,
		Email:      "root@root.root",
		Password:   rootPassword,
		Role:       "root",
		CreatedAt:  time.Now().Round(0),
		IsVerified: true,
	}

	hashedPwd, err := cryptography.HashString(rootUser.Password)
	if err != nil {
		return fmt.Errorf("error hashing root user password: " + err.Error())
	}
	rootUser.Password = hashedPwd
	collection := db.Collection(domain.CollectionUsers)

	_, derr := collection.DeleteMany(context.Background(), bson.D{bson.E{Key: "role", Value: "root"}})
	if derr != nil {
		return fmt.Errorf("error clearing root users: " + derr.Error())
	}

	_, derr = collection.InsertOne(context.Background(), rootUser)
	if derr != nil {
		return fmt.Errorf("error creating root users: " + derr.Error())
	}

	fmt.Println("Root user created successfully")

	return nil
}
