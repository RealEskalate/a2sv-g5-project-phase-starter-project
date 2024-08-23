package mongodb

import (
	"blogApp/internal/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// FilterBlogs filters blogs based on the provided criteria
func (r *MongoBlogRepository) FilterBlogs(ctx context.Context, filter domain.BlogFilter) ([]*domain.GetBlogDTO, error) {
	query := bson.M{}

	// Apply filters based on the provided filter object
	if filter.Title != nil && *filter.Title != "" {
		query["title"] = *filter.Title
	}
	if filter.AuthorID != nil {
		query["ownerID"] = *filter.AuthorID
	}
	if len(filter.Tags) > 0 {
		query["tags.name"] = bson.M{"$in": filter.Tags}
	}
	if filter.DateRange != nil {
		query["created_at"] = bson.M{
			"$gte": filter.DateRange.From,
			"$lte": filter.DateRange.To,
		}
	}
	if filter.Content != nil && *filter.Content != "" {
		query["content"] = bson.M{"$regex": *filter.Content, "$options": "i"}
	}
	if filter.Keyword != nil && *filter.Keyword != "" {
		keyword := *filter.Keyword
		query["$or"] = []bson.M{
			{"title": bson.M{"$regex": keyword, "$options": "i"}},
			{"content": bson.M{"$regex": keyword, "$options": "i"}},
			{"tags.name": bson.M{"$regex": keyword, "$options": "i"}},
		}
	}

	// Retrieve blogs matching the filter
	cursor, err := r.blogsCollection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var blogs []*domain.GetBlogDTO
	for cursor.Next(ctx) {
		var blog domain.GetBlogDTO
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}
