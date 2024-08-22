package repository

import (
	"blog_api/domain"
	"blog_api/domain/dtos"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogRepository struct {
	// database collection
	collection *mongo.Collection
}

var _ domain.BlogRepositoryInterface = &BlogRepository{}

func NewBlogRepository(coll *mongo.Collection) *BlogRepository {
	return &BlogRepository{
		collection: coll,
	}
}

// FetchBlogPostByID retrieves a blog post by its ID and increments the view count.
func (b *BlogRepository) FetchBlogPostByID(ctx context.Context, blogId string, incrementView bool) (*domain.Blog, domain.CodedError) {
	objID, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return nil, domain.NewError("Invalid blog ID", domain.ERR_BAD_REQUEST)
	}

	filter := bson.D{{Key: "_id", Value: objID}}

	var post dtos.BlogDTO
	if !incrementView {
		err := b.collection.FindOne(ctx, filter).Decode(&post)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return nil, domain.NewError("Blog post not found", domain.ERR_NOT_FOUND)
			}
			return nil, domain.NewError("Internal server error: "+err.Error(), domain.ERR_INTERNAL_SERVER)
		}
	} else {
		update := bson.D{{Key: "$inc", Value: bson.D{{Key: "view_count", Value: 1}}}}
		opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	
		err = b.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&post)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return nil, domain.NewError("Blog post not found", domain.ERR_NOT_FOUND)
			}
			return nil, domain.NewError("Internal server error: "+err.Error(), domain.ERR_INTERNAL_SERVER)
		}
	}

	convertedPost := *toDomain(&post)
	return &convertedPost, nil
}

