package blog_repository

import (
	"blog-api/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (br *BlogRepository) FilterBlog(ctx context.Context, filter domain.FilterRequest) ([]*domain.Blog, error) {
	filters := bson.M{"title": filter.Title, "tags": filter.Tags, "date": filter.Date}

	var blogs []*domain.Blog
	cursor, err := br.collection.Find(ctx, filters, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var b domain.Blog
		if err := cursor.Decode(&b); err != nil {
			return nil, err
		}
		blogs = append(blogs, &b)
	}

	return blogs, nil
}
