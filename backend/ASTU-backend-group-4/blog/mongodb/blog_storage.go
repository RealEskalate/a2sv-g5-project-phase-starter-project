package mongodb

import (
	"context"
	"log"

	blogDomain "github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/blog"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	blogCollection    = "blogs"
	commentCollection = "comments"
	likeCollection    = "likes"
	dislikeCollection = "dislikes"
)

type BlogStorage struct {
	db *mongo.Database
}

func NewBlogStorage(db *mongo.Database) *BlogStorage {
	return &BlogStorage{db: db}
}

// CreateBlog implements BlogRepository.
func (b *BlogStorage) CreateBlog(ctx context.Context, blog blogDomain.Blog) (string, error) {
	result, err := b.db.Collection(blogCollection).InsertOne(ctx, blog)
	if err != nil {
		log.Default().Printf("Failed to create blog: %v", err)
		return "", blogDomain.ErrUnableToCreateBlog
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

// CreateComment implements BlogRepository.
func (b *BlogStorage) CreateComment(ctx context.Context, comment blogDomain.Comment) (string, error) {
	result, err := b.db.Collection(commentCollection).InsertOne(ctx, comment)
	if err != nil {
		log.Default().Printf("Failed to create comment: %v", err)
		return "", blogDomain.ErrUnableToCreatComment
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

// DeleteBlog implements BlogRepository.
func (b *BlogStorage) DeleteBlog(ctx context.Context, id string) error {
	blogIDPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return blogDomain.ErrInvalidID
	}

	filter := bson.D{{Key: "_id", Value: blogIDPrimitive}}

	result, err := b.db.Collection(blogCollection).DeleteOne(ctx, filter)

	if err != nil {
		log.Default().Printf("Failed to delete blog: %v", err)
		return blogDomain.ErrUnableToDeleteBlog
	}

	if result.DeletedCount == 0 {
		return blogDomain.ErrBlogNotFound
	}

	return nil
}

// DeleteComment implements BlogRepository.
func (b *BlogStorage) DeleteComment(ctx context.Context, id string) error {
	commentIDPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return blogDomain.ErrInvalidID
	}

	filter := bson.D{{Key: "_id", Value: commentIDPrimitive}}

	result, err := b.db.Collection(commentCollection).DeleteOne(ctx, filter)

	if err != nil {
		log.Default().Printf("Failed to delete comment: %v", err)
		return blogDomain.ErrUnableToDeleteBlog
	}

	if result.DeletedCount == 0 {
		return blogDomain.ErrBlogNotFound
	}

	return nil
}

// DislikeBlog implements BlogRepository.
func (b *BlogStorage) DislikeBlog(ctx context.Context, dislike blogDomain.Dislike) error {
	_, err := b.db.Collection(dislikeCollection).InsertOne(ctx, dislike)
	if err != nil {
		log.Default().Printf("Failed to dislike blog: %v", err)
		return blogDomain.ErrUnableToDislikeBlog
	}

	return nil
}

// GetBlogByID implements BlogRepository.
func (b *BlogStorage) GetBlogByID(ctx context.Context, id string) (blogDomain.Blog, error) {
	blogIDPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return blogDomain.Blog{}, blogDomain.ErrInvalidID
	}

	filter := bson.D{{Key: "_id", Value: blogIDPrimitive}}

	var blogData blogDomain.Blog
	err = b.db.Collection(blogCollection).FindOne(ctx, filter).Decode(&blogData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return blogDomain.Blog{}, blogDomain.ErrBlogNotFound
		}
		log.Default().Printf("Failed to get blog by ID: %v", err)
		return blogDomain.Blog{}, blogDomain.ErrUnabletoGetBlog
	}

	return blogData, nil
}

// GetBlogs implements BlogRepository.
func (b *BlogStorage) GetBlogs(ctx context.Context, filterQuery blogDomain.FilterQuery, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[blogDomain.Blog], error) {
	filter := bson.D{}

	if filterQuery.Tags != nil {
		filter = append(filter, bson.E{Key: "tags", Value: bson.D{{Key: "$in", Value: filterQuery.Tags}}})
	}
	if filterQuery.CreatedAtFrom != "" && filterQuery.CreatedAtTo != "" {
		filter = append(filter, bson.E{Key: "created_at", Value: bson.D{
			{Key: "$gte", Value: filterQuery.CreatedAtFrom},
			{Key: "$lte", Value: filterQuery.CreatedAtTo},
		}})
	}

	if filterQuery.Popularity != 0 {
		if filterQuery.Popularity > 0 {
			filter = append(filter, bson.E{Key: "$sort", Value: bson.D{{Key: "popularity", Value: 1}}})
		} else {
			filter = append(filter, bson.E{Key: "$sort", Value: bson.D{{Key: "popularity", Value: -1}}})
		}
	}

	findOptions := options.Find()
	findOptions.SetSkip(int64(pagination.Limit*pagination.Page - 1))
	findOptions.SetLimit(int64(pagination.Limit))
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	count, err := b.db.Collection(blogCollection).CountDocuments(ctx, filter)
	if err != nil {
		return infrastructure.PaginationResponse[blogDomain.Blog]{}, err
	}

	cursor, err := b.db.Collection(blogCollection).Find(ctx, filter, findOptions)
	if err != nil {
		log.Default().Printf("Failed to get blogs: %v", err)
		return infrastructure.PaginationResponse[blogDomain.Blog]{}, blogDomain.ErrUnabletoGetBlogs
	}

	var blogs []blogDomain.Blog
	cursor.All(ctx, &blogs)

	return infrastructure.NewPaginationResponse[blogDomain.Blog](pagination.Limit, pagination.Page, count, blogs), nil
}

