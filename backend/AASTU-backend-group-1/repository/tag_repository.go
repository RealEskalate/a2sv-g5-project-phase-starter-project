package repository

import (
	"blogs/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TagRepository struct {
	collection *mongo.Collection
}
func NewTagRepository(database *mongo.Database) *TagRepository {
	return &TagRepository{collection: database.Collection("tags")}

}
func (t *TagRepository) InsertTag(tag *domain.Tag)  error {
	_, err := t.collection.InsertOne(context.Background(), tag)
	if err != nil {
		return err
	}
	return nil
}

func (t *TagRepository) GetTags() (tags []*domain.Tag,err error) {
	cursor, err := t.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var tag domain.Tag
		err := cursor.Decode(&tag)
		if err != nil {
			return nil, err
		}
		tags = append(tags, &tag)
	}
	return tags, nil
}

func (t *TagRepository) DeleteTag(tag *domain.Tag) error {
	_, err := t.collection.DeleteOne(context.Background(), tag)
	if err != nil {
		return err
	}
	return nil
}

func (t *TagRepository) CheckTag(tags []string) error {
	count, err := t.collection.CountDocuments(context.Background(), bson.M{"_id": bson.M{"$in": tags}})
	if err != nil {
		return errors.New("failed to verify tags")
	}

	if count != int64(len(tags)) {
		return errors.New("invalid tags")
	}

	return nil
}