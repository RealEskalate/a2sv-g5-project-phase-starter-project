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

func (bpr *BlogPopularityRepository) GetPopularBlogs(sortBys []domain.SortBy, sortOrders []domain.SortOrder) ([]domain.Blog, error) {
	var popularBlogs []domain.Blog

	if len(sortBys) == 0 || len(sortOrders) == 0 || len(sortBys) != len(sortOrders) {
		return nil, domain.ErrInvalidSortBy
	}

	opts := options.Find()
	sortFields := make([]bson.E, len(sortBys))

	for i, sortBy := range sortBys {
		sortKey := mapSortByToKey(sortBy)
		if sortKey == "" {
			return nil, domain.ErrInvalidSortBy
		}
		sortFields[i] = bson.E{Key: sortKey, Value: int(sortOrders[i])}
	}

	opts.SetSort(bson.D(sortFields))

	cursor, err := bpr.collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &popularBlogs); err != nil {
		return nil, err
	}

	return popularBlogs, nil
}

func mapSortByToKey(sortBy domain.SortBy) string {
	switch sortBy {
	case domain.SortByLikeCount:
		return "likes"
	case domain.SortByCommentCount:
		return "comments"
	case domain.SortByPublishDate:
		return "created_at"
	case domain.SortByEngagement:
		return "engagement"
	case domain.SortByDislikeCount:
		return "dislikes"
	default:
		return ""
	}
}
