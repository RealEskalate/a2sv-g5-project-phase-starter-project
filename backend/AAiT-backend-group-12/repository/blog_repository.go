package repository

import (
	"blog_api/domain"
	"blog_api/domain/dtos"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepository struct {
	// database collection
	collection mongo.Collection
}

var _ domain.BlogRepositoryInterface = &BlogRepository{}

func NewBlogRepository(coll mongo.Collection) *BlogRepository {
	return &BlogRepository{
		collection: coll,
	}
}

// FindBlogPostByID implements domain.BlogRepositoryInterface.
func (b *BlogRepository) FindBlogPostByID(ctx context.Context, blogId string) (*domain.Blog, error) {
	panic("to be implemented by robel")
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
		UserID:    blogDTO.UserID.Hex(),
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
	userID, err := primitive.ObjectIDFromHex(blog.UserID)
	if err != nil {
		return nil, err
	}
	// Similarly, map LikedBy, DislikedBy, and Comments.
	return &dtos.BlogDTO{
		ID:        blogID,
		Title:     blog.Title,
		Content:   blog.Content,
		UserID:    userID,
		Tags:      blog.Tags,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
		ViewCount: blog.ViewCount,
	}, nil
}

// Similar functions for User and Comment...
