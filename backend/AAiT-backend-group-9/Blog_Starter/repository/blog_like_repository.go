package repository

import (
	"Blog_Starter/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikeRepository struct {
	database       *mongo.Database
	repoCollection string
}

func NewLikeRepository(db *mongo.Database, collection string) domain.LikeRepository {
	return &LikeRepository{
		database:       db,
		repoCollection: collection,
	}
}

// GetByID implements domain.LikeRepository.
func (lr *LikeRepository) GetByID(c context.Context, userID string, blogID string) (*domain.Like, error) {
	collection := lr.database.Collection(lr.repoCollection)
	var like domain.Like
	err := collection.FindOne(c, bson.M{"user_id": userID, "blog_id": blogID}).Decode(&like)
	if err != nil {
		return nil, err
	}

	return &like, nil
}

// LikeBlog implements domain.LikeRepository.
func (lr *LikeRepository) LikeBlog(c context.Context, like *domain.Like) (*domain.Like, error) {
	collection := lr.database.Collection(lr.repoCollection)

	like.LikeID = primitive.NewObjectID()
	like.CreatedAt = time.Now()

	_, err := collection.InsertOne(c, like)
	if err != nil {
		return nil, err
	}
	return like, nil
}

// UnlikeBlog implements domain.LikeRepository.
func (lr *LikeRepository) UnlikeBlog(c context.Context, likeID string) (*domain.Like, error) {
	collection := lr.database.Collection(lr.repoCollection)

	objectID, err := primitive.ObjectIDFromHex(likeID)
	if err != nil {
		return nil, err
	}

	var like domain.Like
	err = collection.FindOne(c, bson.M{"_id": objectID}).Decode(&like)
	_, err2 := collection.DeleteOne(c, bson.M{"_id": objectID})
	if err != nil {
		return nil, err2
	}

	return &like, nil
}

// DeleteLikeByBlogID implements domain.LikeRepository.
func (lr *LikeRepository) DeleteLikeByBlogID(c context.Context, blogID string) error {

	collection := lr.database.Collection(lr.repoCollection)
	_, err := collection.DeleteMany(c, bson.M{"blog_id": blogID})
	return err
}
