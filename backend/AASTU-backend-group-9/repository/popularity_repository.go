package repository

import (
	"blog/database"
	"blog/domain"
	"context"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PopularityRepositoryStruct struct {
	popdb      database.Database
	collection string
}

// hasUserLiked implements domain.PopularityRepository.

func NewPopularityRepository(db database.Database, collection string) domain.PopularityRepository {
	return &PopularityRepositoryStruct{
		popdb:      db,
		collection: collection,
	}
}

func (pr *PopularityRepositoryStruct) HasUserLiked(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (bool, error) {
	filter := bson.M{"post_id": id, "user_id": userID, "interaction_type": "Like"}
	collection := pr.popdb.Collection(pr.collection)
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (pr *PopularityRepositoryStruct) HasUserDisliked(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (bool, error) {
	filter := bson.M{"post_id": id, "user_id": userID, "interaction_type": "Dislike"}
	collection := pr.popdb.Collection(pr.collection)
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (pr *PopularityRepositoryStruct) UserInteractionsAdder(ctx context.Context, user domain.UserInteraction) error {
	collection := pr.popdb.Collection(pr.collection)
	_, err := collection.InsertOne(ctx, user)
	return err

}

func (pr *PopularityRepositoryStruct) UserInteractionsDelete(ctx context.Context, user domain.UserInteraction) error {
	collection := pr.popdb.Collection(pr.collection)
	_,err:= collection.DeleteOne(ctx, user)
	return err
}
