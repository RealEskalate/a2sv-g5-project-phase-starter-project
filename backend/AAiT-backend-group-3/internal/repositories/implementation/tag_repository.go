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
        filter := bson.M{"name": tag}
        update := bson.M{"$addToSet": bson.M{"blogs": blogIDObj}}

        updateResult, err := r.collection.UpdateOne(ctx, filter, update)
        if err != nil {
            return err
        }
        if updateResult.MatchedCount == 0 {
            newTag := models.Tag{
                ID:    primitive.NewObjectID(),
                Name:  tag,
                Blogs: []primitive.ObjectID{blogIDObj},
            }
            _, err := r.collection.InsertOne(ctx, newTag)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func (r *MongoTagRepository) GetBlogsByTags(tags []string) ([]string, error) {
	var blogIDStrings []string
	blogIDSet := make(map[primitive.ObjectID]struct{})
	for _, tag := range tags {
		var tagObj models.Tag
		err := r.collection.FindOne(ctx, bson.M{"name": tag}).Decode(&tagObj)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				continue
			}
			return nil, err
		}
		for _, id := range tagObj.Blogs {
			if _, exists := blogIDSet[id]; !exists {
				blogIDSet[id] = struct{}{}
				blogIDStrings = append(blogIDStrings, id.Hex())
			}
		}
	}
	return blogIDStrings, nil
}


func (r *MongoTagRepository) RemoveBlogFromTagList(tagNames []string, blogID string) error {
    blogIDObj, err := primitive.ObjectIDFromHex(blogID)
    if err != nil {
        return err
    }
    for _, tagName := range tagNames {
        updateResult, err := r.collection.UpdateOne(
            ctx,
            bson.M{"name": tagName},
            bson.M{"$pull": bson.M{"blogs": blogIDObj}},
        )
        if err != nil {
            return err
        }
        if updateResult.ModifiedCount > 0 {
            var tag models.Tag
            err = r.collection.FindOne(ctx, bson.M{"name": tagName}).Decode(&tag)
            if err != nil {
                if err == mongo.ErrNoDocuments {
                    continue
                }
                return err
            }
            if len(tag.Blogs) == 0 {
                _, err = r.collection.DeleteOne(ctx, bson.M{"name": tagName})
                if err != nil {
                    return err
                }
            }
        }
    }

    return nil
}

