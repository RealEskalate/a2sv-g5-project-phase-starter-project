package repository

import (
	"context"
	"errors"

	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LikeRepository struct {
	collection database.CollectionInterface
}

func NewLikeRepository(collection database.CollectionInterface) *LikeRepository {
	return &LikeRepository{
		collection: collection,
	}
}

func (LR *LikeRepository) GetLikes(post_id string) ([]domain.Like, error) {
	var likes []domain.Like
	query := bson.M{"post_id": post_id}

	cursor, err := LR.collection.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var like domain.Like
		if err := cursor.Decode(&like); err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}

	return likes, nil
}

func (LR *LikeRepository) CreateLike(user_id string, post_id string) error {
	userObjectID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return err
	}

	postObjectID, err := primitive.ObjectIDFromHex(post_id)
	if err != nil {
		return err
	}
	like := domain.Like{
		ID:     primitive.NewObjectID(),
		UserID: userObjectID,
		PostID: postObjectID,
	}

	_, err = LR.collection.InsertOne(context.TODO(), like)
	return err
}

func (LR *LikeRepository) RemoveLike(user_id string, post_id string) error {
	userObjectID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return err
	}

	postObjectID, err := primitive.ObjectIDFromHex(post_id)
	if err != nil {
		return err
	}

	query := bson.M{"user_id": userObjectID, "post_id": postObjectID}
	res, err := LR.collection.DeleteOne(context.TODO(), query)
	if err != nil {
		return err
	}

	if res.DeletedCount() == 0 {
		return errors.New("like not found")
	}

	_, err = LR.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": postObjectID},
		bson.M{"$inc": bson.M{"like_count": -1}},
	)

	return err
}

func (LR *LikeRepository) ToggleLike(user_id string, post_id string) error {
	userObjectID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return err
	}

	postObjectID, err := primitive.ObjectIDFromHex(post_id)
	if err != nil {
		return err
	}

	dislikeRepo := NewDislikeRepository(LR.collection)

	count, err := LR.collection.CountDocuments(context.TODO(), bson.M{"user_id": userObjectID, "post_id": postObjectID})
	if err != nil {
		return err
	}

	if count > 0 {
		return LR.RemoveLike(user_id, post_id)
	} else {
		if err := dislikeRepo.RemoveDislike(user_id, post_id); err != nil && err.Error() != "dislike not found" {
			return err
		}
		return LR.CreateLike(user_id, post_id)
	}
}
