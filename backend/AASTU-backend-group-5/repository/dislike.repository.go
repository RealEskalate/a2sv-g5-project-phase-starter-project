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
		return err
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

func (DR *DislikeRepository) RemoveDislike(user_id string, post_id string) error {
	userObjectID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return err
	}

	postObjectID, err := primitive.ObjectIDFromHex(post_id)
	if err != nil {
		return err
	}

	query := bson.M{"user_id": userObjectID, "post_id": postObjectID}
	res, err := DR.collection.DeleteOne(context.TODO(), query)
	if err != nil {
		return err
	}

	if res.DeletedCount() == 0 {
		return errors.New("dislike not found")
	}

	_, err = DR.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": postObjectID},
		bson.M{"$inc": bson.M{"dislike_count": -1}},
	)

	return err
}

func (DR *DislikeRepository) ToggleDislike(user_id string, post_id string) error {
	userObjectID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return err
	}

	postObjectID, err := primitive.ObjectIDFromHex(post_id)
	if err != nil {
		return err
	}

	likeRepo := NewLikeRepository(DR.collection)

	count, err := DR.collection.CountDocuments(context.TODO(), bson.M{"user_id": userObjectID, "post_id": postObjectID})
	if err != nil {
		return err
	}

	if count > 0 {
		return DR.RemoveDislike(user_id, post_id)
	} else {
		if err := likeRepo.RemoveLike(user_id, post_id); err != nil && err.Error() != "like not found" {
			return err
		}
		return DR.CreateDisLike(user_id, post_id)
	}
}
