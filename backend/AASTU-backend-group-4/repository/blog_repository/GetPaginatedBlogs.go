package blog_repository

import (
	"blog-api/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (br *BlogRepository) GetPaginatedBlogs(ctx context.Context, page, limit int, sortBy string) ([]domain.Blog, error) {
	var posts []domain.Blog
	skip := int64((page - 1) * limit)
	sort := bson.D{{Key: "created_at", Value: -1}} // default sort by recent
	if sortBy == "popularity" {
		sort = bson.D{{Key: "view_count", Value: -1}}
	}

	findOptions := options.Find().
		SetSkip(skip).
		SetLimit(int64(limit)).
		SetSort(sort)

	cursor, err := br.collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
