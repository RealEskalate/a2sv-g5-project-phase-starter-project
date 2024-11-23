package blog_repository

import (
	"blog-api/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *BlogRepository) SearchBlogs(ctx context.Context, filters domain.BlogFilters) ([]domain.Blog, error) {
	var blogs []domain.Blog

	filter := bson.M{}
	if filters.Author != "" {
		filter["author"] = filters.Author
	}
	if filters.Title != "" {
		filter["title"] = bson.M{"$regex": filters.Title, "$options": "i"}
	}
	if len(filters.Tags) > 0 {
		filter["tags"] = bson.M{"$in": filters.Tags}
	}
	if filters.DateFrom != nil || filters.DateTo != nil {
		dateFilter := bson.M{}
		if filters.DateFrom != nil {
			dateFilter["$gte"] = *filters.DateFrom
		}
		if filters.DateTo != nil {
			dateFilter["$lte"] = *filters.DateTo
		}
		filter["date"] = dateFilter
	}

	cursor, err := r.collection.Find(ctx, filter, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}
