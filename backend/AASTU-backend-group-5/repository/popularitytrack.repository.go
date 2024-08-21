package repository

import (
	"context"

	"github.com/RealEskalate/blogpost/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogPopularityRepository struct {
	collection *mongo.Collection
}

func NewBlogPopularityRepository(coll *mongo.Collection) domain.BlogPopularityRepository {
	return &BlogPopularityRepository{collection: coll}
}

func (bpr *BlogPopularityRepository) GetPopularBlogs(sortBy string, sortOrder int) ([]domain.Blog, error) {
	var popularBlogs []domain.Blog
	opts := options.Find().SetSort(bson.D{{Key: sortBy, Value: sortOrder}})
	cursor, err := bpr.collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &popularBlogs); err != nil {
		return nil, err
	}
	return popularBlogs, nil
}
