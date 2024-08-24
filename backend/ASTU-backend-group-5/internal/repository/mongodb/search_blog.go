package mongodb

import (
	"blogApp/internal/domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *MongoBlogRepository) FindBlogs(
	ctx context.Context,
	filter domain.BlogFilter,
	page int,
	pageSize int,
	orderBy []string, // Accept a list of sorting criteria

) ([]*domain.GetBlogDTO, int, error) {
	collection := r.blogsCollection

	fmt.Println(filter, "page=", page, "pageSize=", pageSize, "orderBy=", orderBy)

	// Building the query
	query := bson.M{}

	// Filter by Author ID
	if filter.AuthorID != nil {
		query["ownerID"] = *filter.AuthorID
	}

	// Filter by Tags (normal tags)
	if filter.Tags != nil && len(filter.Tags) > 0 {
		query["tags.name"] = bson.M{"$in": filter.Tags}
	}

	// Filter by Keyword using $text search
	if filter.Keyword != nil {
		query["$text"] = bson.M{"$search": *filter.Keyword}
	}

	// Additional filters like title, content, date range, etc.
	if filter.Title != nil {
		query["title"] = bson.M{"$regex": *filter.Title, "$options": "i"}
	}

	if filter.Content != nil {
		query["content"] = bson.M{"$regex": *filter.Content, "$options": "i"}
	}

	if filter.DateRange != nil {
		query["created_at"] = bson.M{
			"$gte": filter.DateRange.From,
			"$lte": filter.DateRange.To,
		}
	}

	// Set sorting options based on the list of orderBy parameters
	sortOptions := bson.D{}
	for _, order := range orderBy {
		switch order {
		case "oldest":
			sortOptions = append(sortOptions, bson.E{Key: "created_at", Value: 1})
		case "newest":
			sortOptions = append(sortOptions, bson.E{Key: "created_at", Value: -1})
		case "most_likes":
			sortOptions = append(sortOptions, bson.E{Key: "likes_count", Value: -1})
		case "most_views":
			sortOptions = append(sortOptions, bson.E{Key: "views_count", Value: -1})
		case "most_comments":
			sortOptions = append(sortOptions, bson.E{Key: "comments_count", Value: -1})
		default:
			// Handle unknown sorting parameters (you might want to ignore or log this)
		}
	}

	opr := options.Find().
		SetSkip(int64((page - 1) * pageSize)).
		SetLimit(int64(pageSize)).
		SetSort(sortOptions)

	// Execute the query
	cursor, err := collection.Find(ctx, query, opr)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var blogs []*domain.GetBlogDTO
	if err := cursor.All(ctx, &blogs); err != nil {
		return nil, 0, err
	}

	// Get the total count of blogs matching the query
	count, err := collection.CountDocuments(ctx, query)
	if err != nil {
		return nil, 0, err
	}

	return blogs, int(count), nil
}

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
