package repository

import (
	"blog_api/domain"
	"blog_api/domain/dtos"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogRepository struct {
	// database collection
	collection *mongo.Collection
}

func NewBlogRepository(coll *mongo.Collection) *BlogRepository {
	return &BlogRepository{
		collection: coll,
	}
}

// FindBlogPostByID implements domain.BlogRepositoryInterface.
func (b *BlogRepository) FetchBlogPostByID(ctx context.Context, postID string) (*domain.Blog, error) {
	filter := bson.D{{Key: "_id", Value: postID}}
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "view_count", Value: 1}}}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var post domain.Blog
	err := b.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// fetches blogs based on filter.The filtering options are defined in the domain named BlogFilterOptions.
func (b *BlogRepository) GetBlogPosts(ctx context.Context, filters domain.BlogFilterOptions) ([]domain.Blog, int, error) {
	var query bson.D

	// Search by title
	if filters.Title != "" {
		query = append(query, bson.E{Key: "title", Value: bson.D{{Key: "$regex", Value: filters.Title}, {Key: "$options", Value: "i"}}}) // Case-insensitive
	}

	// Search by author name
	if filters.Author != "" {
		query = append(query, bson.E{Key: "username", Value: bson.D{{Key: "$regex", Value: filters.Author}, {Key: "$options", Value: "i"}}}) // Case-insensitive
	}

	// Filter by tag
	if len(filters.Tags) > 0 {
		query = append(query, bson.E{Key: "tags", Value: bson.D{{Key: "$in", Value: filters.Tags}}})
	}

	// Filter by date range
	if !filters.DateFrom.IsZero() && !filters.DateTo.IsZero() {
		query = append(query, bson.E{Key: "created_at", Value: bson.D{{Key: "$gte", Value: filters.DateFrom}, {Key: "$lte", Value: filters.DateTo}}})
	} else if !filters.DateFrom.IsZero() {
		query = append(query, bson.E{Key: "created_at", Value: bson.D{{Key: "$gte", Value: filters.DateFrom}}})
	} else if !filters.DateTo.IsZero() {
		query = append(query, bson.E{Key: "created_at", Value: bson.D{{Key: "$lte", Value: filters.DateTo}}})
	}

	// Filter by popularity metrics
	if filters.MinLikes > 0 {
		query = append(query, bson.E{Key: "liked_by", Value: bson.D{{Key: "$size", Value: bson.D{{Key: "$gte", Value: filters.MinLikes}}}}})
	}
	if filters.MinDislikes > 0 {
		query = append(query, bson.E{Key: "disliked_by", Value: bson.D{{Key: "$size", Value: bson.D{{Key: "$gte", Value: filters.MinDislikes}}}}})
	}
	if filters.MinComments > 0 {
		query = append(query, bson.E{Key: "comments", Value: bson.D{{Key: "$size", Value: bson.D{{Key: "$gte", Value: filters.MinComments}}}}})
	}
	if filters.MinViewCount > 0 {
		query = append(query, bson.E{Key: "view_count", Value: bson.D{{Key: "$gte", Value: filters.MinViewCount}}})
	}

	findOptions := options.Find()

	// Sorting
	sort := bson.D{}
	if filters.SortBy != "" {
		sortDirection := 1 // Default ascending
		if filters.SortDirection == "desc" {
			sortDirection = -1
		}
		sort = append(sort, bson.E{Key: filters.SortBy, Value: sortDirection})
		findOptions.SetSort(sort)
	}

	// Pagination
	if filters.Page > 0 && filters.PostsPerPage > 0 {
		skip := (filters.Page - 1) * filters.PostsPerPage
		findOptions.SetSkip(int64(skip))
		findOptions.SetLimit(int64(filters.PostsPerPage))
	}

	// Execute the query
	cursor, err := b.collection.Find(ctx, query, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var blogDTOs []dtos.BlogDTO
	if err := cursor.All(ctx, &blogDTOs); err != nil {
		return nil, 0, err
	}

	// Convert DTOs to domain models
	var blogs []domain.Blog
	for _, blogDTO := range blogDTOs {
		blogs = append(blogs, *toDomain(&blogDTO))
	}

	// Get total count for pagination metadata
	totalCount, err := b.collection.CountDocuments(ctx, query)
	if err != nil {
		return nil, 0, err
	}

	return blogs, int(totalCount), nil
}


// DeleteBlogPost implements domain.BlogRepositoryInterface.
func (b *BlogRepository) DeleteBlogPost(ctx context.Context, blogId string) error {
	objID, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	_, err = b.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

// InsertBlogPost implements domain.BlogRepositoryInterface.
func (b *BlogRepository) InsertBlogPost(ctx context.Context, blog *domain.Blog) error {

	newBlog, err := toDTO(blog)
	newBlog.ID = primitive.NewObjectID()
	if err != nil {
		return err
	}
	_, err = b.collection.InsertOne(ctx, newBlog)
	if err != nil {
		return err
	}

	return nil
}

// UpdateBlogPost implements domain.BlogRepositoryInterface.
func (b *BlogRepository) UpdateBlogPost(ctx context.Context, blogId string, blog *domain.Blog) error {
	objID, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	update := bson.M{
		"$set": bson.M{
			"title":     blog.Title,
			"content":   blog.Content,
			"updatedAt": blog.UpdatedAt,
		},
	}

	_, err = b.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

// Converts BlogDTO to Blog domain model.
func toDomain(blogDTO *dtos.BlogDTO) *domain.Blog {
	return &domain.Blog{
		ID:        blogDTO.ID.Hex(),
		Title:     blogDTO.Title,
		Content:   blogDTO.Content,
		Username:  blogDTO.Username,
		Tags:      blogDTO.Tags,
		CreatedAt: blogDTO.CreatedAt,
		UpdatedAt: blogDTO.UpdatedAt,
		ViewCount: blogDTO.ViewCount,
		// Map LikedBy, DislikedBy, and Comments appropriately
	}
}

// Converts Blog domain model to BlogDTO.
func toDTO(blog *domain.Blog) (*dtos.BlogDTO, error) {
	blogID, err := primitive.ObjectIDFromHex(blog.ID)
	if err != nil {
		return nil, err
	}

	// Similarly, map LikedBy, DislikedBy, and Comments.
	return &dtos.BlogDTO{
		ID:        blogID,
		Title:     blog.Title,
		Content:   blog.Content,
		Username:  blog.Username,
		Tags:      blog.Tags,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
		ViewCount: blog.ViewCount,
	}, nil
}

// Similar functions for User and Comment...
