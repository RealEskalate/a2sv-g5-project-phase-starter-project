package repositories

import (
	"context"
	"time"

	"aait.backend.g10/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikeRepository struct {
	Collection *mongo.Collection
}

func NewLikeRepository(db *mongo.Database, collectionName string) *LikeRepository {
	collection := db.Collection(collectionName)
	return &LikeRepository{
		Collection: collection,
	}
}

// LikeBlog implements usecases.LikeUsecaseInterface.
func (l *LikeRepository) LikeBlog(like domain.Like) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{
		{Key: "blog_id", Value: like.BlogID},
		{Key: "user_id", Value: like.UserID},
	}

	var existingLike domain.Like
	err := l.Collection.FindOne(ctx, filter).Decode(&existingLike)
	if err == mongo.ErrNoDocuments {
		// Insert a new document
		like.ID = uuid.New()
		_, err = l.Collection.InsertOne(ctx, like)
	} else if err == nil {
		// Update the existing document
		update := bson.D{{Key: "$set", Value: bson.D{{Key: "is_like", Value: like.IsLike}}}}
		_, err = l.Collection.UpdateOne(ctx, filter, update)
	}
	return err
}

// DeleteBlog implements usecases.LikeUsecaseInterface.
func (l *LikeRepository) DeleteLike(like domain.Like) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{
		{Key: "blog_id", Value: like.BlogID},
		{Key: "user_id", Value: like.UserID},
	}
	_, err := l.Collection.DeleteOne(ctx, filter)
	return err
}

func (l *LikeRepository) BlogLikeCount(blog_id uuid.UUID) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{
		{Key: "blog_id", Value: blog_id},
	}
	count, err := l.Collection.CountDocuments(ctx, filter)
	return int(count), err
}