// fetches blogs based on filter.The filtering options are defined in the domain named BlogFilterOptions.
func (b *BlogRepository) FetchBlogPosts(ctx context.Context, filters domain.BlogFilterOptions) ([]domain.Blog, int, domain.CodedError) {
	query := bson.D{}

	// Search by title
	if filters.Title != "" {
		query = append(query, bson.E{Key: "title", Value: bson.D{{Key: "$regex", Value: filters.Title}, {Key: "$options", Value: "i"}}}) // Case-insensitive
	}

	// Search by author name
	if filters.Author != "" {
		query = append(query, bson.E{Key: "username", Value: bson.D{{Key: "$regex", Value: filters.Author}, {Key: "$options", Value: "i"}}}) // Case-insensitive
	}

	// Filter by tags
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

	// Initial match stage for filtering
	aggregation := mongo.Pipeline{
		{{Key: "$match", Value: query}},
	}

	// Add computed fields for like_count and dislike_count
	computeStage := bson.D{{Key: "$addFields", Value: bson.D{
		{Key: "like_count", Value: bson.D{{Key: "$size", Value: "$liked_by"}}},
		{Key: "dislike_count", Value: bson.D{{Key: "$size", Value: "$disliked_by"}}},
	}}}
	aggregation = append(aggregation, computeStage)

	// Sorting stage
	if filters.SortBy != "" {
		sortDirection := 1 // Default ascending
		if filters.SortDirection == "desc" {
			sortDirection = -1
		}
		aggregation = append(aggregation, bson.D{{Key: "$sort", Value: bson.D{{Key: filters.SortBy, Value: sortDirection}}}})
	}

	// Pagination
	if filters.Page > 0 && filters.PostsPerPage > 0 {
		skip := (filters.Page - 1) * filters.PostsPerPage
		aggregation = append(aggregation, bson.D{{Key: "$skip", Value: skip}})
		aggregation = append(aggregation, bson.D{{Key: "$limit", Value: filters.PostsPerPage}})
	}

	// Execute the aggregation pipeline
	cursor, err := b.collection.Aggregate(ctx, aggregation)
	if err != nil {
		return nil, 0, domain.NewError("Internal server error: "+err.Error(), domain.ERR_INTERNAL_SERVER)
	}
	defer cursor.Close(ctx)

	var blogDTOs []dtos.BlogDTO
	if err := cursor.All(ctx, &blogDTOs); err != nil {
		return nil, 0, domain.NewError("Internal server error: "+err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	// Convert DTOs to domain models
	var blogs []domain.Blog
	for _, blogDTO := range blogDTOs {
		blogs = append(blogs, *toDomain(&blogDTO))
	}

	return blogs, len(blogs), nil
}

// DeleteBlogPost deletes a blog post by its ID.
func (b *BlogRepository) DeleteBlogPost(ctx context.Context, blogId string) domain.CodedError {
	objID, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return domain.NewError("Invalid blog ID", domain.ERR_BAD_REQUEST)
	}

	filter := bson.M{"_id": objID}

	_, err = b.collection.DeleteOne(ctx, filter)

	if err != nil {
		return domain.NewError("Internal server error: "+err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

// InsertBlogPost inserts a new blog post into the database.
func (b *BlogRepository) InsertBlogPost(ctx context.Context, blog *domain.Blog) domain.CodedError {

	newBlog, err := toDTO(blog)
	if err != nil {
		return domain.NewError("Internal server error: "+err.Error(), domain.ERR_INTERNAL_SERVER)
	}
	newBlog.ID = primitive.NewObjectID()

	_, err = b.collection.InsertOne(ctx, newBlog)
	if err != nil {
		return domain.NewError("Internal server error: "+err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

// UpdateBlogPost implements domain.BlogRepositoryInterface.
func (b *BlogRepository) UpdateBlogPost(ctx context.Context, blogId string, blog *domain.NewBlog) domain.CodedError {
	objID, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return domain.NewError("Invalid blog ID", domain.ERR_BAD_REQUEST)
	}

	filter := bson.M{"_id": objID}

	update := bson.M{
		"$set": bson.M{
			"title":      blog.Title,
			"content":    blog.Content,
			"tags":       blog.Tags,
			"updated_at": time.Now(),
		},
	}

	_, err = b.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.NewError("Internal server error: "+err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

func (b *BlogRepository) TrackBlogPopularity(ctx context.Context, blogId string, action string, state bool, username string) domain.CodedError {
	objID, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return domain.NewError("Invalid blog ID", domain.ERR_BAD_REQUEST)
	}

	var update bson.D

	switch action {
	case "like":
		if !state {
			update = bson.D{
				{Key: "$pull", Value: bson.D{{Key: "liked_by", Value: username}}},
			}
		} else {
			update = bson.D{
				{Key: "$addToSet", Value: bson.D{{Key: "liked_by", Value: username}}},
				{Key: "$pull", Value: bson.D{{Key: "disliked_by", Value: username}}},
			}
		}
	case "dislike":
		if !state {
			update = bson.D{
				{Key: "$pull", Value: bson.D{{Key: "disliked_by", Value: username}}},
			}
		} else {
			update = bson.D{
				{Key: "$addToSet", Value: bson.D{{Key: "disliked_by", Value: username}}},
				{Key: "$pull", Value: bson.D{{Key: "liked_by", Value: username}}},
			}
		}
	default:
		return domain.NewError("Invalid action", domain.ERR_BAD_REQUEST)
	}

	_, err = b.collection.UpdateOne(ctx, bson.D{{Key: "_id", Value: objID}}, update)
	if err != nil {
		return domain.NewError("Failed to update blog popularity", domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

// Converts BlogDTO to Blog domain model.
func toDomain(blogDTO *dtos.BlogDTO) *domain.Blog {
	return &domain.Blog{
		ID:         blogDTO.ID.Hex(),
		Title:      blogDTO.Title,
		Content:    blogDTO.Content,
		Username:   blogDTO.Username,
		Tags:       blogDTO.Tags,
		CreatedAt:  blogDTO.CreatedAt,
		UpdatedAt:  blogDTO.UpdatedAt,
		ViewCount:  blogDTO.ViewCount,
		LikedBy:    blogDTO.LikedBy,
		DislikedBy: blogDTO.DislikedBy,
	}
}

// Converts Blog domain model to BlogDTO.
func toDTO(blog *domain.Blog) (*dtos.BlogDTO, error) {
	return &dtos.BlogDTO{
		Title:      blog.Title,
		Content:    blog.Content,
		Username:   blog.Username,
		Tags:       blog.Tags,
		CreatedAt:  blog.CreatedAt,
		UpdatedAt:  blog.UpdatedAt,
		ViewCount:  blog.ViewCount,
		LikedBy:    blog.LikedBy,
		DislikedBy: blog.DislikedBy,
	}, nil
}

// CreateComment implements domain.BlogRepositoryInterface.
func (b *BlogRepository) CreateComment(ctx context.Context, comment *domain.Comment, blogId string, createdBy string) domain.CodedError {
	// convert blogId to objectID
	blogID, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return domain.NewError("Invalid blog ID", domain.ERR_BAD_REQUEST)
	}

	//filter the blog using the ID
	filter := bson.M{"_id": blogID}
	var foundBlog dtos.BlogDTO
	err = b.collection.FindOne(ctx, filter).Decode(&foundBlog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.NewError("Blog not found", domain.ERR_NOT_FOUND)
		}
		return domain.NewError("Internal Server Error", domain.ERR_INTERNAL_SERVER)
	}

	// create a new commentDTO
	newComment := dtos.CommentDTO{
		ID:        primitive.NewObjectID(),
		Content:   comment.Content,
		Username:  createdBy,
		CreatedAt: time.Now(),
	}

	//// Append the new comment and update the blog
	foundBlog.Comments = append(foundBlog.Comments, newComment)
	update := bson.M{"$set": bson.M{"comments": foundBlog.Comments}}

	result, err := b.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.NewError("Internal Server Error"+err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	// Ensure the update affected one document
	if result.ModifiedCount == 0 {
		return domain.NewError("No blog was updated", domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

// DeleteComment implements domain.BlogRepositoryInterface.
func (b *BlogRepository) DeleteComment(ctx context.Context, commentId, blogId, userName string) domain.CodedError {
	blogID, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return domain.NewError("Invalid blog ID", domain.ERR_BAD_REQUEST)
	}

	//filter the blog using the ID
	filter := bson.M{"_id": blogID}
	var foundBlog dtos.BlogDTO
	err = b.collection.FindOne(ctx, filter).Decode(&foundBlog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.NewError("Blog not found"+err.Error(), domain.ERR_NOT_FOUND)
		}
		return domain.NewError("Internal Server Error"+err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	// Find and delete the comment if it exists and the user is the owner
	var updatedComments []dtos.CommentDTO
	commentDeleted := false
	for _, comment := range foundBlog.Comments {
		if comment.ID.Hex() == commentId {
			if comment.Username != userName {
				return domain.NewError("Unauthorized", domain.ERR_FORBIDDEN)
			}
			commentDeleted = true
			continue //jump the current comment
		}
		updatedComments = append(updatedComments, comment)
	}

	if !commentDeleted {
		return domain.NewError("Comment not found", domain.ERR_NOT_FOUND)
	}

	update := bson.M{"$set": bson.M{"comments": updatedComments}}
	result, err := b.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.NewError("Internal Server Error"+err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	// Ensure the update affected one document
	if result.ModifiedCount == 0 {
		return domain.NewError("No blog was updated", domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

// UpdateComment implements domain.BlogRepositoryInterface.
func (b *BlogRepository) UpdateComment(ctx context.Context, updateComment *domain.NewComment, commentId, blogId, userName string) domain.CodedError {
	blogID, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return domain.NewError("Invalid blog ID", domain.ERR_BAD_REQUEST)
	}

	//filter the blog using the ID
	filter := bson.M{"_id": blogID}
	var foundBlog dtos.BlogDTO
	err = b.collection.FindOne(ctx, filter).Decode(&foundBlog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.NewError("Blog not found"+err.Error(), domain.ERR_NOT_FOUND)
		}
		return domain.NewError("Internal Server Error"+err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	// Find and update the comment if it exists and the user is the owner
	commentUpdated := false
	for i := range foundBlog.Comments {
		if foundBlog.Comments[i].ID.Hex() == commentId {
			if foundBlog.Comments[i].Username != userName {
				return domain.NewError("Unauthorized request for update", domain.ERR_FORBIDDEN)
			}
			// Update the comment
			foundBlog.Comments[i].Content = updateComment.Content
			foundBlog.Comments[i].UpdatedAt = time.Now()
			commentUpdated = true
			break
		}
	}

	if !commentUpdated {
		return domain.NewError("Comment not found", domain.ERR_NOT_FOUND)
	}

	update := bson.M{"$set": bson.M{"comments": foundBlog.Comments}}
	result, err := b.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.NewError("Internal Server Error"+err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	// Ensure the update affected one document
	if result.ModifiedCount == 0 {
		return domain.NewError("No blog was updated", domain.ERR_INTERNAL_SERVER)
	}

	return nil
}
