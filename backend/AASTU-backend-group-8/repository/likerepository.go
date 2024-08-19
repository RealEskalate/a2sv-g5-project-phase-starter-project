package repository

import (
	"context"
	"meleket/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LikeRepository struct {
	collection domain.Collection
}

type LikeRepositoryInterface interface {
	AddLike(like *domain.Like) error
	RemoveLike(likeID primitive.ObjectID) error
	GetLikesByBlogID(blogID primitive.ObjectID) ([]domain.Like, error)
}

func NewLikeRepository(col domain.Collection) *LikeRepository {
	return &LikeRepository{collection: col}
}

func (r *LikeRepository) AddLike(like *domain.Like) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, like)
	return err
}

func (r *LikeRepository) RemoveLike(likeID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": likeID})
	return err
}

func (r *LikeRepository) GetLikesByBlogID(blogID primitive.ObjectID) ([]domain.Like, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var likes []domain.Like
	filter := bson.M{"blog_id": blogID}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var like domain.Like
		cursor.Decode(&like)
		likes = append(likes, like)
	}
	return likes, nil
}