// GetCommentsByBlogID implements BlogRepository.
func (b *BlogStorage) GetCommentsByBlogID(ctx context.Context, blogID string, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[blogDomain.Comment], error) {
	blogIDPrimitive, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return infrastructure.PaginationResponse[blogDomain.Comment]{}, blogDomain.ErrInvalidID
	}

	filter := bson.D{{Key: "blog_id", Value: blogIDPrimitive}}

	findOptions := options.Find()
	findOptions.SetSkip(int64(pagination.Limit*pagination.Page - 1))
	findOptions.SetLimit(int64(pagination.Limit))
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	count, err := b.db.Collection(commentCollection).CountDocuments(ctx, filter)
	if err != nil {
		return infrastructure.PaginationResponse[blogDomain.Comment]{}, err
	}

	cursor, err := b.db.Collection(commentCollection).Find(ctx, filter, findOptions)
	if err != nil {
		log.Default().Printf("Failed to get comments by blog ID: %v", err)
		return infrastructure.PaginationResponse[blogDomain.Comment]{}, blogDomain.ErrUnableToGetComments
	}

	var comments []blogDomain.Comment
	cursor.All(ctx, &comments)

	return infrastructure.NewPaginationResponse[blogDomain.Comment](pagination.Limit, pagination.Page, count, comments), nil
}

// LikeBlog implements BlogRepository.
func (b *BlogStorage) LikeBlog(ctx context.Context, like blogDomain.Like) error {
	_, err := b.db.Collection(likeCollection).InsertOne(ctx, like)
	if err != nil {
		log.Default().Printf("Failed to like blog: %v", err)
		return blogDomain.ErrUnableToLikeBlog
	}

	return nil
}

// SearchBlogs implements BlogRepository.
func (b *BlogStorage) SearchBlogs(ctx context.Context, query string, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[blogDomain.Blog], error) {
	filter := bson.D{{Key: "$text", Value: bson.D{
		{Key: "$search", Value: query},
		{Key: "$caseSensitive", Value: false},
	}}}

	findOptions := options.Find()
	findOptions.SetSkip(int64(pagination.Limit*pagination.Page - 1))
	findOptions.SetLimit(int64(pagination.Limit))
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	count, err := b.db.Collection(blogCollection).CountDocuments(ctx, filter)
	if err != nil {
		return infrastructure.PaginationResponse[blogDomain.Blog]{}, err
	}

	cursor, err := b.db.Collection(blogCollection).Find(ctx, filter, findOptions)
	if err != nil {
		log.Default().Printf("Failed to search blogs: %v", err)
		return infrastructure.PaginationResponse[blogDomain.Blog]{}, blogDomain.ErrUnabletoSearchBlogs
	}

	var blogs []blogDomain.Blog
	cursor.All(ctx, &blogs)

	return infrastructure.NewPaginationResponse[blogDomain.Blog](pagination.Limit, pagination.Page, count, blogs), nil
}

// UpdateBlog implements BlogRepository.
func (b *BlogStorage) UpdateBlog(ctx context.Context, id string, blog blogDomain.Blog) error {
	blogIDPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return blogDomain.ErrInvalidID
	}

	filter := bson.D{{Key: "_id", Value: blogIDPrimitive}}

	result, err := b.db.Collection(blogCollection).ReplaceOne(ctx, filter, blog)
	if err != nil {
		log.Default().Printf("Failed to update blog: %v", err)
		return blogDomain.ErrUnableToUpdateBlog
	}

	if result.ModifiedCount == 0 {
		return blogDomain.ErrBlogNotFound
	}

	return nil
}

// UnlikeBlog implements BlogRepository.
func (b *BlogStorage) UnlikeBlog(ctx context.Context, like blogDomain.Like) error {
	_, err := b.db.Collection(likeCollection).DeleteOne(ctx, like)
	if err != nil {
		log.Default().Printf("Failed to unlike blog: %v", err)
		return blogDomain.ErrUnableToUnLikeBlog
	}

	return nil
}

// UndislikeBlog implements BlogRepository.
func (b *BlogStorage) UndislikeBlog(ctx context.Context, dislike blogDomain.Dislike) error {
	_, err := b.db.Collection(dislikeCollection).DeleteOne(ctx, dislike)
	if err != nil {
		log.Default().Printf("Failed to unlike blog: %v", err)
		return blogDomain.ErrUnableToUnDislikeBlog
	}

	return nil
}
