package repository

import (
	domain "AAiT-backend-group-8/Domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikeRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewLikeRepository(collection *mongo.Collection, ctx context.Context) *LikeRepository {
	return &LikeRepository{
		collection: collection,
		ctx:        ctx,
	}
}

func (l *LikeRepository) LikeBlog(likeObj domain.Like) error {
	_, err := l.collection.InsertOne(l.ctx, likeObj)
	if err != nil {
		return errors.New("failed to like blog")
	}
	return nil
}

func (l *LikeRepository) UnlikeBlog(likeID primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: likeID}}
	result, err := l.collection.DeleteOne(l.ctx, filter)
	if err != nil {
		return errors.New("error deleting like")
	}
	if result.DeletedCount == 0 {
		return errors.New("like not found")
	}
	return nil
}

func (l *LikeRepository) GetLikes(blogID primitive.ObjectID) ([]domain.Like, error) {
	var likes []domain.Like
	filter := bson.D{{Key: "blogid", Value: blogID}}
	cursor, err := l.collection.Find(l.ctx, filter)
	if err != nil {
		return nil, errors.New("error reading likes")
	}
	defer cursor.Close(l.ctx)
	for cursor.Next(l.ctx) {
		var like domain.Like
		cursor.Decode(&like)
		likes = append(likes, like)
	}
	if len(likes) == 0 {
		return nil, errors.New("no likes found")
	}
	return likes, nil
}

func (l *LikeRepository) CheckIfLiked(userID, blogID primitive.ObjectID) (bool, domain.Like) {
	filter := bson.D{{Key: "userid", Value: userID}, {Key: "blogid", Value: blogID}}
	var like domain.Like
	l.collection.FindOne(l.ctx, filter).Decode(&like)
	if like.Id == primitive.NewObjectID() {
		return false, domain.Like{}
	}
	return true, like
}
