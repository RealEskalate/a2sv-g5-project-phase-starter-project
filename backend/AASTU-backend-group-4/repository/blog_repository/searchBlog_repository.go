package blog_repository

import (
	domain "blog-api/domain/blog"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (br *BlogRepository) SearchBlog(ctx context.Context, title string, author string) ([]*domain.Blog, error) {

	filter := bson.M{}

	if title != "" {
		filter["title"] = bson.M{"$regex": title, "$options": "i"}
	}
	if author != "" {
		filter["author"] = bson.M{"$regex": author, "$options": "i"}
	}

	var blogs []*domain.Blog
	cursor, err := br.collection.Find(ctx, filter, options.Find())
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
