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

func (LR *LikeRepository) DeleteLike(like_id string) error {
	obID, _ := primitive.ObjectIDFromHex(like_id)
	query := bson.M{"_id": obID}

	res, err := LR.collection.DeleteOne(context.TODO(), query)
	if err != nil {
		return err
	}

	if res.DeletedCount() == 0 {
		return errors.New("no like with this ID found")
	}

	return nil
}
