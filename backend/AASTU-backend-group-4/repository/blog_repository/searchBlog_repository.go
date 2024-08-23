package blog_repository

import (
	"blog-api/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (br *BlogRepository) SearchBlog(ctx context.Context, filter map[string]string) ([]*domain.Blog, error) {

	filters := bson.M{}

	if filter["title"] != "" {
		filters["title"] = bson.M{"$regex": filter["title"], "$options": "i"}
	}

	if filter["author"] != "" {
		filters["author"] = bson.M{"$regex": filter["author"], "$options": "i"}
	}

	var blogs []*domain.Blog
	cursor, err := br.collection.Find(ctx, filters)
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
