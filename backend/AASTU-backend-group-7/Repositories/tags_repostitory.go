package Repositories

import (
	"blogapp/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
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

func (repo *TagRepository) CreateTag(ctx context.Context, tag *Domain.Tag) (error, int) {
	// get tag by slug
	_, err := repo.tagCollection.InsertOne(ctx, tag)
	if err != nil {
		return err, 500
	}
	return nil, 201
}

func (repo *TagRepository) DeleteTag(ctx context.Context, slug string) (error, int) {

	// delete tag by slug
	_, err := repo.tagCollection.DeleteOne(ctx, bson.M{"slug": slug})
	if err != nil {
		return err, 500
	}

	// delete tag from all posts
	_, err = repo.postCollection.UpdateMany(ctx, bson.M{}, bson.M{"$pull": bson.M{"tags": slug}})
	if err != nil {
		return err, 500
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
