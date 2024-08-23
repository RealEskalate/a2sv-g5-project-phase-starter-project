package mongodb

import (
	"AAiT-backend-group-8/Domain"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CacheRepository struct {
	db *mongo.Collection
}

func NewCacheRepository(db *mongo.Collection) *CacheRepository {
	return &CacheRepository{
		db: db,
	}
}

func (cache *CacheRepository) Update(id primitive.ObjectID, value string) error {
	findOneResult := cache.db.FindOne(context.Background(), bson.M{"_id": id})

	if findOneResult.Err() != nil {
		cache.db.InsertOne(context.Background(), bson.M{"_id": id, "keys": []string{}})
	}

	update := bson.M{"$push": bson.M{"keys": value}}
	filter := bson.M{"_id": id}
	_, err := cache.db.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}
	return nil
}

func (cache *CacheRepository) Delete(id primitive.ObjectID) ([]string, error) {
	filter := bson.M{"_id": id}

	var id2key Domain.IDToKeys
	findResult := cache.db.FindOne(context.Background(), filter)

	if findResult.Decode(&id2key) != nil {
		fmt.Println("jojjknh")
		return []string{}, errors.New("not found")
	}
	_, err := cache.db.DeleteOne(context.Background(), filter)
	if err != nil {
		fmt.Println(err.Error())
		return []string{}, err
	}

	fmt.Println("reached here")
	return id2key.Keys, nil
}
