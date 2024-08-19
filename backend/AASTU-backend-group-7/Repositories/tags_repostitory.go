package Repositories

import (
	"blogapp/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type tagsRepository struct {
	tagCollection  Domain.Collection
	postCollection Domain.Collection
}

func NewTagsRepository(blogCollection Domain.BlogCollections) *tagsRepository {
	return &tagsRepository{
		tagCollection:  blogCollection.Tags,
		postCollection: blogCollection.Posts,
	}
}

func (repo *tagsRepository) CreateTag(ctx context.Context, tag *Domain.Tag) (error, int) {

	_, err := repo.tagCollection.InsertOne(ctx, tag)
	if err != nil {
		return err, 500
	}
	return nil, 201
}

func (repo *tagsRepository) DeleteTag(ctx context.Context, id primitive.ObjectID) (error, int) {
	_, err := repo.tagCollection.DeleteOne(ctx, id)
	if err != nil {
		return err, 500
	}
	// delete tag from all posts
	_, err = repo.postCollection.UpdateMany(ctx, bson.M{"tags": id}, bson.M{"$pull": bson.M{"tags": id}})
	if err != nil {
		return err, 500
	}

	return nil, 200
}

func (repo *tagsRepository) GetAllTags(ctx context.Context) ([]*Domain.Tag, error, int) {
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

func (repo *tagsRepository) GetTagBySlug(ctx context.Context, slug string) (*Domain.Tag, error, int) {
	tag := &Domain.Tag{}
	err := repo.tagCollection.FindOne(ctx, bson.M{"slug": slug}).Decode(tag)
	if err != nil {
		return nil, err, 500
	}
	return tag, nil, 200
}