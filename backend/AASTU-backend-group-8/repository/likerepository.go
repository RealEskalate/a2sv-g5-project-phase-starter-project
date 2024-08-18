package repository

import (
	"context"
	"meleket/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikeRepository struct {
	db *mongo.Collection
}

func NewLikeRepository(db *mongo.Database) *LikeRepository {
	return &LikeRepository{
		db: db.Collection("likes"),
	}
}

type LikeRepositoryInterface interface {
	AddLike(like *domain.Like) error
	RemoveLike(likeID primitive.ObjectID) error
	GetLikesByBlogID(blogID primitive.ObjectID) ([]domain.Like, error)
}

func (r *LikeRepository) AddLike(like *domain.Like) error {
	_, err := r.db.InsertOne(context.TODO(), like)
	return err
}

func (r *LikeRepository) RemoveLike(likeID primitive.ObjectID) error {
	_, err := r.db.DeleteOne(context.TODO(), bson.M{"_id": likeID})
	return err
}

func (r *LikeRepository) GetLikesByBlogID(blogID primitive.ObjectID) ([]domain.Like, error) {
	var likes []domain.Like
	filter := bson.M{"blog_id": blogID}
	cursor, err := r.db.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var like domain.Like
		cursor.Decode(&like)
		likes = append(likes, like)
	}
	return likes, nil
}
