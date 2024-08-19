package mongodb

import (
	"context"
	"errors"
	"log"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/blog"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	blogCollection    = "blogs"
	commentCollection = "comments"
	likeCollection    = "likes"
	dislikeCollection = "dislikes"
)

var ErrUnableToCreateBlog = errors.New("unable to create blog")
var ErrUnableToCreatComment = errors.New("unable to create comment")
var ErrInvalidID = errors.New("invalid ID")
var ErrUnableToDeleteBlog = errors.New("unable to delete blog")
var ErrBlogNotFound = errors.New("blog not found")
var ErrUnableToDeleteComment = errors.New("unable to delete comment")
var ErrCommentNotFound = errors.New("comment not found")
var ErrUnableToDislikeBlog = errors.New("unable to dislike blog")
var ErrUnableToLikeBlog = errors.New("unable to like blog")
var ErrUnableToUnLikeBlog = errors.New("unable to unlike blog")
var ErrUnableToUnDislikeBlog = errors.New("unable to unlike blog")
var ErrUnabletoGetBlog = errors.New("unable to get blog")

type BlogStorage struct {
	db *mongo.Database
}

func NewBlogStorage(db *mongo.Database) *BlogStorage {
	return &BlogStorage{db: db}
}

// CreateBlog implements BlogRepository.
func (b *BlogStorage) CreateBlog(ctx context.Context, blog blog.Blog) (string, error) {
	result, err := b.db.Collection(blogCollection).InsertOne(ctx, blog)
	if err != nil {
		log.Default().Printf("Failed to create blog: %v", err)
		return "", ErrUnableToCreateBlog
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

// CreateComment implements BlogRepository.
func (b *BlogStorage) CreateComment(ctx context.Context, comment blog.Comment) (string, error) {
	result, err := b.db.Collection(commentCollection).InsertOne(ctx, comment)
	if err != nil {
		log.Default().Printf("Failed to create comment: %v", err)
		return "", ErrUnableToCreatComment
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

// DeleteBlog implements BlogRepository.
func (b *BlogStorage) DeleteBlog(ctx context.Context, id string) error {
	blogIDPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID
	}

	filter := bson.D{{Key: "_id", Value: blogIDPrimitive}}

	result, err := b.db.Collection(blogCollection).DeleteOne(ctx, filter)

	if err != nil {
		log.Default().Printf("Failed to delete blog: %v", err)
		return ErrUnableToDeleteBlog
	}

	if result.DeletedCount == 0 {
		return ErrBlogNotFound
	}

	return nil
}

// DeleteComment implements BlogRepository.
func (b *BlogStorage) DeleteComment(ctx context.Context, id string) error {
	commentIDPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID
	}

	filter := bson.D{{Key: "_id", Value: commentIDPrimitive}}

	result, err := b.db.Collection(commentCollection).DeleteOne(ctx, filter)

	if err != nil {
		log.Default().Printf("Failed to delete comment: %v", err)
		return ErrUnableToDeleteBlog
	}

	if result.DeletedCount == 0 {
		return ErrBlogNotFound
	}

	return nil
}

// DislikeBlog implements BlogRepository.
func (b *BlogStorage) DislikeBlog(ctx context.Context, dislike blog.Dislike) error {
	_, err := b.db.Collection(dislikeCollection).InsertOne(ctx, dislike)
	if err != nil {
		log.Default().Printf("Failed to dislike blog: %v", err)
		return ErrUnableToDislikeBlog
	}

	return nil
}

// GetBlogByID implements BlogRepository.
func (b *BlogStorage) GetBlogByID(ctx context.Context, id string) (blog.Blog, error) {
	blogIDPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return blog.Blog{}, ErrInvalidID
	}

	filter := bson.D{{Key: "_id", Value: blogIDPrimitive}}

	var blogData blog.Blog
	err = b.db.Collection(blogCollection).FindOne(ctx, filter).Decode(&blogData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return blog.Blog{}, ErrBlogNotFound
		}
		log.Default().Printf("Failed to get blog by ID: %v", err)
		return blog.Blog{}, ErrUnabletoGetBlog
	}

	return blogData, nil
}

// GetBlogs implements BlogRepository.
func (b *BlogStorage) GetBlogs(ctx context.Context, filterOptions []blog.FilterOption, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[blog.Blog], error) {
	panic("unimplemented")
}

// GetCommentsByBlogID implements BlogRepository.
func (b *BlogStorage) GetCommentsByBlogID(ctx context.Context, blogID string, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[blog.Comment], error) {
	panic("unimplemented")
}

// LikeBlog implements BlogRepository.
func (b *BlogStorage) LikeBlog(ctx context.Context, like blog.Like) error {
	_, err := b.db.Collection(likeCollection).InsertOne(ctx, like)
	if err != nil {
		log.Default().Printf("Failed to like blog: %v", err)
		return ErrUnableToLikeBlog
	}

	return nil
}

// SearchBlogs implements BlogRepository.
func (b *BlogStorage) SearchBlogs(ctx context.Context, query string) (infrastructure.PaginationResponse[blog.Blog], error) {
	panic("unimplemented")
}

// UpdateBlog implements BlogRepository.
func (b *BlogStorage) UpdateBlog(ctx context.Context, id string, blog blog.Blog) error {
	panic("unimplemented")
}

// UnlikeBlog implements BlogRepository.
func (b *BlogStorage) UnlikeBlog(ctx context.Context, like blog.Like) error {
	_, err := b.db.Collection(likeCollection).DeleteOne(ctx, like)
	if err != nil {
		log.Default().Printf("Failed to unlike blog: %v", err)
		return ErrUnableToUnLikeBlog
	}

	return nil
}

// UndislikeBlog implements BlogRepository.
func (b *BlogStorage) UndislikeBlog(ctx context.Context, dislike blog.Dislike) error {
	_, err := b.db.Collection(dislikeCollection).DeleteOne(ctx, dislike)
	if err != nil {
		log.Default().Printf("Failed to unlike blog: %v", err)
		return ErrUnableToUnDislikeBlog
	}

	return nil
}
