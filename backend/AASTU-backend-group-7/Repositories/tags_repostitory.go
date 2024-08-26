package Repositories

import (
	"blogapp/Domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TagRepository struct {
	tagCollection  Domain.Collection
	postCollection Domain.Collection
}

func NewTagRepository(blogCollection Domain.BlogCollections) *TagRepository {
	return &TagRepository{
		tagCollection:  blogCollection.Tags,
		postCollection: blogCollection.Posts,
	}
}

var ErrDuplicateKey = errors.New("a tag with the same slug already exists")

func (repo *TagRepository) CreateTag(ctx context.Context, tag *Domain.Tag) (error, int) {
	// Attempt to insert a new tag into the collection
	_, err := repo.tagCollection.InsertOne(ctx, tag)
	if err != nil {
		// Check if the error is a MongoDB duplicate key error
		if mongo.IsDuplicateKeyError(err) {
			return ErrDuplicateKey, 409 // HTTP 409 Conflict
		}
		// For other errors, return a generic server error
		return err, 500
	}
	// If successful, return nil error and HTTP 201 Created
	return nil, 201
}

// Define custom errors for clarity and reuse
var (
	ErrTagNotFound = errors.New("tag not found")
	ErrInternal    = errors.New("internal server error")
)

func (repo *TagRepository) DeleteTag(ctx context.Context, slug string) (error, int) {
	// Attempt to delete the tag by slug
	result, err := repo.tagCollection.DeleteOne(ctx, bson.M{"slug": slug})
	if err != nil {
		return ErrInternal, 500
	}
	if result.DeletedCount == 0 {
		// No tag was found with the given slug
		return ErrTagNotFound, 404
	}

	// Attempt to remove the tag from all posts
	_, err = repo.postCollection.UpdateMany(ctx, bson.M{}, bson.M{"$pull": bson.M{"tags": slug}})
	if err != nil {
		return ErrInternal, 500
	}

	return nil, 200
}

func (repo *TagRepository) GetAllTags(ctx context.Context) ([]*Domain.Tag, error, int) {
	tags := []*Domain.Tag{}
	cursor, err := repo.tagCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err, 500
	}
	err = cursor.All(ctx, &tags)
	if err != nil {
		return nil, err, 500
	}
	return tags, nil, 200
}

func (repo *TagRepository) GetTagBySlug(ctx context.Context, slug string) (*Domain.Tag, error, int) {
	tag := &Domain.Tag{}
	err := repo.tagCollection.FindOne(ctx, bson.M{"slug": slug}).Decode(tag)
	if err != nil {
		return nil, err, 500
	}
	return tag, nil, 200
}
