package repositories

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/repositories/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoTagRepository struct {
	collection *mongo.Collection
}

func NewMongoTagRepository(db *mongo.Database, collectionName string) repository_interface.TagRepositoryInterface {
	return &MongoTagRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *MongoTagRepository) AddBlogToTheTagList(tags []string, blogID string) error {
	blogIDObj, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}
	for _, tag := range tags {
		_, err = r.collection.UpdateOne(
			ctx,
			bson.M{"name": tag},
			bson.M{"$addToSet": bson.M{"blogs": blogIDObj}}, 
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *MongoTagRepository) GetBlogsByTags(tags []string) ([]string, error) {
	var blogIDs []primitive.ObjectID
	for _, tag := range tags {
		var tagObj models.Tag
		err := r.collection.FindOne(ctx, bson.M{"name": tag}).Decode(&tagObj)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				continue
			}
			return nil, err
		}
		blogIDs = append(blogIDs, tagObj.Blogs...)
	}

	var blogIDStrings []string
	for _, id := range blogIDs {
		blogIDStrings = append(blogIDStrings, id.Hex())
	}
	return blogIDStrings, nil
}



func (r *MongoTagRepository) RemoveBlogFromTagList(tagNames []string, blogID string) error {
	blogIDObj, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}
	for _, tagName := range tagNames {
		_, err = r.collection.UpdateOne(
			ctx,
			bson.M{"name": tagName},
			bson.M{"$pull": bson.M{"blogs": blogIDObj}},
		)
		if err != nil {
			return err
		}
	}
	return nil
}