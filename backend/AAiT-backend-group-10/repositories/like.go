package repositories

import (
	"context"
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
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
func (l *LikeRepository) GetLike(blogID uuid.UUID, reacterID uuid.UUID) (*domain.Like, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{
		{Key: "blog_id", Value: blogID},
		{Key: "reacter_id", Value: reacterID},
	}
	var like *domain.Like
	err := l.Collection.FindOne(ctx, filter).Decode(&like)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrLikeNotFound
		}
		return nil, domain.ErrLikeCountFetchFailed
	}
	return like, nil
}

// LikeBlog implements usecases.LikeUsecaseInterface.
func (l *LikeRepository) AddLike(like domain.Like) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := l.Collection.InsertOne(ctx, like)
	if err != nil {
		return domain.ErrLikeCreationFailed
	}
	return nil
}

func (l *LikeRepository) UpdateLike(like domain.Like) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{
		{Key: "blog_id", Value: like.BlogID},
		{Key: "reacter_id", Value: like.ReacterID},
	}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "is_like", Value: like.IsLike}}}}
	_, err := l.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.ErrLikeUpdateFailed
	}
	return nil
}

// DeleteBlog implements usecases.LikeUsecaseInterface.
func (l *LikeRepository) DeleteLike(like dto.UnlikeDto) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{
		{Key: "blog_id", Value: like.BlogID},
		{Key: "reacter_id", Value: like.ReacterID},
	}
	result, err := l.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return domain.ErrLikeDeletionFailed
	} else if result.DeletedCount == 0 {
		return domain.ErrLikeNotFound
	}
	return nil
}

func (l *LikeRepository) BlogLikeCount(blog_id uuid.UUID, isLike bool) (int, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{
		{Key: "blog_id", Value: blog_id},
		{Key: "is_like", Value: isLike},
	}
	count, err := l.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, domain.ErrLikeCountFetchFailed
	}
	return int(count), nil
}

func (l *LikeRepository) DeleteLikesByBlog(blogID uuid.UUID) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{
		{Key: "blog_id", Value: blogID},
	}
	_, err := l.Collection.DeleteMany(ctx, filter)
	if err != nil {
		return domain.ErrLikeDeletionFailed
	}
	return nil
}
