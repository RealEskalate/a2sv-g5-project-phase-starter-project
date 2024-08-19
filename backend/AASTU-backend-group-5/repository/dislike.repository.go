package repository

import (
	"context"
	"errors"

	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DislikeRepository struct {
	collection database.CollectionInterface
}

func NewDislikeRepository(collection database.CollectionInterface) *DislikeRepository {
	return &DislikeRepository{
		collection: collection,
	}
}

func (DR *DislikeRepository) GetDisLikes(post_id string) ([]domain.DisLike, error) {
	var dislikes []domain.DisLike
	query := bson.M{"post_id": post_id}

	cursor, err := DR.collection.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var dislike domain.DisLike
		if err := cursor.Decode(&dislike); err != nil {
			return nil, err
		}
		dislikes = append(dislikes, dislike)
	}

	return dislikes, nil
}

func (DR *DislikeRepository) CreateDisLike(user_id string, post_id string) error {
	userObjectID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return err // Return an error if the conversion fails
	}

	postObjectID, err := primitive.ObjectIDFromHex(post_id)
	if err != nil {
		return err
	}

	dislike := domain.DisLike{
		ID:     primitive.NewObjectID(),
		UserID: userObjectID,
		PostID: postObjectID,
	}

	_, err = DR.collection.InsertOne(context.TODO(), dislike)
	return err
}

func (DR *DislikeRepository) DeleteDisLike(dislike_id string) error {
	obID, _ := primitive.ObjectIDFromHex(dislike_id)
	query := bson.M{"_id": obID}

	res, err := DR.collection.DeleteOne(context.TODO(), query)
	if err != nil {
		return err
	}

	if res.DeletedCount() == 0 {
		return errors.New("no dislike with this ID found")
	}

	return nil
}
