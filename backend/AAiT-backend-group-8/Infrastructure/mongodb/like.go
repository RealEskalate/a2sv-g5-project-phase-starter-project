package mongodb

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

func (repo *LikeRepository) LikeBlog(likeObj domain.Like) error {
	_, err := repo.collection.InsertOne(repo.ctx, likeObj)
	if err != nil {
		return errors.New("failed to like blog")
	}
	return nil
}

func (repo *LikeRepository) UnlikeBlog(likeID primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: likeID}}
	result, err := repo.collection.DeleteOne(repo.ctx, filter)
	if err != nil {
		return errors.New("error deleting like")
	}
	if result.DeletedCount == 0 {
		return errors.New("like not found")
	}
	return nil
}

func (repo *LikeRepository) GetLikes(blogID primitive.ObjectID) ([]domain.Like, error) {
	var likes []domain.Like
	filter := bson.D{{Key: "blog_id", Value: blogID}}
	cursor, err := repo.collection.Find(repo.ctx, filter)
	if err != nil {
		return nil, errors.New("error reading likes")
	}
	defer cursor.Close(repo.ctx)
	for cursor.Next(repo.ctx) {
		var like domain.Like
		cursor.Decode(&like)
		likes = append(likes, like)
	}
	if len(likes) == 0 {
		return nil, errors.New("no likes found")
	}
	return likes, nil
}

func (repo *LikeRepository) CheckIfLiked(userID, blogID primitive.ObjectID) (bool, domain.Like) {
	filter := bson.D{{Key: "user_id", Value: userID}, {Key: "blog_id", Value: blogID}}
	var like domain.Like
	err := repo.collection.FindOne(repo.ctx, filter).Decode(&like)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, domain.Like{}
		}
	}
	return true, like
}

func (repo *LikeRepository) DeleteByBLogID(blogID primitive.ObjectID) error {
	filter := bson.D{{Key: "blog_id", Value: blogID}}
	result, err := repo.collection.DeleteMany(repo.ctx, filter)
	if err != nil {
		return errors.New("error deleting comments")
	}
	if result.DeletedCount == 0 {
		return errors.New("no comments found")
	}
	return nil
}

func (repo *LikeRepository) DropDataBase() error {
	filter := bson.M{}
	_, err := repo.collection.DeleteMany(repo.ctx, filter)
	if err != nil {
		return nil
	}
	return nil
}
